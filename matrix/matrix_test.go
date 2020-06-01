package matrix

import "testing"

var matricies = [](func() (*Matrix, *Matrix, *Matrix, int)){
    newZero,
    newUpper,
    newRank1,
    newMaxRank,
    newMaxRankLong,
    newNonMaxRankSquare,
    newNonMaxRankLong,
}

var randomArgs = []([2]int){
    [2]int{10, 10},
    [2]int{100, 100},
    [2]int{10, 20},
    [2]int{100, 200},
    [2]int{20, 10},
    [2]int{200, 100},
}

//TestEvaluatiotEchelonForm test the evaluation of Matrix echelon form
func TestEvaluatiotEchelonForm(t *testing.T) {
    for _, m := range matricies {
        mat, echelon, _, _ := m()
        if !echelon.Equal(mat.Echelon()) {
            t.Errorf("%v != %v", mat.Echelon(), echelon)
        }
    }
}

//TestEvaluatiotDiagonalForm test the evaluation of Matrix diagonal form
func TestEvaluatiotDiagonalForm(t *testing.T) {
    for _, m := range matricies {
        mat, _, diag, _ := m()
        if !diag.Equal(mat.Diagonal()) {
            t.Errorf("%v != %v", mat.Diagonal(), diag)
        }
    }
}

//TestEvaluatiotRank test the evaluation of Matrix's rank
func TestEvaluatiotRank(t *testing.T) {
    for _, m := range matricies {
        mat, _, _, rank := m()
        if mat.Rank() != rank {
            t.Errorf("%v.Rank() = %v, but got %v", mat, mat.Rank(), rank)
        }
    }
}

//TestOrthogonal test the evaluation of Matrix's rank
func TestOrthogonal(t *testing.T) {
    for _, m := range matricies {
        mat, _, _, rank := m()
        orth := mat.Orthogonal()
        orthT := orth.T()
        mul := mat.Mul(orthT)
        res := New(mat.Nrows(), mat.Ncolumns()-rank)
        if mat.Ncolumns() == rank {
            res = New(mat.Nrows(), 1)
        }
        eq := mul.Equal(res)
        if !eq {
            t.Errorf("GH^T != 0, mat:\n%v,\northogonal:\n%v,\northogonal^T:\n%v,\nbut GH^T=\n%v\nexpected=\n%v",
                mat, orth, orthT, mul, res)
        }
    }
}

//TestRandomOrthogonal test the evaluation of orthogonal matrix for random matrix
func TestRandomOrthogonal(t *testing.T) {
    for _, arg := range randomArgs {
        matArray, _ := makeRandomMatrix(arg[0], arg[1])
        mat := New(arg[0], matArray)
        rank := mat.Rank()
        orth := mat.Orthogonal()
        orthT := orth.T()
        mul := mat.Mul(orthT)
        res := New(mat.Nrows(), mat.Ncolumns()-rank)
        if mat.Ncolumns() == rank {
            res = New(mat.Nrows(), 1)
        }
        eq := mul.Equal(res)
        if !eq {
            t.Errorf("GH^T != 0, mat:\n%v,\northogonal:\n%v,\northogonal^T:\n%v,\nbut GH^T=\n%v\nexpected=\n%v",
                mat, orth, orthT, mul, res)
        }
    }
}

//TestInv test the evaluation of generalized inverse of Matrix
func TestInv(t *testing.T) {
    for _, m := range matricies {
        mat, _, _, _ := m()
        inv := mat.Inv()
        mul := inv.Mul(mat)
        res := mat.Diagonal()
        eq := mul.Equal(res)
        if !eq {
            t.Errorf("G^{-1}*G != E, mat:\n%v,\nG^{-1}:\n%v,\nbut G^{-1}*G=\n%v\nexpected=\n%v",
                mat, inv, mul, res)
        }
    }
}

//TestRandomInv test the evaluation of generalized inverse for random matrix
func TestRandomInv(t *testing.T) {
    for _, arg := range randomArgs {
        matArray, _ := makeRandomMatrix(arg[0], arg[1])
        mat := New(arg[0], matArray)
        inv := mat.Inv()
        mul := inv.Mul(mat)
        res := mat.Diagonal()
        eq := mul.Equal(res)
        if !eq {
            t.Errorf("G^{-1}*G != E, mat:\n%v,\nG^{-1}:\n%v,\nbut G^{-1}*G=\n%v\nexpected=\n%v",
                mat, inv, mul, res)
        }
    }
}

func newUpper() (*Matrix, *Matrix, *Matrix, int) {
    body := []uint8{
        1, 1, 1, 1,
        0, 1, 1, 1,
        0, 0, 1, 1,
        0, 0, 0, 1,
    }
    diag := []uint8{
        1, 0, 0, 0,
        0, 1, 0, 0,
        0, 0, 1, 0,
        0, 0, 0, 1,
    }
    return New(4, body), New(4, body), New(4, diag), 4
}

func newMaxRank() (*Matrix, *Matrix, *Matrix, int) {
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
    diag := []uint8{
        1, 0, 0, 0,
        0, 1, 0, 0,
        0, 0, 1, 0,
        0, 0, 0, 1,
    }
    return New(4, body), New(4, echelon), New(4, diag), 4
}

func newNonMaxRankSquare() (*Matrix, *Matrix, *Matrix, int) {
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
    diag := []uint8{
        1, 0, 0, 1, 0,
        0, 1, 0, 1, 1,
        0, 0, 1, 0, 1,
        0, 0, 0, 0, 0,
    }
    return New(4, body), New(4, echelon), New(4, diag), 3
}

func newMaxRankLong() (*Matrix, *Matrix, *Matrix, int) {
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
    diag := []uint8{
        1, 0, 0, 0, 0,
        0, 1, 0, 0, 0,
        0, 0, 1, 0, 0,
        0, 0, 0, 1, 0,
        0, 0, 0, 0, 1,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
    }
    return New(7, body), New(7, echelon), New(7, diag), 5
}

func newZero() (*Matrix, *Matrix, *Matrix, int) {
    body := []uint8{
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
    }
    return New(7, body), New(7, body), New(7, body), 0
}

func newRank1() (*Matrix, *Matrix, *Matrix, int) {
    body := []uint8{
        0, 1, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 1, 0, 0, 0,
        0, 0, 0, 0, 0,
    }
    echelon := []uint8{
        0, 1, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
    }
    return New(7, body), New(7, echelon), New(7, echelon), 1
}

func newNonMaxRankLong() (*Matrix, *Matrix, *Matrix, int) {
    body := []uint8{
        0, 1, 1, 1, 0, 0,
        0, 0, 1, 0, 1, 0,
        1, 1, 0, 0, 1, 0,
        1, 1, 1, 0, 0, 0,
        1, 0, 0, 1, 0, 0,
        1, 1, 1, 1, 1, 0,
        0, 1, 0, 1, 0, 0,
    }
    echelon := []uint8{
        1, 0, 0, 1, 0, 0,
        0, 1, 1, 1, 0, 0,
        0, 0, 1, 0, 1, 0,
        0, 0, 0, 1, 1, 0,
        0, 0, 0, 0, 1, 0,
        0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0,
    }
    diag := []uint8{
        1, 0, 0, 0, 0, 0,
        0, 1, 0, 0, 0, 0,
        0, 0, 1, 0, 0, 0,
        0, 0, 0, 1, 0, 0,
        0, 0, 0, 0, 1, 0,
        0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0,
    }
    return New(7, body), New(7, echelon), New(7, diag), 5
}
