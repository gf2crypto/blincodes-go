package vector

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type word uint

const WordSize = uint(unsafe.Sizeof(word(0))) << 3
const MaxInteger = ^word(0)

//Vector represents binary vector
//This type is immutable
type Vector struct {
	body    []word // packed binary vector
	lenLast uint   //len of the last element of body array
}

//New returns empty vector
func New() *Vector {
	return &Vector{body: nil, lenLast: WordSize}
}

//reAllocate reallocates memory for vector, n is new length of vector, returns m
func (v *Vector) reAllocate(n uint) *Vector {
	var lenBlock uint
	lenBlock, v.lenLast = getShapes(n)
	switch {
	case lenBlock == 0:
		v.body = nil
	case cap(v.body) < int(lenBlock):
		v.body = make([]word, lenBlock)
	default:
		v.body = v.body[:lenBlock]
	}
	return v
}

//allocate reallocates memory for vector, n is new length of vector, returns v
func (v *Vector) allocate(n uint) *Vector {
	var lenBlock uint
	lenBlock, v.lenLast = getShapes(n)
	switch {
	case lenBlock == 0:
		v.body = nil
	default:
		v.body = make([]word, lenBlock)
	}
	return v
}

//getShapes calculates the shapes of vector by the length.
//It returns length of vector's body and length of the vector's last block
func getShapes(n uint) (uint, uint) {
	lenBlock := n / WordSize
	lenLast := n % WordSize
	if lenLast != 0 && n != 0 {
		lenBlock++
	} else {
		lenLast = WordSize
	}
	return lenBlock, lenLast
}

//SetZero sets v to a units vector of length n, returns v
func (v *Vector) SetZero(n uint) *Vector {
	return v.allocate(n)
}

//SetV sets v to u, returns v
func (v *Vector) SetV(u *Vector) *Vector {
	if u == v {
		return v
	}
	v.reAllocate(u.Len())
	copy(v.body, u.body)
	return v
}

// Len returns length of vector v
func (v *Vector) Len() uint {
	if v.body == nil {
		return 0
	}
	return (uint(len(v.body))-1)*WordSize + v.lenLast
}

//SetUnits sets v to a units vector of length n
func (v *Vector) SetUnits(n uint) *Vector {
	v.reAllocate(n)
	for i := 0; i < len(v.body); i++ {
		v.body[i] = ^word(0)
		if i+1 == len(v.body) {
			v.body[i] &= ((1 << v.lenLast) - 1) << (WordSize - v.lenLast)
		}
	}
	return v
}

//SetBytes packs byte array b into the vector of length n and sets v to that new vector
//  Examples:
//  SetBytes([]byte{0x01, 0x02, 0x03}, 24) -> 000000010000001000000011
//  SetBytes([]byte{0x01, 0x02, 0x03}, 11) -> 00000001000
//  SetBytes([]byte{0x01, 0x02, 0x03}, 30) -> 000000010000001000000011000000
//  SetBytes([]byte{}, 3) -> 000
func (v *Vector) SetBytes(b []byte, n uint) *Vector {
	nBytes := n / 8
	if n%8 != 0 {
		nBytes += 1
	}
	r := WordSize / 8
	v.reAllocate(n)
	for i := uint(0); i < uint(len(v.body)); i++ {
		if r*i >= uint(len(b)) {
			break
		}
		v.body[i] = 0
		for j := uint(0); j < r; j++ {
			v.body[i] <<= 8
			if r*i+j < uint(len(b)) {
				v.body[i] ^= word(b[r*i+j])
			}
		}
	}
	return v
}

//SetRandom sets v to a random vector of length n
func (v *Vector) SetRandom(n uint) *Vector {
	b := make([]byte, n/8+1)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return v.SetBytes(b, n)
}

// Parse converts string to Vector and sets v to result, returns v
// Function supports the following filler for zero symbol:
//       "0" == "0", "-", "*"
// Example:
//   "--110-1- - - 1-  * 11 0 " -> 0011001000100110
func (v *Vector) Parse(s string) (*Vector, error) {
	repl := map[string]string{
		" ": "",
		"-": "0",
		"*": "0",
	}
	for oldS, newS := range repl {
		s = strings.ReplaceAll(s, oldS, newS)
	}
	if len(s) == 0 {
		v.body = nil
		v.lenLast = WordSize
		return v, nil
	}
	w := New().allocate(uint(len(s)))
	for i, j := uint(0), 0; j < len(w.body); j++ {
		max := i + WordSize
		if max > uint(len(s)) {
			max = uint(len(s))
		}
		t, err := strconv.ParseUint(s[i:max], 2, int(WordSize))
		if err != nil {
			return v, err
		}
		w.body[j] = word(t)
		i = max
	}
	w.body[len(w.body)-1] <<= WordSize - w.lenLast
	v.body, v.lenLast = w.body, w.lenLast
	return v, nil
}

//SetBitArray converts bit array to Vector and sets v to the result
func (v *Vector) SetBitArray(array []byte) *Vector {
	v.reAllocate(uint(len(array)))
	for i, b := range array {
		switch {
		case b == 1:
			v.body[i/int(WordSize)] ^= 1 << (WordSize - (uint(i) % WordSize) - 1)
		case b != 0:
			panic(fmt.Errorf("vector: unexpected digit %d in position %d, "+
				"possible only {0, 1}", b, i))
		}
	}
	return v
}

//SetSupport sets v to vector with units from the support array
// Example:
//     v.SetSupport(10, []int{0, 1, 5, 9}) -> 1100010001
func (v *Vector) SetSupport(n uint, sup []uint) *Vector {
	v.SetZero(n)
	for _, i := range sup {
		if i >= n {
			continue
		}
		// 0 <= index <= n - 1
		v.body[i/WordSize] ^= 1 << (WordSize - (i % WordSize) - 1)
	}
	return v
}

//Concatenate sets v to concatenation of u and w and returns v
func (v *Vector) Concatenate(u, w *Vector) *Vector {
	if u.Len() == 0 {
		return v.SetV(w)
	}
	if w.Len() == 0 {
		return v.SetV(u)
	}
	tmpv := New().allocate(u.Len() + w.Len())
	copy(tmpv.body, u.body)
	tmpv.body[len(u.body) - 1] ^= w.body[0] >> u.lenLast
	for i := len(u.body); i < len(tmpv.body); i++ {
			j := i - len(u.body)
			tmpv.body[i] = w.body[j] << (WordSize - u.lenLast)
			if j+1 < len(w.body) {
				tmpv.body[i] ^= w.body[j+1] >> u.lenLast
			}
	}
	v.body, v.lenLast = tmpv.body, tmpv.lenLast
	return v
}

//Resize resizes vector w and sets the result to v, returns v
func (v *Vector) Resize(w *Vector, r int) *Vector {
	newLen := uint(0)
	v.SetV(w)
	if r < 0 {
		if v.Len() <= uint(-r) {
			return v.SetZero(0)
		}
		newLen = v.Len() - uint(-r)
 	} else {
 		newLen = v.Len() + uint(r)
	}
	u := new(Vector).reAllocate(newLen)
	copy(u.body, w.body)
	v.body, v.lenLast = u.body, u.lenLast
	v.body[len(v.body) - 1] &= ^word(0) << (WordSize - v.lenLast)
	return v
}

// String returns string representation of Vector
func (v *Vector) String() string {
	var buf bytes.Buffer
	fs := fmt.Sprintf("%%0%db", WordSize)
	for _, w := range v.body {
		_, err := fmt.Fprintf(&buf, fs, w)
		if err != nil {
			panic(err)
		}
	}
	s := buf.String()
	return s[:(len(s) - int(WordSize-v.lenLast))]
}

// PrettyString returns pretty formatted string of vector representation
// Example:
// 0101011 -> -1-1-11
func (v *Vector) PrettyString() string {
	var buf bytes.Buffer
	zeroSub := "-"
	fs := fmt.Sprintf("%%0%db", WordSize)
	for _, w := range v.body {
		_, err := fmt.Fprint(&buf, strings.ReplaceAll(fmt.Sprintf(fs, w), "0", zeroSub))
		if err != nil {
			panic(err)
		}
	}
	s := buf.String()
	return s[:(len(s) - int(WordSize-v.lenLast))]
}

// LaTeXString returns string of vector representation to use in LaTeX matrix environment
// Example:
// 0101011 -> 0&1&0&1&0&1&1
func (v *Vector) LaTeXString() string {
	return strings.Join(strings.Split(v.String(), ""), "&")
}

// Cmp compares vector v and vector u and returns
// -1 if v < u
//  0 if v == u
//  1 if v > u
// v < u if len(u) > len(v) or len(u)==len(v) and for some i holds
// v[0]=u[0],..., v[i-1]=u[i-1], v[i] < u[i]
// For example,
//     10 < 000110
//     000010 < 000110
func (v *Vector) Cmp(u *Vector) int {
	if v.Len() < u.Len() {
		return -1
	}
	if v.Len() > u.Len() {
		return 1
	}
	for i, b := range v.body {
		switch {
		case b < u.body[i]:
			return -1
		case b > u.body[i]:
			return 1
		}
	}
	return 0
}

// Get returns i-th bit of vector
func (v *Vector) Get(i uint) byte {
	if i < 0 {
		panic(fmt.Errorf("vector: index error %d (expected non-negative integer)", i))
	}
	if i >= v.Len() {
		panic(fmt.Errorf("vector: index %d out of range, expected integer in [0, %d)",
			i, v.Len()))
	}
	if v.body[i/WordSize]&(1<<(WordSize-i%WordSize-1)) == 0 {
		return 0
	}
	return 1
}

// SetBit sets v to vector w with change coordinate i by bit b % 2
// m = w[i] <- b%2
func (v *Vector) SetBit(w *Vector, i uint, b byte) *Vector {
    v.SetV(w)
    n := i / WordSize
    if b % 2 == 0 {
		v.body[n] &= ^word(0) ^ (1 << (WordSize-i%WordSize-1))
	} else {
		v.body[n] |= 1 << (WordSize-i % WordSize-1)
	}
	return v
}

// Bits returns all bits of vector as slice of bytes
func (v *Vector) Bits() []byte {
	bits := make([]byte, 0, v.Len())
	for i := uint(0); i < v.Len(); i++ {
		bits = append(bits, v.Get(i))
	}
	return bits
}

//Wt returns Hamming weight of vector
func (v *Vector) Wt() int {
	wt := 0
	for i := uint(0); i < v.Len(); i++ {
		wt += int(v.Get(i))
	}
	return wt
}

//Support returns support of vector.
//support is set of 1's indexes of vector
func (v *Vector) Support() []uint {
	sup := make([]uint, 0, v.Len())
	for i := uint(0); i < v.Len(); i++ {
		if v.Get(i) == 1 {
			sup = append(sup, i)
		}
	}
	return sup
}

//Zeros returns set of 0's indexes of vector.
func (v *Vector) Zeros() []uint {
	z := make([]uint, 0, v.Len())
	for i := uint(0); i < v.Len(); i++ {
		if v.Get(i) == 0 {
			z = append(z, i)
		}
	}
	return z
}

//Xor sets v to u XOR w and returns v
func (v *Vector) Xor(u, w *Vector) *Vector {
	if u == nil || w == nil {
		return v
	}
	if u.Len() != w.Len() {
		panic(fmt.Errorf("vector: vectors have different length: %d != %d",
			u.Len(), w.Len()))
	}
	v.reAllocate(u.Len())
	for i, b := range u.body {
		v.body[i] = b ^ w.body[i]
	}
	return v
}

////XorV sets v to u1 XOR u2 XOR ... XOR un and returns v
////XorV is vectorised version of Xor function.
//func (v *Vector) XorV(u []*Vector) *Vector {
//	body := make([]word, 0)
//	var lenLast uint
//	var t uint
//	for i, w := range u {
//		if i == 0 {
//			t = w.Len()
//			body, lenLast = makeVector(t)
//		} else {
//			if w.Len() != t {
//				panic(fmt.Errorf("vector: vectors have different length: %d != %d",
//					t, w.Len()))
//			}
//		}
//		for j, b := range w.body {
//			body[j] ^= b
//		}
//	}
//	v.body, v.lenLast = body, lenLast
//	return v
//}

//Or sets v to u OR w and returns v
func (v *Vector) Or(u, w *Vector) *Vector {
	if u == nil || w == nil {
		return v
	}
	if u.Len() != w.Len() {
		panic(fmt.Errorf("vector: vectors have different length: %d != %d",
			u.Len(), w.Len()))
	}
	v.reAllocate(u.Len())
	for i, b := range u.body {
		v.body[i] = b | w.body[i]
	}
	return v
}

////OrV sets v to u1 OR u2 OR ... OR un and returns v
////OrV is vectorised version of Or function.
//func (v *Vector) OrV(u []*Vector) *Vector {
//	body := make([]word, 0)
//	var lenLast uint
//	var t uint
//	for i, w := range u {
//		if i == 0 {
//			t = w.Len()
//			body, lenLast = makeVector(t)
//		} else {
//			if w.Len() != t {
//				panic(fmt.Errorf("vector: vectors have different length: %d != %d",
//					t, w.Len()))
//			}
//		}
//		for j, b := range w.body {
//			body[j] |= b
//		}
//	}
//	v.body, v.lenLast = body, lenLast
//	return v
//}

//And sets v to u AND w and returns v
func (v *Vector) And(u, w *Vector) *Vector {
	if u == nil || w == nil {
		return v
	}
	if u.Len() != w.Len() {
		panic(fmt.Errorf("vector: vectors have different length: %d != %d",
			u.Len(), w.Len()))
	}
	v.reAllocate(u.Len())
	for i, b := range u.body {
		v.body[i] = b & w.body[i]
	}
	return v
}

////AndV sets v to u1 AND u2 AND ... AND un and returns v
////AndV is vectorised version of And function.
//func (v *Vector) AndV(u []*Vector) *Vector {
//	body := make([]word, 0)
//	var lenLast uint
//	var t uint
//	for i, w := range u {
//		if i == 0 {
//			t = w.Len()
//			body, lenLast = makeVector(t)
//		} else {
//			if w.Len() != t {
//				panic(fmt.Errorf("vector: vectors have different length: %d != %d",
//					t, w.Len()))
//			}
//		}
//		for j, b := range w.body {
//			if i == 0 {
//				body[j] = MaxInteger
//			}
//			body[j] &= b
//		}
//	}
//	v.body, v.lenLast = body, lenLast
//	return v
//}

//Not sets v to ^u and returns v
func (v *Vector) Not(u *Vector) *Vector {
	if u == nil {
		return v
	}
	v.reAllocate(u.Len())
	for i, b := range u.body {
		v.body[i] = ^b
		if i+1 == len(u.body) {
			v.body[i] &= ((1 << v.lenLast) - 1) << (WordSize - v.lenLast)
		}
	}
	return v
}

//Iter iterates over elements
func (v *Vector) Iter() <-chan byte {
	ch := make(chan byte)
	go func() {
		defer close(ch)
		for i := uint(0); i < v.Len(); i++ {
			ch <- byte(v.Get(i))
		}
	}()
	return ch
}

//ShiftRight shifts vector w for r position right and sets vector v to shifted w, returns v
// That's it v = w >> r
func (v *Vector) ShiftRight(w *Vector, r uint) *Vector {
	//body, lenLast := makeVector(w.Len())
	start := int(r / WordSize)
	l := r % WordSize
	mask := MaxInteger >> (WordSize - l)
	tmpv := New().allocate(w.Len())
	for i := start; i < len(w.body); i++ {
		tmpv.body[i] = w.body[i-start] >> l
		if i > start && mask != 0 {
			tmpv.body[i] ^= (w.body[i-start-1] & mask) << (WordSize - l)
		}
		if i == len(tmpv.body)-1 {
			tmpv.body[i] &= MaxInteger << (WordSize - tmpv.lenLast)
		}
	}
	v.body, v.lenLast = tmpv.body, tmpv.lenLast
	return v
}

//ShiftLeft shifts vector w for r position left and sets vector v to shifted w, returns v
// That's it v = w << r
func (v *Vector) ShiftLeft(w *Vector, r uint) *Vector {
	//body, lenLast := makeVector(w.Len())
	start := int(r / WordSize)
	l := r % WordSize
	mask := MaxInteger << (WordSize - l)
	tmpv := New().allocate(w.Len())
	for i := start; i < len(w.body); i++ {
		tmpv.body[i-start] = w.body[i] << l
		if i < len(w.body)-1 {
			tmpv.body[i-start] ^= (w.body[i+1] & mask) >> (WordSize - l)
		}
	}
	v.body, v.lenLast = tmpv.body, tmpv.lenLast
	return v
}

//RotateLeft cyclical shifts vector w for r position left and sets vector v to shifted w, returns v
// That's it v = w rotl r
func (v *Vector) RotateLeft(w *Vector, r uint) *Vector {
	if w.Len() == 0 {
		v.allocate(0)
		return v
	}
	r = r % w.Len()
	u := new(Vector).ShiftLeft(w, r) // u == w << r
	v.ShiftRight(w, w.Len()-r)       // v == w >> (len - r)
	return v.Xor(v, u)
}

//RotateRight cyclical shifts vector w for r position right and sets vector v to shifted w, returns v
// That's it v = w rotr r
func (v *Vector) RotateRight(w *Vector, r uint) *Vector {
	if w.Len() == 0 {
		v.allocate(0)
		return v
	}
	r = r % w.Len()
	u := new(Vector).ShiftLeft(w, w.Len()-r) // u == w << (len - r)
	v.ShiftRight(w, r)                       // v == w >> r
	return v.Xor(v, u)
}
