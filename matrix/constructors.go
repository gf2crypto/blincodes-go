package matrix

import "github.com/gf2crypto/blincodes-go/vector"
import "strings"
import "fmt"

// New creates Matrix object
// Parameters:
//     nrows int - number of rows of the Matrix
//     param - int, string, []string, []int{,8,16,32,64}, []uint{, 8, 16, 32, 64}
func New(params ...interface{}) *Matrix {
    if len(params) == 0 {
        return newEmpty(0)
    }
    if len(params) == 1 {
        switch t := params[0].(type) {
        case int:
            return newEmpty(int(t))
        case int8:
            return newEmpty(int(t))
        case int16:
            return newEmpty(int(t))
        case int32:
            return newEmpty(int(t))
        case int64:
            return newEmpty(int(t))
        case uint:
            return newEmpty(int(t))
        case uint8:
            return newEmpty(int(t))
        case uint16:
            return newEmpty(int(t))
        case uint32:
            return newEmpty(int(t))
        case uint64:
            return newEmpty(int(t))
        case []string:
            return newFromStrings(t)
        case string:
            return newFromStrings(strings.Split(t, "\n"))
        default:
            panic(fmt.Errorf("matrix: cannot convert %T to Matrix", t))
        }
    }
    if len(params) > 2 {
        e := fmt.Errorf("matrix: expected number of parameters less or equal 2, but got %v > 2", len(params))
        panic(e)
    }
    // len(params) == 2
    var nrows int
    switch p := params[0].(type) {
    case int:
        nrows = int(p)
    case int8:
        nrows = int(p)
    case int16:
        nrows = int(p)
    case int32:
        nrows = int(p)
    case int64:
        nrows = int(p)
    case uint:
        nrows = int(p)
    case uint8:
        nrows = int(p)
    case uint16:
        nrows = int(p)
    case uint32:
        nrows = int(p)
    case uint64:
        nrows = int(p)
    default:
        e := fmt.Errorf("matrix: expected type of nrows integer, type %T is not supported", nrows)
        panic(e)
    }
    if nrows == 0 {
        return newEmpty(0)
    }
    switch b := params[1].(type) {
    case nil:
        return newEmpty(nrows)
    case []uint:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []uint8:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []uint16:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []uint32:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []uint64:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []int:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []int8:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []int16:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []int32:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    case []int64:
        {
            if len(b)%nrows != 0 {
                e := fmt.Errorf("matrix: cannot create Matrix with nrows from %T, len(%v) %% %v", b, b, nrows)
                panic(e)
            }
            ncolumns := len(b) % nrows
            body := make([](*vector.Vector), nrows)
            for i := 0; i < nrows; i++ {
                body[i], _ = vector.New(b[i*ncolumns : (i+1)*ncolumns])
            }
            return &Matrix{body: body, ncolumns: ncolumns, nrows: nrows}
        }
    default:
        e := fmt.Errorf("matrix: unsupported type %T, cannot create matrix form (%v, %T)", b, nrows, b)
        panic(e)
    }
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
        return &Matrix{body: make([](*vector.Vector), 0), nrows: 0, ncolumns: 0}
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
    return &Matrix{body: body, nrows: len(body), ncolumns: body[0].Len()}
}

func newEmpty(n int) *Matrix {
    if n <= 0 {
        return &Matrix{body: make([](*vector.Vector), 0, 0), nrows: 0, ncolumns: 0}
    }
    body := make([](*vector.Vector), n)
    for i := range body {
        body[i], _ = vector.New(n)
    }
    return &Matrix{body: body, nrows: n, ncolumns: n}
}
