package matrix

import (
    "bytes"
    "crypto/rand"
    "math/big"
    "strings"
)
import "fmt"
import "github.com/gf2crypto/blincodes-go/vector"

//Matrix represents binary Matrix
//This type is immutable
type Matrix struct {
    body     []*vector.Vector
    nColumns uint
}

func New() *Matrix {
    return new(Matrix).SetZero(0,0)
}

//SetV sets m to matrix from vectors list, returns m
func (m *Matrix) SetV(vectors []*vector.Vector) *Matrix {
    if vectors == nil {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, len(vectors))
    nColumns := uint(0)
    flag := false
    for i, v := range vectors {
        if flag && v.Len() != nColumns {
            panic(fmt.Errorf("matrix: expected all vectors have the same length, but %d!=%d",
                nColumns, v.Len()))
        }
        if !flag {
            flag = true
            nColumns = v.Len()
        }
        body[i] = new(vector.Vector).SetV(v)
    }
    m.body, m.nColumns = body, nColumns
    if m.nColumns == 0 {
        m.body = make([]*vector.Vector, 0)
    }
    return m
}

////SetV sets m to matrix from one vector v, returns m
//func (m *Matrix) SetV(vec *vector.Vector) *Matrix {
//    m.body, m.nColumns = []*vector.Vector{vec}, vec.Len()
//    return m
//}

//SetZero sets m to zero (k x n)-matrix, returns m
func (m *Matrix) SetZero(k, n uint) *Matrix {
    if k == 0 || n == 0 {
        m.body, m.nColumns = nil, 0
        return m
    }
    m.body, m.nColumns = make([]*vector.Vector, int(k)), n
    for i:=0; i < int(k); i++ {
        m.body[i] = new(vector.Vector).SetZero(m.nColumns)
    }
    return m
}

//SetBytes packs the byte 2-D array into the binary matrix
func (m *Matrix)SetBytes(b [][]byte, n uint) *Matrix {
    if n == 0 || len(b) == 0 {
        return m
    }
    v := make([]*vector.Vector, len(b))
    for i, a := range b {
        v[i] = new(vector.Vector).SetBytes(a, n)
    }
    return m.SetV(v)
}

//SetBits converts the array of bit array to the binary matrix, sets m to the result and returns m
func (m *Matrix)SetBits(b [][]byte) *Matrix {
    if len(b) == 0 {
        return m.SetZero(0,0)
    }
    v := make([]*vector.Vector, len(b))
    for i, a := range b {
        v[i] = new(vector.Vector).SetBitArray(a)
    }
    return m.SetV(v)
}

//SetUnit sets m to units (k x n)-matrix, returns m
func (m *Matrix) SetUnit(k, n uint) *Matrix {
    if k == 0 || n == 0 {
        m.body, m.nColumns = make([]*vector.Vector,0), 0
        return m
    }
    m.body, m.nColumns = make([]*vector.Vector, int(k)), n
    for i:=0; i < int(k); i++ {
        m.body[i] = new(vector.Vector).SetUnits(m.nColumns)
    }
    return m
}

//SetIdentity sets m to the identity matrix of size n
func (m *Matrix) SetIdentity(n uint) *Matrix {
    if n == 0 {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, n)
    for i := 0; i < int(n); i++ {
        body[i] = new(vector.Vector).SetSupport(n, []uint{uint(i)})
    }
    return m.SetV(body)
}

//SetRandom sets m to the random (k x n) - matrix, returns m
func (m *Matrix) SetRandom(k, n uint) *Matrix {
    if k ==0 || n == 0 {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, k)
    for i := 0; i < int(k); i++ {
        body[i] = new(vector.Vector).SetRandom(n)
    }
    return m.SetV(body)
}

// SetStrings converts array string to Matrix, sets m to the result and returns m.
// Function supports the following filler for zero symbol:
//       '0' == '0', '-', '*'
func (m *Matrix) SetStrings(s []string) *Matrix {
    body := make([]*vector.Vector, 0 ,len(s))
    for _, r := range s {
        if r == "" {
            continue
        }
        v, err := new(vector.Vector).Parse(r)
        if err != nil {
            panic(err)
        }
        body = append(body, v)
    }
    return m.SetV(body)
}

// Parse converts string to Matrix, sets m to the result and returns m.
// Function supports the following filler for zero symbol:
//       '0' == '0', '-', '*'
func (m *Matrix) Parse(s string) *Matrix {
    return m.SetStrings(strings.Split(s, "\n"))
}

//SetPerm returns permutation matrix corresponding to the right-action permutation p
// Attention! Perm does not check that p is correct permutation, i.e. p contains differ elements.
// For example,
// Permutation matrix for p = {5, 1, 2, 0, 4, 3} is
// 000100
// 010000
// 001000
// 000001
// 000010
// 100000
// Permutation matrix for p = {1, 1, 1, 1, 1, 1} is
// 000000
// 111111
// 000000
// 000000
// 000000
// 000000
// Permutation matrix for p = {11, 7, 8, 6, 10, 9} is
// 000100
// 010000
// 001000
// 000001
// 000010
// 100000
func (m *Matrix) SetPerm(p []uint) *Matrix {
    if len(p) == 0 {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, len(p))
    for i := 0; i < len(p); i++ {
        body[p[i] % uint(len(p))] = new(vector.Vector).SetSupport(uint(len(p)), []uint{uint(i)})
    }
    return m.SetV(body)
}

//SetPermL returns permutation matrix corresponding to the left-action permutation p
// Attention! Perm does not check that p is correct permutation, i.e. p contains differ elements.
// For example,
// Permutation matrix for p = {5, 1, 2, 0, 4, 3} is
// 000001
// 010000
// 001000
// 100000
// 000010
// 000100
// Permutation matrix for p = {1, 1, 1, 1, 1, 1} is
// 010000
// 010000
// 010000
// 010000
// 010000
// 010000
// Permutation matrix for p = {11, 7, 8, 6, 10, 9} is
// 000001
// 010000
// 001000
// 100000
// 000010
// 000100
func (m *Matrix) SetPermL(p []uint) *Matrix {
    if len(p) == 0 {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, len(p))
    for i := 0; i < len(p); i++ {
        body[i] = new(vector.Vector).SetSupport(uint(len(p)), []uint{p[i] % uint(len(p))})
    }
    return m.SetV(body)
}

//RandomMaxRank sets m to random (k x n) - matrix of maximal rank, returns m
//     nonsing = nonsingular(min(nrows, ncolumns))
//     perm_matrix = permutation(sample(range(max(nrows, ncolumns)),
//                                      max(nrows, ncolumns)))
//     if nrows < ncolumns:
//         return concatenate(nonsing, random(ncolumns - nrows)) * perm_matrix
//     return perm_matrix * concatenate(nonsing,
//                                      random(nrows - ncolumns),
//                                      by_rows=True)
func (m *Matrix)RandomMaxRank(k, n uint) *Matrix {
    if k == n {
        return m.NonSing(k)
    }
    if k > n {
        m.ConRows(new(Matrix).NonSing(n),new(Matrix).SetRandom(k - n, n))
        return m.Dot(new(Matrix).SetPermL(GetRandomPerm(k)), m)
    }
    // k < n
    m.ConCols(new(Matrix).NonSing(k), new(Matrix).SetRandom(k, n - k))
    return m.Dot(m, new(Matrix).SetPerm(GetRandomPerm(n)))
}

func GetRandomPerm(n uint) []uint{
    p := make(map[uint] byte, n)
    for i:=uint(0); i < n; i++ {
        for isDuplicate := true; isDuplicate; {
            jB, err := rand.Int(rand.Reader, new(big.Int).SetUint64(uint64(n)))
            if err != nil {
                panic(err)
            }
            j := uint(jB.Uint64())
            _, ok := p[j]
            if !ok {
                p[j] = 0
                isDuplicate = false
            }
        }
    }
    res := make([]uint, 0, n)
    for i := range p {
        res = append(res, i)
    }
    return res
}

//NonSing sets mo to random non singular matrix, returns m
// Function uses algorithm of Dana Randall
// https://www.researchgate.net/publication/2729950_Efficient_Generation_of_Random_Nonsingular_Matrices
func (m *Matrix) NonSing(n uint) *Matrix {
    if n == 0 {
        return m.SetZero(0,0)
    }
    a := new(Matrix).SetZero(n, n)
    m.body, m.nColumns = make([]*vector.Vector, n), n
    active := make([]uint, n)
    for i := uint(0); i < n; i++ {
        active[i] = i
    }
    for i := uint(0); i < n; i++ {
        //Generate random v != 0
        // And found its first 1
        v := new(vector.Vector).SetZero(0)
        var r uint  // position of the first one
        switch i {
        case n-1:
            v = new(vector.Vector).SetUnits(1)
            r = 0
        default:
            //fmt.Println(active, passive)
            for isZero := true; isZero; {
                v = new(vector.Vector).SetRandom(n-i)
                for j := uint(0); j < v.Len(); j++ {
                    if v.Get(j) != 0 {
                        isZero = false
                        r = j // index of the first 1
                    }
                }
            }
        }
        //Update matrix A
        a.body[i].SetBit(a.body[i], active[r], 1)  // A[i][r] = 1
        // There is a mistake in the paper of Dana Randall: this code was missing in her paper
        w := new(vector.Vector).SetRandom(n-i-1)
        for j := i + 1; j < n; j++ {
            a.body[j].SetBit(a.body[j], active[r], w.Get(j - i - 1)) // A[i+1, n-1][r] <- random bit
        }
        //Update matrix T: T[r]  <- v
        m.body[active[r]] = new(vector.Vector).SetZero(n)
        for k, j := range active {
            m.body[active[r]].SetBit(m.body[active[r]], j, v.Get(uint(k)))
        }
        active = append(active[:r], active[r+1:]...)
    }
    return m.Dot(a, m)
}

func (m *Matrix) String() string {
    var buf bytes.Buffer
    for i, row := range m.body {
        _, err := fmt.Fprintf(&buf, "%s", row)
        if err != nil {
            panic(err)
        }
        if i < len(m.body)-1 {
            _, err = fmt.Fprint(&buf, "\n")
            if err != nil {
                panic(err)
            }
        }
    }
    return buf.String()
}

// PrettyString returns pretty formatted string of Matrix representation
// Example:
// 0101011 -> -1-1-11
// 1011001 -> 1-11--1
func (m *Matrix) PrettyString() string {
    var buf bytes.Buffer
    for i, row := range m.body {
        _, err := fmt.Fprintf(&buf, "%s", row.PrettyString())
        if err != nil {
            panic(err)
        }
        if i < len(m.body)-1 {
            _, err = fmt.Fprint(&buf, "\n")
            if err != nil {
                panic(err)
            }
        }
    }
    return buf.String()
}

// LaTeXString returns LaTeX string of Matrix representation
// Example:
// 0101011 -> 0&1&0&1&0&1&1\\
// 1011001 -> 1&0&1&1&0&0&1\\
func (m *Matrix) LaTeXString() string {
    var buf bytes.Buffer
    for i, row := range m.body {
        _, err := fmt.Fprintf(&buf, "%s", row.LaTeXString())
        if err != nil {
            panic(err)
        }
        if i < len(m.body)-1 {
            _, err =fmt.Fprint(&buf, "\\\n")
            if err != nil {
                panic(err)
            }
        }
    }
    return buf.String()
}

//NRows returns number of matrix rows
func (m *Matrix) NRows() uint {
    return uint(len(m.body))
}

//NColumns returns number of matrix rows
func (m *Matrix) NColumns() uint {
    return m.nColumns
}

//Shapes returns number of rows and columns of Matrix
func (m *Matrix) Shapes() (uint, uint) {
    return m.NRows(), m.NColumns()
}

//Rank returns rank of Matrix
func (m *Matrix) Rank() uint {
    return uint(len(new(Matrix).GaussElimination(m,true, nil)))
}

//Echelon sets m to echelon form of a.
// Parameters:
//   columns []uint - if len(columns) != 0 then construct echelon form only for submatrix on columns list.
func (m *Matrix) Echelon(a *Matrix, columns []uint) *Matrix {
    m.GaussElimination(a, true, columns)
    return m
}

//Diagonal sets m to diagonal form of Matrix a
// Parameters:
//   columns []uint - if len(columns) != 0 then construct diagonal form only for submatrix on columns list.
func (m *Matrix) Diagonal(a *Matrix, columns []uint) *Matrix {
    m.GaussElimination(a, false, columns)
    return m
}

//Ort returns orthogonal matrix for input matrix
//Orthogonal matrix H is matrix satisfied the following condition
//       mat * H^T = 0.
// Moreover the matrix H has highest possible rank
func (m *Matrix) Ort(a *Matrix) *Matrix {
    if a.NRows() == 0 || a.NColumns() == 0 {
        return m.SetZero(0,0)
    }
    diag := new(Matrix)
    iw := diag.GaussElimination(a, false, nil)
    r := a.NColumns() - uint(len(iw))
    if r == 0 {
        return m.SetZero(1, a.NColumns())
    }
    //matGIndexes := make([]uint, 0, int(a.NColumns())-len(iw))
    tmp := make(map[uint] byte, len(iw)) // for fast search of elements
    for i:=0; i < len(iw); i++ {
        tmp[iw[i]] = 0
    }
    body := make([]*vector.Vector, r)
    for i:=uint(0); i < r; i++ {
        body[i] = new(vector.Vector).SetZero(a.NColumns())
    }
    id := uint(0)
    p := uint(0)
    for i := uint(0); i < a.NColumns(); i++ {
        _, ok := tmp[i]
        switch ok {
        case true: // i belongs to information window
            for j, l := uint(0), uint(0); l < r; j++ {
                _, ok := tmp[j]
                if !ok {
                    body[l].SetBit(body[l], i, diag.body[p].Get(j))
                    l++
                }
            }
            p++
        case false: // i does not belong to information window
            body[id].SetBit(body[id], i, 1)
            id++
        }
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//Inv sets m generalized inverse of matrix a
func (m *Matrix) Inv(a *Matrix) *Matrix {
    if a.NRows() == 0 || a.NColumns() == 0 {
        return m.SetZero(0,0)
    }
    t := new(Matrix).SetM(a)
    m.SetIdentity(a.NRows())
    for i,j:=uint(0), uint(0); j < t.NColumns() && i < t.NRows(); j++  {
        k := i
        isZero := true
        for ; k < t.NRows(); k++ {
            if t.body[k].Get(j) == 1 {
                isZero = false
                break
            }
        }
        if isZero {
            continue
        }
        if i != k {
            t.body[i].Xor(t.body[i], t.body[k])
            m.body[i].Xor(m.body[i], m.body[k])
        }
        for l := uint(0); l < t.NRows(); l++ {
            if l != i && t.body[l].Get(j) == 1 {
                t.body[l].Xor(t.body[l], t.body[i])
                m.body[l].Xor(m.body[l], m.body[i])
            }
        }
        i++
    }
    return m
}

// GaussElimination sets the result of Gaussian elimination of matrix a to m.
//It returns []uint - information window, i.e. positions of maximum rank submatrix.
// Parameters:
//   onlyForward bool - if true then apply only forward step
//   columns []uint - if len(columns) != 0 then apply elimination only to submatrix on columns list.
func (m *Matrix) GaussElimination(a *Matrix, onlyForward bool, columns []uint) []uint {
    if a.NRows() == 0 || a.NColumns() == 0 {
        m.SetZero(0,0)
        return make([]uint, 0)
    }
    m.SetM(a)  // m = a
    edge := m.NColumns()
    if len(columns) != 0 {
        edge = uint(len(columns))
    }
    infoWindow := make([]uint, 0, edge)
    for i,j:=uint(0), uint(0); j < edge && i < m.NRows(); j++  {
        k := i
        isZero := true
        c := j
        if len(columns) != 0 {
            c = columns[j]
        }
        for ; k < m.NRows(); k++ {
            if m.body[k].Get(c) == 1 {
              isZero = false
              break
            }
        }
        if isZero {
            continue
        }
        infoWindow = append(infoWindow, c)
        if i != k {
            m.body[i].Xor(m.body[i], m.body[k])
        }
        start := uint(0)
        if onlyForward {
            start = i + 1
        }
        for t := start; t < m.NRows(); t++ {
            if t != i && m.body[t].Get(c) == 1 {
                m.body[t].Xor(m.body[t], m.body[i])
            }
        }
        i++
    }
    return infoWindow
}

//SetM sets m to a, i.e. m := a, return m
func (m *Matrix) SetM(a *Matrix) *Matrix {
    if a == m {
        return m
    }
    if a.NColumns() == 0 || a.NRows() == 0 {
        return m.SetZero(0,0)
    }
    if m.NRows() < a.NRows() {
        m.body = make([]*vector.Vector, a.NRows())
    } else {
        m.body = m.body[:a.NRows()]
    }
    m.nColumns = a.nColumns
    for i, row := range a.body {
        m.body[i] = new(vector.Vector).SetV(row)
    }
    return m
}

//T sets m to transpose of matrix a, returns m
// m = a^T
func (m *Matrix) T(a *Matrix) *Matrix {
    if a.NRows() == 0 || a.NColumns() == 0 {
        return m.SetZero(0,0)
    }
    body := make([]*vector.Vector, 0, a.NColumns())
    for j := 0; j < int(a.NColumns()); j++ {
        body = append(body, a.GetCol(uint(j)))
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//SubMatrix sets m to submatrix of Matrix a defined by index array of rows and
//index array of columns.
//It returns m.
func (m *Matrix) SubMatrix(a *Matrix, rows, cols []uint) *Matrix {

    isInArray := func(el uint, array []uint) bool {
        if array == nil || len(array) == 0 {
            return true
        }
        for _, t := range array{
            if t == el {
                return true
            }
        }
        return false
    }
    nr := uint(len(rows))
    if rows == nil || len(rows) == 0 {
        nr = a.NRows()
    }
    nc := uint(len(cols))
    if cols == nil || len(cols) == 0 {
        body := make([]*vector.Vector, 0, nr)
        for i := uint(0); i < a.NRows(); i++ {
            if !isInArray(i, rows) {
                continue
            }
            body = append(body, a.GetRow(i))
        }
        if len(body) == 0 {
            return m.SetZero(0,0)
        }
        m.body, m.nColumns = body, body[0].Len()
        return m
    }
    body := make([]*vector.Vector, 0, nr)
    for i := uint(0); i < a.NRows(); i++ {
        if !isInArray(i, rows) {
            continue
        }
        row := make([]uint8, 0, nc)
        for j:=uint(0); j < a.NColumns(); j++ {
            if !isInArray(j, cols) {
                continue
            }
            row = append(row, a.body[i].Get(j))
        }
        if len(row) == 0 {
            return m.SetZero(0,0)
        }
        body = append(body, new(vector.Vector).SetBitArray(row))
    }
    if len(body) == 0 {
        return m.SetZero(0, 0)
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//ConRows sets m to concatenation a's and b's rows, returns m
//      a
// m =  -
//      b
func (m *Matrix) ConRows(a, b *Matrix) *Matrix {
    if a.NColumns() != b.NColumns() {
        panic(fmt.Errorf("matrix: cannon concatenate of matrices, "+
            "because they have different number of columns, %v != %v ",
            a.NColumns(), b.NColumns()))
    }
    body := append(make([]*vector.Vector, 0, a.NRows()+b.NRows()), a.body...)
    body = append(body, b.body...)
    if len(body) == 0 {
        return m.SetZero(0,0)
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//ConCols sets m to concatenation of columns of matrices a and b, returns m
// m = (a | b)
func (m *Matrix) ConCols(a, b *Matrix) *Matrix {
    if a.NRows() != b.NRows() {
        panic(fmt.Errorf("matrix: cannon concatenate of matrices, "+
            "because they have different number of rows, %v != %v ", a.NRows(), b.NRows()))
    }
    body := make([]*vector.Vector, a.NRows())
    for i:=0; i < len(body); i++ {
        body[i] = new(vector.Vector).Concatenate(a.body[i], b.body[i])
    }
    if len(body) == 0 {
        return m.SetZero(0,0)
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//Add sets m to sum of matrices a and b, returns m
// m = a + b
func (m *Matrix) Add(a, b *Matrix) *Matrix {
    if a.NRows() != b.NRows() {
        panic(fmt.Errorf("matrix: cannon evaluate sum of matrices, "+
            "because they have different number of rows, %v != %v ", a.NRows(), b.NRows()))
    }
    if a.NColumns() != b.NColumns() {
        panic(fmt.Errorf("matrix: cannon evaluate sum of matrices, "+
            "because they have different number of columns, %v != %v ", a.NColumns(), b.NColumns()))
    }
    body := make([]*vector.Vector, a.NRows())
    for i, row := range b.body {
        body[i] = new(vector.Vector).Xor(a.body[i], row)
    }
    m.body, m.nColumns = body, a.NColumns()
    return m
}

//Dot sets m to "dot" product of matrices a and b, returns m
// m = a.b
func (m *Matrix) Dot(a, b *Matrix) *Matrix {
    if a.NColumns() != b.NRows() {
        panic(fmt.Errorf("matrix: cannon multiplicate matrices, "+
            "because they have wrong dimension, mat.nrows != mat0.nColumns (%v != %v) ", a.NColumns(), b.NRows()))
    }
    body := make([]*vector.Vector, a.NRows())
    for i:=0; i < len(body); i++ {
        body[i] = new(vector.Vector).SetZero(b.NColumns())
        for _, j := range a.body[i].Support() {
            body[i].Xor(body[i], b.body[j])
        }
    }
    if len(body) == 0 {
        return m.SetZero(0,0)
    }
    m.body, m.nColumns = body, body[0].Len()
    return m
}

//IsEqual returns true if m == a
func (m *Matrix) IsEqual(a *Matrix) bool {
   if m.NRows() != a.NRows() {
       return false
   }
   for i, row := range m.body {
       if row.Cmp(a.body[i]) != 0 {
           return false
       }
   }
   return true
}

//IsZero returns true if m is zero matrix
func (m *Matrix) IsZero()  bool {
    return m.IsEqual(new(Matrix).SetZero(m.NRows(), m.NColumns()))
}

//GetRow returns copy of i-th row of m
func (m *Matrix) GetRow(i uint) *vector.Vector {
    if m.NRows() == 0 {
        return new(vector.Vector).SetZero(0)
    }
    return new(vector.Vector).SetV(m.body[i % m.NRows()])
}

//GetCol returns copy of i-th column of m
func (m *Matrix) GetCol(i uint) *vector.Vector {
    if m.NColumns() == 0 || m.NRows() == 0{
        return new(vector.Vector).SetZero(0)
    }
    row := make([]uint8, 0, m.NRows())
    for j := 0; j < int(m.NRows()); j++ {
        row = append(row, m.body[j].Get(i % m.NColumns()))
    }
    return new(vector.Vector).SetBitArray(row)
}

//Solve solves linear equation Ax^T=b^T and returns fundamental system and any solution of it.
// Fundamental system is basis of linear space that consists of such x: Ax^T=0.
// And solution of it is any vector x which satisfied linear system Ax^T=b^T.
func (m *Matrix) Solve(v *vector.Vector) ([]*vector.Vector, *vector.Vector) {
    if m.NRows() != v.Len() {
        return nil, nil
    }
    e := new(Matrix).ConCols(m, new(Matrix).T(new(Matrix).SetV([]*vector.Vector{v})))
    e.Ort(e)
    fund := make([]*vector.Vector, 0, e.NRows()-1)
    var solution *vector.Vector = nil
    for _, row := range e.body {
        if row.Get(row.Len()-1) == 0 {
            fund = append(fund, row.Resize(row, -1))
        } else {
            solution = row.Resize(row, -1)
        }
    }
    if solution == nil {
        solution = new(vector.Vector).SetZero(0)
    }
    return fund, solution
}

//Iter iterates over rows
func (m *Matrix) Iter() <-chan *vector.Vector {
    ch := make(chan *vector.Vector)
    go func() {
        defer close(ch)
        for i := 0; i < len(m.body); i++ {
            ch <- m.GetRow(uint(i))
        }
    }()
    return ch
}

//Bits returns byte array of matrix's bits
func (m *Matrix) Bits() [][]byte {
    b := make([][]byte, len(m.body))
    for i, row := range m.body{
        b[i] = row.Bits()
    }
    return b
}
