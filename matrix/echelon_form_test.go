package matrix

import "testing"

// import "fmt"

//TestEvaluatiotEchelonForm test the evaluation of Matrix echelon form
func TestEvaluatiotEchelonForm(t *testing.T) {
    mat, echelon := newUpper()
    if !mat.Equal(mat.Echelon()) {
        t.Errorf("%v != %v", mat.Echelon(), echelon)
    }
    mat, echelon = newMaxRank()
    if !echelon.Equal(mat.Echelon()) {
        t.Errorf("%v != %v", mat.Echelon(), echelon)
    }
    mat, echelon = newMaxRankSquare()
    if !echelon.Equal(mat.Echelon()) {
        t.Errorf("%v != %v", mat.Echelon(), echelon)
    }
    mat, echelon = newMaxRankLong()
    if !echelon.Equal(mat.Echelon()) {
        t.Errorf("%v != %v", mat.Echelon(), echelon)
    }
}

func newUpper() (*Matrix, *Matrix) {
    body := []uint8{
        1, 1, 1, 1,
        0, 1, 1, 1,
        0, 0, 1, 1,
        0, 0, 0, 1,
    }
    return New(4, body), New(4, body)
}

func newMaxRank() (*Matrix, *Matrix) {
    body := []uint8{
        0, 1, 1, 1,
        1, 0, 0, 0,
        1, 1, 0, 0,
        1, 1, 1, 0,
    }
    echelon := []uint8{
        1, 0, 0, 0,
        0, 1, 1, 1,
        0, 0, 1, 1,
        0, 0, 0, 1,
    }
    return New(4, body), New(4, echelon)
}

func newMaxRankSquare() (*Matrix, *Matrix) {
    body := []uint8{
        0, 1, 1, 1, 0,
        0, 0, 1, 0, 1,
        1, 1, 0, 0, 1,
        1, 1, 1, 0, 0,
    }
    echelon := []uint8{
        1, 0, 0, 1, 0,
        0, 1, 1, 1, 0,
        0, 0, 1, 0, 1,
        0, 0, 0, 0, 0,
    }
    return New(4, body), New(4, echelon)
}

func newMaxRankLong() (*Matrix, *Matrix) {
    body := []uint8{
        0, 1, 1, 1, 0,
        0, 0, 1, 0, 1,
        1, 1, 0, 0, 1,
        1, 1, 1, 0, 0,
        1, 0, 0, 1, 0,
        1, 1, 1, 1, 1,
        0, 1, 0, 1, 0,
    }
    echelon := []uint8{
        1, 0, 0, 1, 0,
        0, 1, 1, 1, 0,
        0, 0, 1, 0, 1,
        0, 0, 0, 1, 1,
        0, 0, 0, 0, 1,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
    }
    return New(7, body), New(7, echelon)
}
