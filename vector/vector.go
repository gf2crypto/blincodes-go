package vector

import "bytes"
import "fmt"

const wordSize = 64
const maxInt = 0xFFFFFFFFFFFFFFFF

//Vector represents binary vector
//This type is mutable
//Almost all functions which operates over Vectors
//change the recipient
type Vector struct {
    body    []uint64 // packed binary vector
    lenLast int      //len of the last element of body array
}

func (v *Vector) String() string {
    var buf bytes.Buffer
    fs := fmt.Sprintf("%%0%db", wordSize)
    for _, w := range v.body {
        fmt.Fprintf(&buf, fs, w)
    }
    if v.lenLast != 0 {
        s := buf.String()
        return s[:(len(s) - (wordSize - v.lenLast))]
    }
    return buf.String()
}

// Len returns len of vector
func (v *Vector) Len() int {
    if len(v.body) == 0 {
        return 0
    }
    l := (len(v.body) << 6)
    if v.lenLast != 0 {
        l -= (wordSize - v.lenLast)
    }
    return l
}

//Equal return true if vector v is equal vector v0
//Return v == v0?
func (v *Vector) Equal(v0 *Vector) bool {
    if v.Len() != v0.Len() {
        return false
    }
    if v.Len() != 0 {
        for i, b := range v.body {
            if v0.body[i] != b {
                return false
            }
        }
    }
    return true
}

// Copy returns copy of vector
func (v *Vector) Copy() *Vector {
    return &Vector{
        body:    append(make([]uint64, 0, len(v.body)), v.body...),
        lenLast: v.lenLast,
    }
}

// Get returns i-th bit of vector
func (v *Vector) Get(i int) byte {
    if i < 0 {
        panic(fmt.Errorf("vector: index error %d (expected non-negative integer)", i))
    }
    if i >= v.Len() {
        panic(fmt.Errorf("vector: index %d out of range, expected integer in [0, %d)",
            i, v.Len()))
    }
    if v.body[i/wordSize]&(1<<(wordSize-i%wordSize-1)) == 0 {
        return 0
    }
    return 1
}

// Set sets i-th bit of vector to value
func (v *Vector) Set(i int, val byte) *Vector {
    if i < 0 {
        panic(fmt.Errorf("vector: index error %d (expected non-negative integer)", i))
    }
    if i >= v.Len() {
        panic(fmt.Errorf("vector: index %d out of range, expected integer in [0, %d)",
            i, v.Len()))
    }
    if val == 0 {
        v.body[i/wordSize] &= (maxInt ^ (1 << (wordSize - i%wordSize - 1)))
    } else {
        v.body[i/wordSize] |= (1 << (wordSize - i%wordSize - 1))
    }
    return v
}

// Bits returns all bits of vector as slice of bytes
func (v *Vector) Bits() []byte {
    bits := make([]byte, 0, v.Len())
    for i := 0; i < v.Len(); i++ {
        bits = append(bits, v.Get(i))
    }
    return bits
}

//Wt returns Hamming weight of vector
func (v *Vector) Wt() int {
    wt := 0
    for i := 0; i < v.Len(); i++ {
        wt += int(v.Get(i))
    }
    return wt
}

//Support returns support of vector.
//support is set of 1's indexes of vector
func (v *Vector) Support() []int {
    sup := make([]int, 0, v.Len())
    for i := 0; i < v.Len(); i++ {
        if v.Get(i) == 1 {
            sup = append(sup, i)
        }
    }
    return sup
}

//Zeroes returns set of 0's indexes of vector.
func (v *Vector) Zeroes() []int {
    zeroes := make([]int, 0, v.Len())
    for i := 0; i < v.Len(); i++ {
        if v.Get(i) == 0 {
            zeroes = append(zeroes, i)
        }
    }
    return zeroes
}

//Xor evaluates xor of two vectors.
//Results of operation is assigned to original vector
//v ^= v0
func (v *Vector) Xor(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    for i, b := range v0.body {
        v.body[i] ^= b
    }
    return v
}

//Or evaluates or of two vectors.
//Results of operation is assigned to original vector
//v |= v0
func (v *Vector) Or(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    for i, b := range v0.body {
        v.body[i] |= b
    }
    return v
}

//And evaluates and of two vectors.
//Results of operation is assigned to original vector
//v &= v0
func (v *Vector) And(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    for i, b := range v0.body {
        v.body[i] &= b
    }
    return v
}

//Not evaluates not of vectors.
//Results of operation is assigned to original vector
//v = ^v0
func (v *Vector) Not() *Vector {
    for i, b := range v.body {
        v.body[i] = ^b
    }
    if v.lenLast != 0 {
        v.body[len(v.body)-1] &= (((1 << v.lenLast) - 1) << (wordSize - v.lenLast))
    }
    return v
}

//Concatenate concatenates of two vectors
//Results of operation is assigned to original vector
// v = (v || v0)
func (v *Vector) Concatenate(v0 *Vector) *Vector {
    if v0 == nil || len(v0.body) == 0 {
        return v
    }
    if v.Len() == 0 {
        // just vopy v0 to v
        v.body = append(make([]uint64, 0, len(v0.body)), v0.body...)
        v.lenLast = v0.lenLast
        return v
    }
    if v.lenLast == 0 {
        v.body = append(v.body, v0.body...)
        v.lenLast = v0.lenLast
        return v
    }
    // v.lenLast !=0, len(v0.body) >= 1
    // We have to shift of v0 to left
    newLen := v.Len() + v0.Len()
    newSize := int(newLen / wordSize)
    newLenLast := newLen % wordSize
    if newLenLast != 0 {
        newSize++
    }
    // resize body
    oldSize := len(v.body) // oldSize >=1
    v.body = append(v.body, make([]uint64, newSize-len(v.body))...)
    var mask uint64 = ((1 << (wordSize - v.lenLast)) - 1) << v.lenLast
    var rest uint64
    j := 0
    for i := oldSize; i < len(v.body)+1; i++ {
        rest = (v0.body[j] & mask) >> v.lenLast
        v.body[i-1] ^= rest
        if i < len(v.body) {
            v.body[i] = v0.body[j] << (wordSize - v.lenLast)
        }
        j++
    }
    v.lenLast = newLenLast
    return v
}
