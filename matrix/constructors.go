package matrix

import "github.com/gf2crypto/blincodes-go/vector"
import "strings"
import "fmt"

// New creates Matrix object
// Parameters:
//     New() creates empty matrix
//     New(n (u)int{, 8, 16, 32}) create (n x n) zero matrix
//     New(m, n (u)int{, 8, 16, 32}) creates (m x n) zero matrix
//     New(s string) creates Matrix from string, uses row-separator \n
//          Example:             Result:
//              ` 11111111             11111111
//                ----1111             00001111
//                00110011             00110011
//                01-101-1`            01010101
//     New(s []string) creates Matrix from string array, every element of this array is row
//           Example:                   Result:
//           []string{
//              "11111111",                11111111
//              "- - --1 111",             00001111
//              "00110011",                00110011
//              "01-10   1 -1",            01010101
//           }
//     New(m, (u)int{, 8, 16, 32}) creates Matrix with m rows from integer array, len of this array must be divided to m
func New(params ...interface{}) *Matrix {
    if len(params) == 0 {
        return newEmpty(0, 0)
    }
    if len(params) == 1 {
        switch t := params[0].(type) {
        case nil:
            return newEmpty(0, 0)
        case []string:
            return newFromStrings(t)
        case string:
            return newFromStrings(strings.Split(t, "\n"))
        default:
            var n int
            var e error
            if n, e = toInt(t); e != nil {
                panic(e)
            }
            return newEmpty(n, n)
        }
    }
    if len(params) > 2 {
        panic(fmt.Errorf("matrix: expected number of parameters less or equal 2, but got %v > 2", len(params)))
    }
    // len(params) == 2
    var nrows int
    var e error
    nrows, e = toInt(params[0])
    if e != nil {
        panic(fmt.Errorf("matrix: expected type of nrows integer, type %T is not supported", nrows))
    }
    if nrows <= 0 {
        return newEmpty(0, 0)
    }
    if params[1] == nil {
        return newEmpty(nrows, nrows)
    }
    var ncolumns int
    ncolumns, e = toInt(params[1])
    if e == nil {
        return newEmpty(nrows, ncolumns)
    }
    var slice vector.Slicer
    slice, e = vector.ToSlicer(params[1])
    if e != nil {
        panic(fmt.Errorf("matrix: unsupported type %T, cannot create matrix form (%v, %T)", params[1], nrows, params[1]))
    }
    if slice.Len()%nrows != 0 {
        panic(fmt.Errorf("matrix: cannot create Matrix with nrows, len of input array must divided to nrows, but %v mod %v = %v != 0", slice.Len(), nrows, slice.Len()%nrows))
    }
    ncolumns = slice.Len() / nrows
    body := make([](*vector.Vector), nrows)
    for i := 0; i < nrows; i++ {
        body[i], _ = vector.New(slice.GetSlice(i*ncolumns, (i+1)*ncolumns))
    }
    return &Matrix{body: body, ncolumns: ncolumns}
}

// newFromStrings converts array string to Matrix
// Function supports the following filler for zero symbol:
//       '0' == '0', '-'
func newFromStrings(s []string) *Matrix {
    nrows := 0
    rows := make([]string, 0, len(s))
    for _, v := range s {
        if v != "" {
            rows = append(rows, v)
            nrows++
        }
    }
    if len(rows) == 0 {
        return &Matrix{body: make([](*vector.Vector), 0), ncolumns: 0}
    }
    body := make([](*vector.Vector), nrows)
    for i, r := range rows {
        var e error
        body[i], e = vector.New(r)
        if e != nil {
            panic(fmt.Errorf("matrix: cannot convert %v to vector, error: %v", r, e))
        }
        if (i > 0) && (body[i].Len() != body[0].Len()) {
            panic(fmt.Errorf("matrix: expected the same length of rows, but len(Matrix[%v]=%v) != len(Matrix[0]=%v): %v != %v",
                i, body[i], body[0], body[i].Len(), body[0].Len()))
        }
    }
    return &Matrix{body: body, ncolumns: body[0].Len()}
}

func newEmpty(m, n int) *Matrix {
    if m <= 0 || n <= 0 {
        return &Matrix{body: make([](*vector.Vector), 0, 0), ncolumns: 0}
    }
    body := make([](*vector.Vector), m)
    for i := range body {
        body[i], _ = vector.New(n)
    }
    return &Matrix{body: body, ncolumns: n}
}
