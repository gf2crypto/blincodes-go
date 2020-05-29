package vector

import "strings"
import "fmt"

// New creates vector object
func New(body interface{}) (*Vector, error) {
    switch b := body.(type) {
    case nil:
        return newEmpty(0), nil
    case string:
        return newFromString(&b)
    case int:
        return newEmpty(b), nil
    default:
        {
            slice, e := ToSlicer(b)
            if e != nil {
                return nil, fmt.Errorf("vector: unsupported type %T, supported only: [](u)int(8,16, 32, 64), (u)int, string, nil", b)
            }
            v := newEmpty(slice.Len())
            for i := 0; i < slice.Len(); i++ {
                switch bit := slice.GetElement(i); {
                case bit == 1:
                    v.body[i/wordSize] ^= (1 << (wordSize - (i % wordSize) - 1))
                case bit != 0:
                    return nil, fmt.Errorf("vector: unexpected digit %d in position %d, possible only {0, 1}",
                        bit, i)
                }
            }
            return v, nil
        }
    }
}

// newFromString converts string to Vector
// Function supports the following filler for zero symbol:
//       '0' == '0', '-'
// Example:
//   "--110-1- - - 1-   0 " -> 0011001000100
func newFromString(s *string) (*Vector, error) {
    if len(*s) == 0 {
        return newEmpty(0), nil
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
            return nil, e
        }
    }
    return v, nil
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
