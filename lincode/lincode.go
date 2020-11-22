package lincode

import "fmt"
import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

//LinearCode represents generic binary linear code
//This type is immutable
type LinearCode struct {
    generator   *matrix.Matrix
    parityCheck *matrix.Matrix
}

//D returns code distance of code
func (code *LinearCode) D() int {
    return D(code)
}

//K return dimension of linear code
func (code *LinearCode) K() int {
    if code.generator != nil {
        if code.generator.Nrows() == 1 && code.generator.GetRow(0).Equal(vector.New(code.generator.Ncolumns())) {
            return 0
        }
        return code.generator.Nrows()
    }
    if code.parityCheck != nil {
        return code.parityCheck.Ncolumns() - code.parityCheck.Nrows()
    }
    return 0
}

//N return length of linear code
func (code *LinearCode) N() int {
    if code.generator != nil {
        return code.generator.Ncolumns()
    }
    if code.parityCheck != nil {
        return code.parityCheck.Ncolumns()
    }
    return 0
}

//Gen returns generator of linear codes
func (code *LinearCode) Gen() *matrix.Matrix {
    if code.generator != nil {
        return code.generator
    }
    if code.parityCheck != nil {
        code.generator = code.parityCheck.Orthogonal()
        return code.generator
    }
    return matrix.New(0)
}

//ParityCheck returns parity check of linear codes
func (code *LinearCode) ParityCheck() *matrix.Matrix {
    if code.parityCheck != nil {
        return code.parityCheck
    }
    if code.generator != nil {
        code.parityCheck = code.generator.Orthogonal()
        return code.parityCheck
    }
    return matrix.New(0)
}

//FromCodeWords returns code defined by list of code words
func FromCodeWords(words interface{}) *LinearCode {
    switch w := words.(type) {
    case *matrix.Matrix:
        m := w.Diagonal()
        v := make([](*vector.Vector), 0, m.Nrows())
        for i := 0; i < m.Nrows(); i++ {
            if r := m.GetRow(i); !r.IsZero() {
                v = append(v, r)
            }
        }
        if len(v) == 0 {
            v = append(v, vector.New(m.Ncolumns()))
        }
        return &LinearCode{generator: matrix.New(v), parityCheck: nil}
    case [](*vector.Vector):
        m := matrix.New(w).Diagonal()
        v := make([](*vector.Vector), 0, m.Nrows())
        for i := 0; i < m.Nrows(); i++ {
            if r := m.GetRow(i); !r.IsZero() {
                v = append(v, r)
            }
        }
        if len(v) == 0 {
            v = append(v, vector.New(m.Ncolumns()))
        }
        return &LinearCode{generator: matrix.New(v), parityCheck: nil}
    default:
        panic(fmt.Errorf("lincode: cannot create code from %T", w))
    }
}

//FromParityChecks returns code defined by list of parity check words
func FromParityChecks(words interface{}) *LinearCode {
    switch w := words.(type) {
    case *matrix.Matrix:
        m := w.Diagonal()
        v := make([](*vector.Vector), 0, m.Nrows())
        for i := 0; i < m.Nrows(); i++ {
            if r := m.GetRow(i); !r.IsZero() {
                v = append(v, r)
            }
        }
        if len(v) == 0 {
            v = append(v, vector.New(m.Ncolumns()))
        }
        return &LinearCode{generator: nil, parityCheck: matrix.New(v)}
    case [](*vector.Vector):
        m := matrix.New(w).Diagonal()
        v := make([](*vector.Vector), 0, m.Nrows())
        for i := 0; i < m.Nrows(); i++ {
            if r := m.GetRow(i); !r.IsZero() {
                v = append(v, r)
            }
        }
        if len(v) == 0 {
            v = append(v, vector.New(m.Ncolumns()))
        }
        return &LinearCode{parityCheck: matrix.New(v), generator: nil}
    default:
        panic(fmt.Errorf("lincode: cannot create code from %T", w))
    }
}
