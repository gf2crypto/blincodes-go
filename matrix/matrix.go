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
