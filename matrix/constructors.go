package matrix

import "github.com/gf2crypto/blincodes-go/vector"
import "strings"
import "fmt"
import "math/rand"
import "container/list"
import "time"

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
        case *vector.Vector:
            return &Matrix{body: [](*vector.Vector){t}, ncolumns: t.Len()}
        case [](*vector.Vector):
            ncolumns := -1
            for _, v := range t {
                if ncolumns >= 0 && v.Len() != ncolumns {
                    panic(fmt.Errorf("matrix: expected all vectors have the same length"))
                }
                if ncolumns < 0 {
                    ncolumns = v.Len()
                }
            }
            return &Matrix{body: t, ncolumns: ncolumns}
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
        body[i] = vector.New(slice.GetSlice(i*ncolumns, (i+1)*ncolumns))
    }
    return &Matrix{body: body, ncolumns: ncolumns}
}

//Identity creates identity matrix of size n
func Identity(n int) *Matrix {
    if n <= 0 {
        return newEmpty(0, 0)
    }
    row := make([]uint8, n)
    body := make([](*vector.Vector), 0, n)
    for i := 0; i < n; i++ {
        if i > 0 {
            row[i-1] = 0
        }
        row[i] = 1
        body = append(body, vector.New(row))
    }
    return &Matrix{body: body, ncolumns: n}
}

//Random returns random (m x n) - matrix.
func Random(m int, n ...interface{}) *Matrix {
    if m <= 0 {
        return newEmpty(0, 0)
    }
    nc := m
    if len(n) > 0 {
        if t, ok := n[0].(int); ok {
            nc = t
        } else {
            panic(fmt.Errorf("matrix: type error, expected number of columns is integer, not %T", n[0]))
        }
    }
    if nc <= 0 {
        return newEmpty(0, 0)
    }
    body := make([](*vector.Vector), 0, m)
    zero := vector.New(m)
    for i := 0; i < m; i++ {
        v := vector.Random(nc)
        for v.Equal(zero) {
            v = vector.Random(nc)
        }
        body = append(body, v)
    }
    return &Matrix{body: body, ncolumns: nc}
}

//RandomMaxRank returns random (m x n) - matrix of maximal rank.
//     nonsing = nonsingular(min(nrows, ncolumns))
//     perm_matrix = permutation(sample(range(max(nrows, ncolumns)),
//                                      max(nrows, ncolumns)))
//     if nrows < ncolumns:
//         return concatenate(nonsing, random(ncolumns - nrows)) * perm_matrix
//     return perm_matrix * concatenate(nonsing,
//                                      random(nrows - ncolumns),
//                                      by_rows=True)
func RandomMaxRank(m int, n ...interface{}) *Matrix {
    if m <= 0 {
        return newEmpty(0, 0)
    }
    nc := m
    if len(n) > 0 {
        if t, ok := n[0].(int); ok {
            nc = t
        } else {
            panic(fmt.Errorf("matrix: type error, expected number of columns is integer, not %T", n[0]))
        }
    }
    if nc <= 0 {
        return newEmpty(0, 0)
    }
    if nc == m {
        return Nonsing(m)
    }
    if m > nc {
        a := Nonsing(nc).ConcatenateRows(Random(m-nc, nc))
        return PermLeft(rand.Perm(m)).Mul(a)
    }
    // m < nc
    a := Nonsing(m).ConcatenateColumns(Random(m, nc-m))
    return a.Mul(Perm(rand.Perm(nc)))
}

//Nonsing returns nonsingular random matrix
// Function uses algorithm of Dana Randall
// https://www.researchgate.net/publication/2729950_Efficient_Generation_of_Random_Nonsingular_Matrices
func Nonsing(n int) *Matrix {
    if n <= 0 {
        return newEmpty(0, 0)
    }
    rand.Seed(time.Now().UnixNano())
    matA := make([]([]uint8), n)
    for i := 0; i < n; i++ {
        matA[i] = make([]uint8, n)
    }
    matT := make([](*vector.Vector), n)
    cols := list.New()
    for i := 0; i < n; i++ {
        cols.PushBack(i)
    }
    for i := 0; i < n; i++ {
        //Generate random v != 0
        // And found its the first 1
        v := make([]uint8, n)
        r := cols.Front()
        for isZero := true; isZero; {
            for e := cols.Front(); e != nil; e = e.Next() {
                v[e.Value.(int)] = uint8(rand.Intn(2))
                if isZero && v[e.Value.(int)] != 0 {
                    isZero = false
                    r = e // index of the first 1
                }
            }
        }
        //Update matrix A
        matA[i][r.Value.(int)] = 1
        // There is a mistake in the paper of Dana Randall: this code was missing in her paper
        for j := i + 1; j < n; j++ {
            matA[j][r.Value.(int)] = uint8(rand.Intn(2))
        }
        //Update matrix T
        a := make([]uint8, n)
        for e := cols.Front(); e != nil; e = e.Next() {
            a[e.Value.(int)] = v[e.Value.(int)]
        }
        matT[r.Value.(int)] = vector.New(a)
        cols.Remove(r)
    }
    bodyA := make([](*vector.Vector), n)
    for i := 0; i < n; i++ {
        bodyA[i] = vector.New(matA[i])
    }
    return (&Matrix{body: bodyA, ncolumns: n}).Mul(&Matrix{body: matT, ncolumns: n})

}

//PermLeft returns permutation matrix corresponding to the left-action permutation
func PermLeft(p []int) *Matrix {
    if len(p) == 0 {
        return newEmpty(0, 0)
    }
    body := make([](*vector.Vector), 0, len(p))
    for _, j := range p {
        v := make([]uint8, len(p))
        if j >= 0 && j < len(p) {
            v[j] = 1
        }
        body = append(body, vector.New(v))
    }
    return &Matrix{body: body, ncolumns: len(p)}
}

//Perm returns permutation matrix corresponding to the right-action permutation
func Perm(p []int) *Matrix {
    if len(p) == 0 {
        return newEmpty(0, 0)
    }
    tmp := make([][]uint8, len(p))
    for i := 0; i < len(p); i++ {
        tmp[i] = make([]uint8, len(p))
    }
    for i, j := range p {
        if j >= 0 && j < len(p) {
            tmp[j][i] = 1
        }
    }
    body := make([](*vector.Vector), 0, len(p))
    for i := 0; i < len(p); i++ {
        body = append(body, vector.New(tmp[i]))
    }
    return &Matrix{body: body, ncolumns: len(p)}
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
        body[i] = vector.New(r)
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
        body[i] = vector.New(n)
    }
    return &Matrix{body: body, ncolumns: n}
}
