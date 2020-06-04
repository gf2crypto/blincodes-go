package matrix

import "fmt"
import "testing"
import "github.com/gf2crypto/blincodes-go/vector"

var matrices = [](func() (*Matrix, *Matrix, *Matrix, int)){
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
    for _, m := range matrices {
        mat, echelon, _, _ := m()
        if !echelon.Equal(mat.Echelon()) {
            t.Errorf("%v != %v", mat.Echelon(), echelon)
        }
    }
}

//TestEvaluatiotDiagonalForm test the evaluation of Matrix diagonal form
func TestEvaluatiotDiagonalForm(t *testing.T) {
    for _, m := range matrices {
        mat, _, diag, _ := m()
        if !diag.Equal(mat.Diagonal()) {
            t.Errorf("%v != %v", mat.Diagonal(), diag)
        }
    }
}

//TestEvaluatiotRank test the evaluation of Matrix's rank
func TestEvaluatiotRank(t *testing.T) {
    for _, m := range matrices {
        mat, _, _, rank := m()
        if mat.Rank() != rank {
            t.Errorf("%v.Rank() = %v, but got %v", mat, mat.Rank(), rank)
        }
    }
}

//TestOrthogonal test the evaluation orthogonal matrix
func TestOrthogonal(t *testing.T) {
    for _, m := range matrices {
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
    for _, m := range matrices {
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

//TestSolveUniqeSol test the solving linear equations
func TestSolveUniqeSol(t *testing.T) {
    vec := vector.New([]uint8{1, 0, 1, 0})
    mat, _, _, _ := newMaxRank()
    fund, base := mat.Solve(vec)
    if len(fund) != 0 {
        t.Errorf("wrong fundamental system of equation %vx^T=(%v)^T, expected [], but got %v",
            mat, vec, fund)
    }
    sol := &Matrix{body: [](*vector.Vector){base}, ncolumns: base.Len()}
    b := &Matrix{body: [](*vector.Vector){vec}, ncolumns: vec.Len()}
    if !mat.Mul(sol.T()).Equal(b.T()) {
        t.Errorf("wrong base solution of equation %vx^T=(%v)^T: %v != %v:\n",
            mat, vec, mat.Mul(sol.T()), b.T())
    }
}

//TestSolveNoSol test the solving linear equations
func TestSolveNoSol(t *testing.T) {
    vec := vector.New([]uint8{1, 0, 1, 0})
    mat, _, _, _ := newNonMaxRankSquare()
    fund, base := mat.Solve(vec)
    if len(fund) != 0 {
        t.Errorf("wrong fundamental system of equation %vx^T=(%v)^T, expected [], but got %v",
            mat, vec, fund)
    }
    if base != nil {
        t.Errorf("wrong solution of equation %vx^T=(%v)^T, expected nil, but got %v:\n",
            mat, vec, base)
    }
}

// TestSolveMultSol test the solving linear equations
func TestSolveMultSol(t *testing.T) {
    vec := vector.New([]uint8{1, 1, 1, 0})
    mat, _, _, _ := newNonMaxRankSquare()
    fund, base := mat.Solve(vec)
    if len(fund) != 2 {
        t.Errorf("wrong dimension fundamental system of equation %vx^T=(%v)^T, expected 2, but got %v",
            mat, vec, len(fund))
    }
    m := &Matrix{body: fund, ncolumns: fund[0].Len()}
    if !mat.Mul(m.T()).Equal(New(4, 2)) {
        t.Errorf("wrong fundamental system of equation %vx^T=(%v)^T, AF^T != 0:\n%v * %v^T !=0",
            mat, vec, mat, m)
    }
    sol := &Matrix{body: [](*vector.Vector){base}, ncolumns: base.Len()}
    b := &Matrix{body: [](*vector.Vector){vec}, ncolumns: vec.Len()}
    if !mat.Mul(sol.T()).Equal(b.T()) {
        t.Errorf("wrong base solution of equation %vx^T=(%v)^T: %v != %v:\n",
            mat, vec, mat.Mul(sol.T()), b.T())
    }
}

//TestRandom tests the generating of random matrices
func TestRandom(t *testing.T) {
    sizes := []([2]int){
        [2]int{5, 10},
        [2]int{10, 10},
        [2]int{100, 10},
        [2]int{200, 10},
        [2]int{300, 10},
        [2]int{400, 10},
        [2]int{500, 10},
        [2]int{1000, 10},
        [2]int{1024, 10},
    }
    for _, s := range sizes {
        if ok, er := basicTestRandom(s[0], s[1]); !ok {
            t.Errorf(er)
        }
    }
}

func basicTestRandom(size, ntests int) (bool, string) {
    mats := make([](*Matrix), ntests)
    for i := 0; i < ntests; i++ {
        mats[i] = Random(size)
        for j := 0; j < i; j++ {
            if mats[i].Equal(mats[j]) {
                return false, fmt.Sprintf("problem with Random functions, found two equal matrices:\n%v\n%v\n", mats[i], mats[j])
            }
        }
    }
    return true, ""
}

//TestNonsingRandom tests the generating of random nonsingular matrices
func TestNonsingRandom(t *testing.T) {
    sizes := []([2]int){
        [2]int{5, 10},
        [2]int{10, 10},
        [2]int{100, 10},
        [2]int{200, 10},
        [2]int{300, 10},
        [2]int{400, 10},
        [2]int{500, 10},
        [2]int{1000, 10},
        [2]int{1024, 10},
    }
    for _, s := range sizes {
        if ok, er := basicNonsingTestRandom(s[0], s[1]); !ok {
            t.Errorf(er)
        }
    }
}

//TestRandomMaxRank tests the generating of random max-rank matrices
func TestRandomMaxRank(t *testing.T) {
    sizes := []([2]int){
        [2]int{5, 2},
        [2]int{6, 2},
        [2]int{10, 10},
        [2]int{100, 10},
        [2]int{200, 10},
        [2]int{300, 10},
        [2]int{400, 10},
        [2]int{500, 10},
        [2]int{1000, 10},
        [2]int{1024, 10},
    }
    for _, s := range sizes {
        if ok, er := basicTestRandomMaxRank(int(s[0]/2), s[0], s[1]); !ok {
            t.Errorf(er)
        }
    }
    for _, s := range sizes {
        if ok, er := basicTestRandomMaxRank(s[0], int(s[0]/2), s[1]); !ok {
            t.Errorf(er)
        }
    }
}

//TestPermLeft tests the generating of the left-action permutation
func TestPermLeft(t *testing.T) {
    permM := []uint8{
        0, 0, 0, 1, 0, 0, 0,
        0, 0, 1, 0, 0, 0, 0,
        0, 0, 0, 0, 1, 0, 0,
        0, 1, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 1, 0,
    }
    perm := []int{3, 2, 4, 1, 6, 0, 5}
    if p, res := PermLeft(perm), New(7, permM); !p.Equal(res) {
        t.Errorf("wrong permutation matrix, expected:\n%v\ngot:\n%v\n", p, res)
    }
}

//TestPerm tests the generating of the right-action permutation
func TestPerm(t *testing.T) {
    permM := []uint8{
        0, 0, 0, 1, 0, 0, 0,
        0, 0, 1, 0, 0, 0, 0,
        0, 0, 0, 0, 1, 0, 0,
        0, 1, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 1, 0,
    }
    perm := []int{5, 3, 1, 0, 2, 6, 4}
    if p, res := Perm(perm), New(7, permM); !p.Equal(res) {
        t.Errorf("wrong permutation matrix, expected:\n%v\ngot:\n%v\n", p, res)
    }
}

func basicNonsingTestRandom(size, ntests int) (bool, string) {
    mats := make([](*Matrix), ntests)
    for i := 0; i < ntests; i++ {
        mats[i] = Nonsing(size)
        if r := mats[i].Rank(); r != size {
            return false, fmt.Sprintf("problem with Nonsing functions, rank of matrix is not maximal:\n%v\n%v!=%v\n", mats[i], r, size)
        }
        for j := 0; j < i; j++ {
            if mats[i].Equal(mats[j]) {
                return false, fmt.Sprintf("problem with Random functions, found two equal matrices:\n%v\n%v\n", mats[i], mats[j])
            }
        }
    }
    return true, ""
}

func basicTestRandomMaxRank(m, n, ntests int) (bool, string) {
    mats := make([](*Matrix), ntests)
    for i := 0; i < ntests; i++ {
        mats[i] = RandomMaxRank(m, n)
        rank := m
        if rank > n {
            rank = n
        }
        if r := mats[i].Rank(); r != rank {
            return false, fmt.Sprintf("problem with Nonsing functions, rank of matrix is not maximal:\n%v\n%v!=%v\n", mats[i], r, rank)
        }
        for j := 0; j < i; j++ {
            if mats[i].Equal(mats[j]) {
                return false, fmt.Sprintf("problem with Random functions, found two equal matrices:\n%v\n%v\n", mats[i], mats[j])
            }
        }
    }
    return true, ""
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
