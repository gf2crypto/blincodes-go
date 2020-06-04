package lincode

import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

//LinearCode defines the basic interface of linear block code
type LinearCode interface {
    Gen() *matrix.Matrix
    ParityCheck() *matrix.Matrix
}

//HadamardProd evaluates the Hadamard product of linear codes
func HadamardProd(codeA, codeB *LinearCode) *LinearCode {
    genA, genB := codeA.Gen(), codeB.Gen()
    b := make(map[int](*vector.Vector))
    hmbasis := make(map[int](*vector.Vector))
    for i := 0; i < genA.Nrows(); i++ {
        for j := 0; j < genB.Nrows(); j++ {
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
func Intersect(codeA, codeB *LinearCode) *LinearCode {
    return FromParityChecks(codeA.ParityCheck().ConcatenateRows(codeB.ParityCheck()))
}

//Sum evaluates sum of codes
func Sum(codeA, codeB *LinearCode) *LinearCode {
    return FromCodeWords(codeA.Gen().ConcatenateRows(codeB.Gen()))
}

//Hull evaluates hull of code
//The code's hull is intersection of code and it's dual.
func Hull(code *LinearCode) *LinearCode {
    return FromParityChecks(code.Gen().ConcatenateRows(code.ParityCheck()))
}
