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
    rank := 0
    echelon := mat.Copy()
    for i, row := range echelon.body {
        firstOne := -1
        for j := 0; j < mat.Ncolumns(); j++ {
            if row.Get(j) != 0 {
                firstOne = j
                break
            }
        }
        if firstOne < 0 {
            continue
        }
        rank++
        for k := i + 1; k < mat.Nrows(); k++ {
            if echelon.body[k].Get(firstOne) != 0 {
                echelon.body[k] = row.Xor(echelon.body[k])
            }
        }
    }
    return rank
}

//Echelon returns echelon form of Matrix
func (mat *Matrix) Echelon() *Matrix {
    if mat.Nrows() == 0 || mat.Ncolumns() == 0 {
        return newEmpty(0, 0)
    }
    echelon := mat.Copy()
    for i, row := range echelon.body {
        firstOne := -1
        for j := 0; j < mat.Ncolumns(); j++ {
            if row.Get(j) != 0 {
                firstOne = j
                break
            }
        }
        if firstOne < 0 {
            continue
        }
        for k := i + 1; k < mat.Nrows(); k++ {
            if echelon.body[k].Get(firstOne) != 0 {
                echelon.body[k] = row.Xor(echelon.body[k])
            }
        }
    }
    sort.Slice(echelon.body, func(i, j int) bool { return echelon.body[j].Less(echelon.body[i]) })
    return echelon
}

//Diagonal returns diagonal form of Matrix
func (mat *Matrix) Diagonal() *Matrix {
    if mat.Nrows() == 0 || mat.Ncolumns() == 0 {
        return newEmpty(0, 0)
    }
    diag := mat.Copy()
    for i, row := range diag.body {
        firstOne := -1
        for j := 0; j < mat.Ncolumns(); j++ {
            if row.Get(j) != 0 {
                firstOne = j
                break
            }
        }
        if firstOne < 0 {
            continue
        }
        for k := 0; k < mat.Nrows(); k++ {
            if (k != i) && (diag.body[k].Get(firstOne) != 0) {
                diag.body[k] = row.Xor(diag.body[k])
            }
        }
    }
    sort.Slice(diag.body, func(i, j int) bool { return diag.body[j].Less(diag.body[i]) })
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
    diag := mat.Copy()
    // identityIndexes - list of indexes which corresponding identity matrix
    // G = (I|A)
    identityIndexes := make([]int, 0, diag.Ncolumns())
    // Make Gaussian elimination
    for i, row := range diag.body {
        firstOne := -1
        for j := 0; j < mat.Ncolumns(); j++ {
            if row.Get(j) != 0 {
                firstOne = j
                break
            }
        }
        if firstOne < 0 {
            continue
        }
        identityIndexes = append(identityIndexes, firstOne)
        for k := 0; k < mat.Nrows(); k++ {
            if (k != i) && (diag.body[k].Get(firstOne) != 0) {
                diag.body[k] = row.Xor(diag.body[k])
            }
        }
    }
    if len(identityIndexes) == mat.Ncolumns() {
        return New(1, mat.Ncolumns())
    }
    sort.Slice(diag.body, func(i, j int) bool { return diag.body[j].Less(diag.body[i]) })
    sort.Slice(identityIndexes, func(i, j int) bool { return identityIndexes[i] < identityIndexes[j] })
    matGIndexes := make([]int, 0, mat.Ncolumns()-len(identityIndexes))
    for i := 0; i < mat.Ncolumns(); i++ {
        j := sort.SearchInts(identityIndexes, i)
        if j >= len(identityIndexes) || (identityIndexes[j] != i) {
            matGIndexes = append(matGIndexes, i)
        }
    }
    matG := (&Matrix{body: diag.body[:len(identityIndexes)], ncolumns: diag.ncolumns}).Submatrix(matGIndexes)
    matI := Identity(mat.Ncolumns() - len(identityIndexes))
    body := make([](*vector.Vector), mat.Ncolumns())
    for i, ind := range identityIndexes {
        body[ind] = matG.body[i]
    }
    for i, ind := range matGIndexes {
        body[ind] = matI.body[i]
    }
    return (&Matrix{body: body, ncolumns: mat.Ncolumns() - len(identityIndexes)}).T()
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
        row := make([]uint8, 0, mat.Ncolumns())
        for i := 0; i < mat.Nrows(); i++ {
            row = append(row, uint8(mat.body[i].Get(j)))
        }
        body = append(body, vector.New(row))
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
