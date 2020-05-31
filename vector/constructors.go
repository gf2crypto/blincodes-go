package vector

import "strings"
import "fmt"

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
