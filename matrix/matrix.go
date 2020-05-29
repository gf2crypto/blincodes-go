package matrix

import "bytes"
import "fmt"
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
