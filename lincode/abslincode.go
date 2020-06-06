package lincode

import "sort"
import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

//AbstractLinearCode defines the abstract interface of linear block code
type AbstractLinearCode interface {
    Gen() *matrix.Matrix
    ParityCheck() *matrix.Matrix
    N() int
    K() int
    D() int
}

//HadamardProd evaluates the Hadamard product of linear codes
func HadamardProd(codeA, codeB AbstractLinearCode) AbstractLinearCode {
    genA, genB := codeA.Gen(), codeB.Gen()
    b := make(map[int](*vector.Vector))
    hmbasis := make(map[int](*vector.Vector))
    for i := 0; i < codeA.K(); i++ {
        for j := 0; j < codeB.K(); j++ {
            row := genA.GetRow(i).And(genB.GetRow(j))
            sum := row.Copy()
            for i, r := range b {
                if row.Get(i) != 0 {
                    sum = sum.Xor(r)
                }
            }
            t := sum.FirstOne()
            if t != sum.Len() {
                b[t] = sum
                hmbasis[len(hmbasis)] = row
            }
        }
    }
    body := make([](*vector.Vector), 0, len(hmbasis))
    for _, r := range hmbasis {
        body = append(body, r)
    }
    return FromCodeWords(body)
}

//Intersect intersects of codes
func Intersect(codeA, codeB AbstractLinearCode) AbstractLinearCode {
    return FromParityChecks(codeA.ParityCheck().ConcatenateRows(codeB.ParityCheck()))
}

//Sum evaluates sum of codes
func Sum(codeA, codeB AbstractLinearCode) AbstractLinearCode {
    return FromCodeWords(codeA.Gen().ConcatenateRows(codeB.Gen()))
}

//Hull evaluates hull of code
//The code's hull is intersection of code and it's dual.
func Hull(code AbstractLinearCode) AbstractLinearCode {
    return FromParityChecks(code.Gen().ConcatenateRows(code.ParityCheck()))
}

//Puncture evaluates of puncture code.
// Punctured code is code obtaining by set the positions
// with indexes from `ncolumns` of every codeword to zero.
// Punctured code is NOT subcode of original code!
func Puncture(code AbstractLinearCode, columns []int) AbstractLinearCode {
    gen := code.Gen()
    sup := make([]int, 0, gen.Ncolumns())
    ones := append(make([]int, 0, len(columns)), columns...)
    sort.Slice(ones, func(i, j int) bool { return ones[i] < ones[j] })
    for i := 0; i < code.N(); i++ {
        if j := sort.SearchInts(ones, i); j >= len(ones) || ones[j] != i {
            sup = append(sup, i)
        }
    }
    mask := vector.New(code.N(), sup)
    body := make([](*vector.Vector), 0, code.K())
    for i := 0; i < code.K(); i++ {
        body = append(body, gen.GetRow(i).And(mask))
    }
    return FromCodeWords(body)
}

//Truncate evaluates of truncated code.
//Truncated code is code obtaining by choose codewords which
//have coordinates with indexes from `columns` is zero.
//Unlike the punctured code truncated code is a subcode of original code.
func Truncate(code AbstractLinearCode, columns []int) AbstractLinearCode {
    gen, _ := code.Gen().GaussElim(false, columns)
    body := make([](*vector.Vector), 0, code.K())
    for i := 0; i < code.K(); i++ {
        flag := true
        for _, j := range columns {
            if j < code.N() && j >= 0 && gen.GetRow(i).Get(j) != 0 {
                flag = false
                break
            }
        }
        if flag {
            body = append(body, gen.GetRow(i))
        }
    }
    return FromCodeWords(body)
}

//IsSubset tests if codeA is subset of codeB or not
func IsSubset(codeA, codeB AbstractLinearCode) bool {
    if codeA.N() != codeB.N() || codeA.K() > codeB.K() {
        return false
    }
    if codeB.ParityCheck().Mul(codeA.Gen().T()).Equal(matrix.New(codeB.N()-codeB.K(), codeA.K())) {
        return true
    }
    return false
}

//Encode encodes the message
func Encode(code AbstractLinearCode, mes *vector.Vector) *vector.Vector {
    m := matrix.New(mes)
    return m.Mul(code.Gen()).GetRow(0)
}

//Syndrome evaluates syndrome of the vector
func Syndrome(code AbstractLinearCode, vec *vector.Vector) *vector.Vector {
    x := matrix.New(vec)
    return code.ParityCheck().Mul(x.T()).T().GetRow(0)
}

//Iter iterates over code words
func Iter(code AbstractLinearCode) <-chan *vector.Vector {
    ch := make(chan *vector.Vector)
    go func() {
        defer close(ch)
        v := make([]int, code.K())
        if len(v) == 0 {
            return
        }
        carry := 0
        for carry == 0 {
            ch <- matrix.New(vector.New(v)).Mul(code.Gen()).GetRow(0)
            for i := 0; i < len(v); i++ {
                v[len(v)-i-1], carry = (v[len(v)-i-1]+1+carry)%2, ((v[len(v)-i-1] + 1 + carry) >> 1)
                if carry == 0 {
                    break
                }
            }
        }
    }()
    return ch
}

//Spectrum returns weight spectrum of code
func Spectrum(code AbstractLinearCode) *map[int]int {
    spec := make(map[int]int)
    for v := range Iter(code) {
        spec[v.Wt()]++
    }
    return &spec
}

//D returns code distance of code
func D(code AbstractLinearCode) int {
    if d := code.D(); d != -1 {
        return d
    }
    sp := Spectrum(code)
    for d := range *sp {
        if d != 0 {
            return d
        }
    }
    return 0
}
