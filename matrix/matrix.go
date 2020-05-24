package matrix

import "github.com/gf2crypto/blincodes-go/vector"

//Matrix represents binary Matrix
//This type is immutable
type Matrix struct {
    body     []vector.Vector
    nrows    int
    ncolumns int
}
