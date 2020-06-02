package vector

import "bytes"
import "fmt"
import "strings"

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

// PrettyString returns pretty formatted string of vector representation
// Example:
// 0101011 -> -1-1-11
func (v *Vector) PrettyString() string {
    var buf bytes.Buffer
    zeroSub := "-"
    fs := fmt.Sprintf("%%0%db", wordSize)
    for _, w := range v.body {
        fmt.Fprint(&buf, strings.ReplaceAll(fmt.Sprintf(fs, w), "0", zeroSub))
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

//Less return true if vector v is less vector v0
//Return v < v0?
func (v *Vector) Less(v0 *Vector) bool {
    if v.Len() < v0.Len() {
        return true
    }
    if v.Len() > v0.Len() {
        return false
    }
    if v.Len() != 0 {
        for i, b := range v.body {
            switch {
            case b < v0.body[i]:
                return true
            case b > v0.body[i]:
                return false
            default:
                continue
            }
        }
    }
    return false
}

//More return true if vector v is more vector v0
//Return v > v0?
func (v *Vector) More(v0 *Vector) bool {
    if v.Len() < v0.Len() {
        return false
    }
    if v.Len() > v0.Len() {
        return true
    }
    if v.Len() != 0 {
        for i, b := range v.body {
            switch {
            case b < v0.body[i]:
                return false
            case b > v0.body[i]:
                return true
            default:
                continue
            }
        }
    }
    return false
}

//NoMore return true if vector v is no more vector v0
//Return v <= v0?
func (v *Vector) NoMore(v0 *Vector) bool {
    if v.Len() < v0.Len() {
        return true
    }
    if v.Len() > v0.Len() {
        return false
    }
    if v.Len() != 0 {
        for i, b := range v.body {
            switch {
            case b < v0.body[i]:
                return true
            case b > v0.body[i]:
                return false
            default:
                continue
            }
        }
    }
    return true
}

//NoLess return true if vector v is no less vector v0
//Return v >= v0?
func (v *Vector) NoLess(v0 *Vector) bool {
    if v.Len() < v0.Len() {
        return false
    }
    if v.Len() > v0.Len() {
        return true
    }
    if v.Len() != 0 {
        for i, b := range v.body {
            switch {
            case b < v0.body[i]:
                return false
            case b > v0.body[i]:
                return true
            default:
                continue
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
// return v ^ v0
func (v *Vector) Xor(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    res := v.Copy()
    for i, b := range v0.body {
        res.body[i] ^= b
    }
    return res
}

//Or evaluates or of two vectors.
//return v | v0
func (v *Vector) Or(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    res := v.Copy()
    for i, b := range v0.body {
        res.body[i] |= b
    }
    return res
}

//And evaluates and of two vectors.
//return v & v0
func (v *Vector) And(v0 *Vector) *Vector {
    if v.Len() != v0.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            v.Len(), v0.Len()))
    }
    res := v.Copy()
    for i, b := range v0.body {
        res.body[i] &= b
    }
    return res
}

//Not evaluates not of vectors.
//return ^v
func (v *Vector) Not() *Vector {
    res := v.Copy()
    for i, b := range res.body {
        res.body[i] = ^b
    }
    if res.lenLast != 0 {
        res.body[len(res.body)-1] &= (((1 << res.lenLast) - 1) << (wordSize - res.lenLast))
    }
    return res
}

//Concatenate concatenates of two vectors
// return (v || v0)
func (v *Vector) Concatenate(v0 *Vector) *Vector {
    if v0 == nil || len(v0.body) == 0 {
        return v
    }
    if v.Len() == 0 {
        // just vopy v0 to v
        return &Vector{body: append(make([]uint64, 0, len(v0.body)), v0.body...), lenLast: v0.lenLast}
    }
    if v.lenLast == 0 {
        return &Vector{body: append(v.body, v0.body...), lenLast: v0.lenLast}
    }
    r := wordSize - v.lenLast
    mask := ((1 << r) - 1) << v.lenLast
    if v0.lenLast <= r {
        body := make([]uint64, 0, len(v.body)+len(v0.body)-1)
        body = append(body, v.body...)
        if len(v0.body) > 1 {
            body = append(body, v0.ShiftLeft(uint(r)).body[:len(v0.body)-1]...)
        }
        body[len(v.body)-1] ^= ((v0.body[0] & uint64(mask)) >> v.lenLast)
        return &Vector{body: body, lenLast: (v0.lenLast + v.lenLast) % wordSize}
    }
    body := append(v.body, v0.ShiftLeft(uint(r)).body...)
    body[len(v.body)-1] ^= ((v0.body[0] & uint64(mask)) >> v.lenLast)
    return &Vector{body: body, lenLast: v0.lenLast - r}
}

//Resize changes the length of vector
//Resize(delta) sets length to v.Len() + delta.
//That is if delta < 0 then vector is zip, and if delta > 0
func (v *Vector) Resize(delta int) *Vector {
    n := v.Len() + delta
    if n <= 0 {
        return newEmpty(0)
    }
    w := newEmpty(n)
    end := len(w.body)
    if end > len(v.body) {
        end = len(v.body)
    }
    for i := 0; i < end; i++ {
        w.body[i] = v.body[i]
    }
    w.body[len(w.body)-1] &= (((1 << w.lenLast) - 1) << (wordSize - w.lenLast))
    return w
}

// set sets i-th bit of vector to value
func (v *Vector) set(i int, val byte) *Vector {
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
