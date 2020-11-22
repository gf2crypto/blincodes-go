package lincode

import "testing"
import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

var matrices = map[string](*matrix.Matrix){
    "rm14": matrix.New(5, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
    }),
    "rm14_add": matrix.New(9, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
    }),
    "rm24": matrix.New(11, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
        0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
        0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
    }),
    "rm24_add": matrix.New(31, []uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
        0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
        0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
        0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0,
        0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
    }),
}

//TestFromCodeWords tests creation the new code from its code words.
func TestFromCodeWords(t *testing.T) {
    if rm14 := FromCodeWords(matrices["rm14_add"]); !rm14.Gen().Mul(matrices["rm24"].T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm14.G is wrong:\n%v", rm14.Gen())
    }
    if rm14 := FromCodeWords(matrices["rm14_add"]); !rm14.Gen().Mul(rm14.ParityCheck().T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm14.H is wrong:\n%v", rm14.ParityCheck())
    }
    body := make([](*vector.Vector), 0, matrices["rm14_add"].Nrows())
    for i := 0; i < matrices["rm14_add"].Nrows(); i++ {
        body = append(body, matrices["rm14_add"].GetRow(i))
    }
    if rm14 := FromCodeWords(body); !rm14.Gen().Mul(matrices["rm24"].T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm14.G is wrong:\n%v", rm14.Gen())
    }
    if rm14 := FromCodeWords(body); !rm14.Gen().Mul(rm14.ParityCheck().T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm14.H is wrong:\n%v", rm14.ParityCheck())
    }
}

//TestFromParityChecks tests creation the new code from its parity checks.
func TestFromParityChecks(t *testing.T) {
    if rm24 := FromParityChecks(matrices["rm24_add"]); !rm24.Gen().Mul(matrices["rm24"].T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm24.G is wrong:\n%v", rm24.Gen())
    }
    if rm24 := FromParityChecks(matrices["rm24_add"]); !rm24.Gen().Mul(rm24.ParityCheck().T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm24.H is wrong:\n%v", rm24.ParityCheck())
    }
    body := make([](*vector.Vector), 0, matrices["rm24_add"].Nrows())
    for i := 0; i < matrices["rm24_add"].Nrows(); i++ {
        body = append(body, matrices["rm24_add"].GetRow(i))
    }
    if rm24 := FromParityChecks(body); !rm24.Gen().Mul(matrices["rm24"].T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm24.G is wrong:\n%v", rm24.Gen())
    }
    if rm24 := FromParityChecks(body); !rm24.Gen().Mul(rm24.ParityCheck().T()).Equal(matrix.New(5, 11)) {
        t.Errorf("rm24.H is wrong:\n%v", rm24.ParityCheck())
    }
}

//TestLinCodeD tests to evaluate of code distance.
func TestLinCodeD(t *testing.T) {
    if d := FromCodeWords(matrices["rm14"]).D(); d != 8 {
        t.Errorf("D(RM(1,4)) is wrong: expected 8, but got %v", d)
    }
    if d := FromCodeWords(matrices["rm24"]).D(); d != 4 {
        t.Errorf("D(RM(2,4)) is wrong: expected 4, but got %v\n", d)
    }
}
