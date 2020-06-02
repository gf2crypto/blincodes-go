package matrix

import "bytes"
import "fmt"
import "sort"
import "github.com/gf2crypto/blincodes-go/vector"

//Matrix represents binary Matrix
//This type is immutable
type Matrix struct {
    body     [](*vector.Vector)
    ncolumns int
}

func (mat *Matrix) String() string {
    var buf bytes.Buffer
    for i, row := range mat.body {
        fmt.Fprintf(&buf, "%s", row)
        if i < len(mat.body)-1 {
            fmt.Fprint(&buf, "\n")
        }
    }
    return buf.String()
}

// PrettyString returns pretty formatted string of Matrix representation
// Example:
// 0101011 -> -1-1-11
// 1011001 -> 1-11--1
func (mat *Matrix) PrettyString() string {
    var buf bytes.Buffer
    for i, row := range mat.body {
        fmt.Fprintf(&buf, "%s", row.PrettyString())
        if i < len(mat.body)-1 {
            fmt.Fprint(&buf, "\n")
        }
    }
    return buf.String()
}

//Nrows returns number of matrix rows
func (mat *Matrix) Nrows() int {
    return len(mat.body)
}

//Ncolumns returns number of matrix rows
func (mat *Matrix) Ncolumns() int {
    return mat.ncolumns
}

//Shapes returns number of rows and columns of Matrix
func (mat *Matrix) Shapes() (int, int) {
    return mat.Nrows(), mat.Ncolumns()
}

//Rank returns rank of Matrix
func (mat *Matrix) Rank() int {
    _, iw := mat.GaussElim(false)
    return len(iw)
}

//Echelon returns echelon form of Matrix
func (mat *Matrix) Echelon() *Matrix {
    ech, _ := mat.GaussElim(false)
    return ech
}

//Diagonal returns diagonal form of Matrix
func (mat *Matrix) Diagonal() *Matrix {
    diag, _ := mat.GaussElim(true)
    return diag
}

//Orthogonal returns orthogonal matrix for input matrix
//Orthogonal matrix H is matrix satisfied the following condition
//       mat * H^T = 0.
// Moreover the matrix H has highest possible rank
func (mat *Matrix) Orthogonal() *Matrix {
    if mat.Nrows() == 0 || mat.Ncolumns() == 0 {
        return newEmpty(0, 0)
    }
    diag, iw := mat.GaussElim(true)
    if len(iw) == mat.Ncolumns() {
        return New(1, mat.Ncolumns())
    }
    sort.Slice(diag.body, func(i, j int) bool { return diag.body[j].Less(diag.body[i]) })
    sort.Slice(iw, func(i, j int) bool { return iw[i] < iw[j] })
    matGIndexes := make([]int, 0, mat.Ncolumns()-len(iw))
    for i := 0; i < mat.Ncolumns(); i++ {
        j := sort.SearchInts(iw, i)
        if j >= len(iw) || (iw[j] != i) {
            matGIndexes = append(matGIndexes, i)
        }
    }
    matG := (&Matrix{body: diag.body[:len(iw)], ncolumns: diag.ncolumns}).Submatrix(matGIndexes)
    matI := Identity(mat.Ncolumns() - len(iw))
    body := make([](*vector.Vector), mat.Ncolumns())
    for i, ind := range iw {
        body[ind] = matG.body[i]
    }
    for i, ind := range matGIndexes {
        body[ind] = matI.body[i]
    }
    return (&Matrix{body: body, ncolumns: mat.Ncolumns() - len(iw)}).T()
}

//Inv evaluates generalized inverse of matrix
func (mat *Matrix) Inv() *Matrix {
    m := make([]([2](*vector.Vector)), mat.Nrows())
    mIdent := Identity(mat.Nrows())
    for i, row := range mat.body {
        m[i] = [2](*vector.Vector){row.Copy(), mIdent.body[i]}
    }
    for i, rows := range m {
        firstOne := -1
        for j := 0; j < mat.Ncolumns(); j++ {
            if rows[0].Get(j) != 0 {
                firstOne = j
                break
            }
        }
        if firstOne < 0 {
            continue
        }
        for k := 0; k < mat.Nrows(); k++ {
            if (k != i) && (m[k][0].Get(firstOne) != 0) {
                m[k][0], m[k][1] = rows[0].Xor(m[k][0]), rows[1].Xor(m[k][1])
            }
        }
    }
    sort.Slice(m, func(i, j int) bool { return m[j][0].Less(m[i][0]) })
    for i := 0; i < mat.Nrows(); i++ {
        mIdent.body[i] = m[i][1]
    }
    return mIdent
}

//GaussElim evaluates Gaussian elimination
// GaussElim(full=false) -> classic, only forward
// GaussElim(full=true) -> forward and reverse
// GaussElim(full bool, edge int) -> only for columns with numbers < edge
// GaussElim(full bool, columns []int) -> only for columns from input slice
//Return:
//     *Matrix - result of elimination
//     []int - information window, i.e. positions of maximum rank submatrix
func (mat *Matrix) GaussElim(full bool, limit ...interface{}) (*Matrix, []int) {
    if mat.Nrows() == 0 || mat.Ncolumns() == 0 {
        return newEmpty(0, 0), make([]int, 0)
    }
    var lim interface{}
    if len(limit) >= 2 {
        panic(fmt.Errorf("matrix: GaussElim expected no more 2 arguments, but got %v", len(limit)))
    }
    if len(limit) == 0 {
        lim = mat.Ncolumns()
    } else {
        lim = limit[0]
    }
    res := mat.Copy()
    infoWindow := make([]int, 0, mat.Ncolumns())
    for i, row := range res.body {
        firstOne := -1
        switch l := lim.(type) {
        case int:
            for j := 0; j < l; j++ {
                if row.Get(j) != 0 {
                    firstOne = j
                    break
                }
            }
        case []int:
            for _, j := range l {
                if row.Get(j) != 0 {
                    firstOne = j
                    break
                }
            }
        }
        if firstOne < 0 {
            continue
        }
        infoWindow = append(infoWindow, firstOne)
        start := 0
        if !full {
            start = i + 1
        }
        for k := start; k < mat.Nrows(); k++ {
            if (k != i) && (res.body[k].Get(firstOne) != 0) {
                res.body[k] = row.Xor(res.body[k])
            }
        }
    }
    sort.Slice(res.body, func(i, j int) bool { return res.body[j].Less(res.body[i]) })
    sort.Slice(infoWindow, func(i, j int) bool { return infoWindow[i] < infoWindow[j] })
    return res, infoWindow
}

//Copy copies matrix
func (mat *Matrix) Copy() *Matrix {
    if mat.Ncolumns() == 0 || mat.Nrows() == 0 {
        return newEmpty(0, 0)
    }
    body := make([](*vector.Vector), 0, mat.Nrows())
    for _, row := range mat.body {
        body = append(body, row.Copy())
    }
    return &Matrix{body: body, ncolumns: mat.ncolumns}
}

//T returns transpose of matrix
func (mat *Matrix) T() *Matrix {
    body := make([](*vector.Vector), 0, mat.Ncolumns())
    // fmt.Printf("debug: ncolumns=%v,matrix:\n%v\n", mat.Ncolumns(), mat)
    for j := 0; j < mat.Ncolumns(); j++ {
        body = append(body, mat.GetColumn(j))
    }
    return &Matrix{body: body, ncolumns: mat.Nrows()}
}

//Submatrix returns submatrix of Matrix defined by array of indexes inds
func (mat *Matrix) Submatrix(inds []int) *Matrix {
    body := make([](*vector.Vector), 0, len(inds))
    ncolumns := 0
    for i := 0; i < mat.Nrows(); i++ {
        row := make([]uint8, 0, len(inds))
        for _, j := range inds {
            if j < mat.Ncolumns() && j >= 0 {
                row = append(row, uint8(mat.body[i].Get(j)))
            }
        }
        if len(row) != 0 {
            body = append(body, vector.New(row))
        }
        if len(row) != ncolumns {
            ncolumns = len(row)
        }
    }
    return &Matrix{body: body, ncolumns: ncolumns}
}

//ConcatenateRows concatenates rows of two matrices
func (mat *Matrix) ConcatenateRows(mat0 *Matrix) *Matrix {
    if mat.Ncolumns() != mat0.Ncolumns() {
        panic(fmt.Errorf("matrix: cannon concatenate of matrices, because they have different number of columns, %v != %v ", mat.Ncolumns(), mat0.Ncolumns()))
    }
    body := append(make([](*vector.Vector), 0, mat.Nrows()+mat0.Nrows()), mat.body...)
    body = append(body, mat0.body...)
    return &Matrix{body: body, ncolumns: mat.Ncolumns()}
}

//ConcatenateColumns concatenates columns of matrices
func (mat *Matrix) ConcatenateColumns(mat0 *Matrix) *Matrix {
    if mat.Nrows() != mat0.Nrows() {
        panic(fmt.Errorf("matrix: cannon concatenate of matrices, because they have different number of rows, %v != %v ", mat.Nrows(), mat0.Nrows()))
    }
    body := make([](*vector.Vector), 0, mat.Nrows())
    for i, row := range mat.body {
        body = append(body, row.Concatenate(mat0.body[i]))
    }
    return &Matrix{body: body, ncolumns: mat.Ncolumns() + mat0.Ncolumns()}
}

//Add returns sum of matrices
func (mat *Matrix) Add(mat0 *Matrix) *Matrix {
    if mat.Nrows() != mat0.Nrows() {
        panic(fmt.Errorf("matrix: cannon evaluate sum of matrices, because they have different number of rows, %v != %v ", mat.Nrows(), mat0.Nrows()))
    }
    if mat.Ncolumns() != mat0.Ncolumns() {
        panic(fmt.Errorf("matrix: cannon evaluate sum of matrices, because they have different number of columns, %v != %v ", mat.Ncolumns(), mat0.Ncolumns()))
    }
    body := make([](*vector.Vector), 0, mat.Nrows())
    for i, row := range mat.body {
        body = append(body, row.Xor(mat0.body[i]))
    }
    return &Matrix{body: body, ncolumns: mat.Ncolumns()}
}

//Mul returns multiplication of matrices
func (mat *Matrix) Mul(mat0 *Matrix) *Matrix {
    if mat.Ncolumns() != mat0.Nrows() {
        panic(fmt.Errorf("matrix: cannon multiplicate matrices, because they have wrong dimension, mat.nrows != mat0.ncolumns (%v != %v) ", mat.Ncolumns(), mat0.Nrows()))
    }
    body := make([](*vector.Vector), 0, mat.Nrows())
    for _, row := range mat.body {
        res := vector.New(mat0.Ncolumns())
        for _, i := range row.Support() {
            res = res.Xor(mat0.body[i])
        }
        body = append(body, res)
    }
    return &Matrix{body: body, ncolumns: mat0.Ncolumns()}
}

//Equal returns true if mat == mat0
func (mat *Matrix) Equal(mat0 *Matrix) bool {
    if mat.Nrows() != mat0.Nrows() {
        return false
    }
    for i, row := range mat.body {
        if !row.Equal(mat0.body[i]) {
            return false
        }
    }
    return true
}

//GetRow returns i-th row
func (mat *Matrix) GetRow(i int) *vector.Vector {
    if (mat.Nrows() <= i) || (i < 0) {
        panic(fmt.Errorf("matrix: index row %v out of range, expected %v <= i <= %v",
            i, 0, mat.Nrows()-1))
    }
    return mat.body[i]
}

//GetColumn returns copy of i-th column
func (mat *Matrix) GetColumn(i int) *vector.Vector {
    if (mat.Ncolumns() <= i) || (i < 0) {
        panic(fmt.Errorf("matrix: index column %v out of range, expected %v <= i <= %v",
            i, 0, mat.Ncolumns()-1))
    }
    row := make([]uint8, 0, mat.Nrows())
    for j := 0; j < mat.Nrows(); j++ {
        row = append(row, uint8(mat.body[j].Get(i)))
    }
    return vector.New(row)
}

//Solve solves linear equation Ax^T=b^T
func (mat *Matrix) Solve(v *vector.Vector) ([](*vector.Vector), *vector.Vector) {
    if mat.Nrows() != v.Len() {
        panic(fmt.Errorf("matrix: cannot solve equation, wrong dimension, expected length of vector equal to number of matrix rows, %v!=%v",
            v.Len(), mat.Nrows()))
    }
    extended := mat.ConcatenateColumns((&Matrix{body: [](*vector.Vector){v}, ncolumns: v.Len()}).T())
    orth := extended.Orthogonal()
    fund := make([](*vector.Vector), 0, orth.Nrows()-1)
    var sol *vector.Vector
    for _, row := range orth.body {
        if row.Get(row.Len()-1) == 0 {
            fund = append(fund, row.Resize(-1))
        } else {
            if sol != nil {
                return [](*vector.Vector){}, nil
            }
            sol = row.Resize(-1)
        }
    }
    if sol == nil {
        return [](*vector.Vector){}, nil
    }
    return fund, sol
}
