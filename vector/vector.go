package vector

import (
    "bytes"
    "unsafe"
)
import "fmt"
import "strings"

type word uint
const WordSize = uint(unsafe.Sizeof(word(0))) << 3
const MaxInteger = ^word(0)

//Vector represents binary vector
//This type is immutable
type Vector struct {
    body    []word // packed binary vector
    lenLast uint      //len of the last element of body array
}

func (v *Vector) String() string {
    var buf bytes.Buffer
    fs := fmt.Sprintf("%%0%db", WordSize)
    for _, w := range v.body {
        fmt.Fprintf(&buf, fs, w)
    }
    s := buf.String()
    return s[:(len(s) - int(WordSize- v.lenLast))]
}

// PrettyString returns pretty formatted string of vector representation
// Example:
// 0101011 -> -1-1-11
func (v *Vector) PrettyString() string {
    var buf bytes.Buffer
    zeroSub := "-"
    fs := fmt.Sprintf("%%0%db", WordSize)
    for _, w := range v.body {
        fmt.Fprint(&buf, strings.ReplaceAll(fmt.Sprintf(fs, w), "0", zeroSub))
    }
    s := buf.String()
    return s[:(len(s) - int(WordSize- v.lenLast))]
}

// Len returns length of vector
func (v *Vector) Len() int {
    if len(v.body) == 0 {
        return 0
    }
    l := len(v.body) << 6
    if v.lenLast != 0 {
        l -= int(WordSize - v.lenLast)
    }
    return l
}

// Cmp compares vector v and vector u and returns
// -1 if v < u
//  0 if v == u
//  1 if v > u
func (v *Vector) Cmp(u *Vector) int {
    if v.Len() < u.Len() {
        return -1
    }
    if v.Len() > u.Len() {
        return 1
    }
    for i, b := range v.body {
        switch {
        case b < u.body[i]: return -1
        case b > u.body[i]: return 1
        }
    }
    return 0
}

//GetFirstOne returns index of the first one or -1 if it is zero
func (v *Vector) GetFirstOne() int {
    for i := 0; i < v.Len(); i++ {
        if v.Get(i) == 1 {
            return i
        }
    }
    return -1
}

// Copy returns copy of vector
func (v *Vector) Copy() *Vector {
    return &Vector{
        body:    append(make([]word, 0, len(v.body)), v.body...),
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
    if v.body[i/int(WordSize)]&(1<<(int(WordSize)-i%int(WordSize)-1)) == 0 {
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

//Xor sets v to u XOR w and returns v
func (v *Vector) Xor(u, w *Vector) *Vector {
    if u.Len() != w.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            u.Len(), w.Len()))
    }
    v.body = make([]word, u.Len())
    for i, b := range u.body {
        v.body[i] = b ^ w.body[i]
    }
    return v
}

//Or sets v to u OR w and returns v
func (v *Vector) Or(u, w *Vector) *Vector {
    if u.Len() != w.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            u.Len(), w.Len()))
    }
    v.body = make([]word, u.Len())
    for i, b := range u.body {
        v.body[i] = b | w.body[i]
    }
    return v
}

//And sets v to u AND w and returns v
func (v *Vector) And(u, w *Vector) *Vector {
    if u.Len() != w.Len() {
        panic(fmt.Errorf("vector: vectors have different length: %d != %d",
            u.Len(), w.Len()))
    }
    v.body = make([]word, u.Len())
    for i, b := range u.body {
        v.body[i] = b & w.body[i]
    }
    return v
}

//Not sets v to ^u and returns v
func (v *Vector) Not(u *Vector) *Vector {
    v.body = make([]word, u.Len())
    for i, b := range u.body {
        v.body[i] = ^b
    }
    v.lenLast = u.lenLast
    if v.lenLast != 0 {
        v.body[len(v.body)-1] &= ((1 << v.lenLast) - 1) << (WordSize - v.lenLast)
    }
    return v
}

//Concatenate sets v to concatenation of u and w and returns v
func (v *Vector) Concatenate(u, w *Vector) *Vector {
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

//Iter iterates over elements
func (v *Vector) Iter() <-chan byte {
    ch := make(chan byte)
    go func() {
        defer close(ch)
        for i := 0; i < v.Len(); i++ {
            ch <- byte(v.Get(i))
        }
    }()
    return ch
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
        v.body[i/int(WordSize)] &= MaxInteger ^ (1 << (int(WordSize) - i%int(WordSize) - 1))
    } else {
        v.body[i/int(WordSize)] |= 1 << (int(WordSize) - i%int(WordSize) - 1)
    }
    return v
}
