package lincode

// import "fmt"
import "testing"
import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

//TestHadamardProd tests evaluation of Hadamard product of codes.
func TestHadamardProd(t *testing.T) {
    if codeA := FromCodeWords(matrices["rm14"]); !HadamardProd(codeA, codeA).Gen().Mul(matrices["rm14"].T()).Equal(matrix.New(11, 5)) {
        t.Errorf("HadamardProduct(RM(1, 4), RM(1, 4)) is wrong:\n%v", HadamardProd(codeA, codeA))
    }

    if codeA := FromCodeWords(matrices["rm24"]); !HadamardProd(codeA, codeA).Gen().Equal(matrix.Identity(16)) {
        t.Errorf("HadamardProduct(RM(2, 4), RM(2, 4)) is wrong:\n%v", HadamardProd(codeA, codeA))
    }

    if codeA, codeB := FromCodeWords(matrices["rm24"]), FromCodeWords(matrices["rm14"]); !HadamardProd(codeA, codeB).ParityCheck().Equal(matrix.New(1, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })) {
        t.Errorf("HadamardProduct(RM(2, 4), RM(1, 4)) is wrong:\n%v", HadamardProd(codeA, codeB))
    }
}

//TestIntersect tests evaluation of code intersection.
func TestIntersect(t *testing.T) {
    if codeA, codeB := FromCodeWords(matrices["rm14_add"]), FromCodeWords(matrices["rm24_add"]); !IsEqual(Intersect(codeA, codeB), codeA) {
        t.Errorf("Intersect(RM(1, 4) , RM(2,4)) is wrong:\n%v", Intersect(codeA, codeB))
    }
    if codeA, codeB := FromCodeWords(matrices["rm24_add"]), FromCodeWords(matrices["rm24_add"]); !IsEqual(Intersect(codeA, codeB), codeA) {
        t.Errorf("Intersect(RM(2, 4) , RM(2,4)) is wrong:\n%v", Intersect(codeA, codeB))
    }
}

//TestSum tests evaluation of code sum.
func TestSum(t *testing.T) {
    if codeA, codeB := FromCodeWords(matrices["rm14_add"]), FromCodeWords(matrices["rm24_add"]); !IsEqual(Sum(codeA, codeB), codeB) {
        t.Errorf("Sum(RM(1, 4) , RM(2,4)) is wrong:\n%v", Sum(codeA, codeB))
    }
    if codeA, codeB := FromCodeWords(matrices["rm24_add"]), FromCodeWords(matrices["rm24_add"]); !IsEqual(Sum(codeA, codeB), codeB) {
        t.Errorf("Sum(RM(2, 4) , RM(2,4)) is wrong:\n%v", Sum(codeA, codeB))
    }
}

//TestHull tests evaluation of code's hull.
func TestHull(t *testing.T) {
    if codeA := FromCodeWords(matrices["rm14_add"]); !IsEqual(Hull(codeA), codeA) {
        t.Errorf("Hull(RM(1, 4)) is wrong:\n%v", Hull(codeA))
    }
    if codeA, codeB := FromCodeWords(matrices["rm24_add"]), FromCodeWords(matrices["rm14_add"]); !IsEqual(Hull(codeA), codeB) {
        t.Errorf("Hull(RM(2, 4)) is wrong:\n%v", Hull(codeA))
    }
}

//TestPuncture tests evaluation of punctured code.
func TestPuncture(t *testing.T) {
    punc := FromCodeWords(matrix.New(5, []uint8{
        0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0,
        0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0,
    }))
    if codeA := Puncture(FromCodeWords(matrices["rm14_add"]), []int{0, 4, 8, 9, 15}); !IsEqual(codeA, punc) {
        t.Errorf("Puncture(RM(1, 4), [0, 4, 8, 9, 15]) is wrong:\n%v", codeA)
    }
    punc = FromCodeWords(matrix.New(4, []uint8{
        0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0,
    }))
    if codeA := Puncture(FromCodeWords(matrices["rm14_add"]), []int{0, 1, 2, 3, 4, 5, 8, 9, 14, 15}); !IsEqual(codeA, punc) {
        t.Errorf("Puncture(RM(1, 4), [0, 1, 2, 3, 4, 5, 8, 9, 14, 15]) is wrong:\n%v", codeA)
    }
}

//TestTruncate tests evaluation of truncated code.
func TestTruncate(t *testing.T) {
    trunc := FromCodeWords(matrix.New(2, []uint8{
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    }))
    if codeA := Truncate(FromCodeWords(matrices["rm14_add"]), []int{0, 4, 8, 9, 15}); !IsEqual(codeA, FromCodeWords([](*vector.Vector){vector.New(16)})) {
        t.Errorf("Truncate(RM(1, 4), [0, 4, 8, 9, 15]) is wrong:\n%v", codeA.Gen())
    }
    if codeA := Truncate(FromCodeWords(matrices["rm14_add"]), []int{0, 1, 2, 3}); !IsEqual(codeA, trunc) {
        t.Errorf("Truncate(RM(1, 4), [0, 1, 2, 3]) is wrong:\n%v", codeA.Gen())
    }
}

//TestEncode tests to encode of message.
func TestEncode(t *testing.T) {
    if rm, v := FromCodeWords(matrices["rm14"]), vector.New([]byte{1, 1, 1, 1, 1}); !Encode(rm, v).Equal(vector.New([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})) {
        t.Errorf("Encode(rm14, %v) is wrong: %v, expected %v", v, Encode(rm, v), vector.New([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
    }
}

//TestSyndrome tests to evaluate of syndrome.
func TestSyndrome(t *testing.T) {
    if rm, v := FromCodeWords(matrices["rm14"]), vector.New([]byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}); !Syndrome(rm, v).Equal(vector.New([]byte{1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0})) {
        t.Errorf("Syndrome(rm14, %v) is wrong: %v, expected %v", v, Syndrome(rm, v), vector.New([]byte{1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0}))
    }
}

//TestIter tests to iterate over code words.
func TestIter(t *testing.T) {
    cws := [](*vector.Vector){
        vector.New([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
        vector.New([]byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}),
        vector.New([]byte{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}),
        vector.New([]byte{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0}),
        vector.New([]byte{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1}),
        vector.New([]byte{0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0}),
        vector.New([]byte{0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0}),
        vector.New([]byte{0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1}),
        vector.New([]byte{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}),
        vector.New([]byte{0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0}),
        vector.New([]byte{0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0}),
        vector.New([]byte{0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1}),
        vector.New([]byte{0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0}),
        vector.New([]byte{0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1}),
        vector.New([]byte{0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1}),
        vector.New([]byte{0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0}),
        vector.New([]byte{1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1}),
        vector.New([]byte{1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0}),
        vector.New([]byte{1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0}),
        vector.New([]byte{1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1}),
        vector.New([]byte{1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0}),
        vector.New([]byte{1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1}),
        vector.New([]byte{1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1}),
        vector.New([]byte{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}),
        vector.New([]byte{1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0}),
        vector.New([]byte{1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1}),
        vector.New([]byte{1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1}),
        vector.New([]byte{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0}),
        vector.New([]byte{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}),
        vector.New([]byte{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}),
        vector.New([]byte{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}),
        vector.New([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
    }
    i := 0
    for c := range Iter(FromCodeWords(matrices["rm14"])) {
        if !c.Equal(cws[i]) {
            t.Errorf("Iter() is wrong: expected c[%v] is %v, but got c[%v] = %v", i, cws[i], i, c)
        }
        i++
    }
}

//TestSpectrum tests to evaluate of code spectrum.
func TestSpectrum(t *testing.T) {
    sp := Spectrum(FromCodeWords(matrices["rm14"]))
    for w, n := range *sp {
        switch {
        case w == 0:
            if n != 1 {
                t.Errorf("Spectrum is wrong: expected sp[%v] is %v, but got sp[%v] = %v", w, 1, w, n)
            }
        case w == 16:
            if n != 1 {
                t.Errorf("Spectrum is wrong: expected sp[%v] is %v, but got sp[%v] = %v", w, 1, w, n)
            }
        case w == 8:
            if n != 30 {
                t.Errorf("Spectrum is wrong: expected sp[%v] is %v, but got sp[%v] = %v", w, 30, w, n)
            }
        default:
            if n != 0 {
                t.Errorf("Spectrum is wrong: expected sp[%v] is %v, but got sp[%v] = %v", w, 0, w, n)
            }
        }
    }
}

//TestD tests to evaluate of code distance.
func TestD(t *testing.T) {
    if d := D(FromCodeWords(matrices["rm14"])); d != 8 {
        t.Errorf("D(RM(1,4)) is wrong: expected 8, but got %v", d)
    }
    if d := D(FromCodeWords(matrices["rm24"])); d != 4 {
        t.Errorf("D(RM(2,4)) is wrong: expected 4, but got %v\n%v", d, *Spectrum(FromCodeWords(matrices["rm24"])))
    }
}
