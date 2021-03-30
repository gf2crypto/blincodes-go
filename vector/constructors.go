package vector

import "strings"
import "fmt"
import "math/rand"
import "time"

// New creates vector object
// Examples:
//   New() creates empty vector
//   New(nil) creates empty vector
//   New(s string) creates vector from string
//   New(n int) creates zero vector of length n
//   New(vec [](u)int(, 8, 16, 32, 64)) creates vector from integer array
//   New(n int, support []int) creates vector of length n from its support
//                             support is set of one's positions
//                             New(5, support []int{0, 3, 1}) -> 11010
//                                                               01234
func New(params ...interface{}) *Vector {
    if len(params) == 0 {
        return newEmpty(0)
    }
    if len(params) == 1 {
        switch b := params[0].(type) {
        case nil:
            return newEmpty(0)
        case string:
            return newFromString(&b)
        case int:
            return newEmpty(b)
        default:
            return fromArray(b)
        }
    }
    if len(params) > 2 {
        panic(fmt.Errorf("vector: too many arguments, expected is 2 or less, but get %v > 2", len(params)))
    }
    if n, ok := params[0].(int); ok {
        if sup, ok1 := params[1].([]int); ok1 {
            return fromSupport(n, sup)
        }
        panic(fmt.Errorf("vector: cannot use constructor New(n int, array []int), expected the second parameter is []int, not %T", params[1]))
    }
    panic(fmt.Errorf("vector: cannot use constructor New(n int, array []int), expected the first parameter is int, not %T", params[0]))
}

//Random returns random vector of length n
func Random(n int) *Vector {
    v := newEmpty(n)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < len(v.body); i++ {
        v.body[i] = rand.Uint64()
    }
    if len(v.body) > 0 && v.lenLast != 0 {
        v.body[len(v.body)-1] &= (((1 << v.lenLast) - 1) << (wordSize - v.lenLast))
    }
    return v
}

//PackBytes packs byte array into the vector
//  Examples:
//  PackBytes([]byte{0x01, 0x02, 0x03}, 24) -> 000000010000001000000011
//  PackBytes([]byte{0x01, 0x02, 0x03}, 11) -> 00000001000
//  PackBytes([]byte{0x01, 0x02, 0x03}, 30) -> 000000010000001000000011000000
//  PackBytes([]byte{}, 3) -> 000
func PackBytes(b []byte, n int) *Vector {
    nBytes := n / 8
    if n % 8 != 0 {
        nBytes += 1
    }
    r := wordSize / 8
    lenBody := n / wordSize
    lenLast := n % wordSize
    if lenLast != 0 {
        lenBody++
    }
    body := make([]uint64, lenBody)
    for i:=0; i < lenBody; i++{
        if r * i >= len(b) {
            break
        }
        for j := 0; j < r; j++ {
            body[i] <<= 8
            if r*i+j < len(b) {
                body[i] ^= uint64(b[r*i+j])
            }
        }
    }
    return &Vector{body: body, lenLast: lenLast}
}

// newFromString converts string to Vector
// Function supports the following filler for zero symbol:
//       '0' == '0', '-'
// Example:
//   "--110-1- - - 1-   0 " -> 0011001000100
func newFromString(s *string) *Vector {
    if len(*s) == 0 {
        return newEmpty(0)
    }
    // len > 0
    // remove all white spaces
    tmp := strings.ReplaceAll(*s, " ", "")
    v := newEmpty(len(tmp))
    for i, bit := range tmp {
        switch {
        case bit == '1':
            v.body[i/wordSize] ^= (1 << (wordSize - (i % wordSize) - 1))
        case bit == '0':
        default:
            e := fmt.Errorf("vector: parse string %s error, unexpected symbol %c in position %d, possible only {' ', '0', '1', '-'}",
                tmp, bit, i)
            panic(e)
        }
    }
    return v
}

func fromArray(array interface{}) *Vector {
    slice, e := ToSlicer(array)
    if e != nil {
        panic(fmt.Errorf("vector: unsupported type %T, supported only: [](u)int(8,16, 32, 64), (u)int, string, nil", array))
    }
    v := newEmpty(slice.Len())
    for i := 0; i < slice.Len(); i++ {
        switch bit := slice.GetElement(i); {
        case bit == 1:
            v.body[i/wordSize] ^= (1 << (wordSize - (i % wordSize) - 1))
        case bit != 0:
            panic(fmt.Errorf("vector: unexpected digit %d in position %d, possible only {0, 1}",
                bit, i))
        }
    }
    return v
}

func newEmpty(n int) *Vector {
    if n <= 0 {
        return &Vector{body: make([]uint64, 0, 0), lenLast: 0}
    }
    lenBlock := int(n / wordSize)
    lenLast := n % wordSize
    if lenLast != 0 {
        lenBlock++
    }
    return &Vector{body: make([]uint64, lenBlock), lenLast: lenLast}
}

func fromSupport(n int, sup []int) *Vector {
    if n <= 0 {
        return newEmpty(0)
    }
    // n > 1
    v := newEmpty(n)
    for _, index := range sup {
        if index >= n || index < 0 {
            continue
        }
        // 0 <= index <= n - 1
        v.body[index/wordSize] ^= (1 << (wordSize - (index % wordSize) - 1))
    }
    return v
}
