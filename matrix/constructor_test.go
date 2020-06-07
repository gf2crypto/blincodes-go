package matrix

import (
    "fmt"
    "math/rand"
    "strings"
    "testing"
)

//TestNoParams tests function New()
func TestNoParams(t *testing.T) {
    if res := checkResult("New()", New(), "", 0, 0); res != "" {
        t.Errorf(res)
    }
}

//TestNewSquare tests function New(n)
func TestNewSquare(t *testing.T) {
    if res := checkResult("New()", New(0), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    mat := "0"
    if res := checkResult("New(1)", New(1), mat, 1, 1); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(11)", New(11), makeZeroMatrix(11, 11), 11, 11); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(64)", New(64), makeZeroMatrix(64, 64), 64, 64); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(97)", New(97), makeZeroMatrix(97, 97), 97, 97); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(213)", New(213), makeZeroMatrix(213, 213), 213, 213); res != "" {
        t.Errorf(res)
    }
}

//TestNewRectangle test to make rectangular matrix using function New(m, n (u)int{, 8, 16, 32})
func TestNewRectangle(t *testing.T) {
    if res := checkResult("New(0, 0)", New(0, 0), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(0, 1)", New(0, 1), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(110, 0)", New(110, 0), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(4, 11)", New(4, 11), makeZeroMatrix(4, 11), 4, 11); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(1, 11)", New(1, 11), makeZeroMatrix(1, 11), 1, 11); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(11, 1)", New(11, 1), makeZeroMatrix(11, 1), 11, 1); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(258, 64)", New(258, 64), makeZeroMatrix(258, 64), 258, 64); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(97, 1035)", New(97, 1035), makeZeroMatrix(97, 1035), 97, 1035); res != "" {
        t.Errorf(res)
    }
    if res := checkResult("New(97, 213)", New(97, 213), makeZeroMatrix(97, 213), 97, 213); res != "" {
        t.Errorf(res)
    }
}

//TestNewFromString test function New(s string) that creates Matrix from string, uses row-separator \n
func TestNewFromString(t *testing.T) {
    if res := checkResult("New(\"\")", New(""), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    _, smat := makeRandomMatrix(1, 1)
    if res := checkResult("New(len(1))", New(smat), smat, 1, 1); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(4, 11)
    if res := checkResult("New(string, 4, 11)", New(smat), smat, 4, 11); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(1, 11)
    if res := checkResult("New(string, 1, 11)", New(smat), smat, 1, 11); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(11, 1)
    if res := checkResult("New(string, 11, 1)", New(smat), smat, 11, 1); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(258, 64)
    if res := checkResult("New(string, 258, 64)", New(smat), smat, 258, 64); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(97, 1035)
    if res := checkResult("New(string, 97, 1035)", New(smat), smat, 97, 1035); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(213, 97)
    if res := checkResult("New(string, 213, 97)", New(smat), smat, 213, 97); res != "" {
        t.Errorf(res)
    }
}

//TestNewFromStrings test function New(s []string) that creates Matrix from string array,
//every element of this array is row
func TestNewFromStrings(t *testing.T) {
    if res := checkResult("New(\"\")", New([]string{""}), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    _, smat := makeRandomMatrix(1, 1)
    if res := checkResult("New(len(1))", New(strings.Split(smat, "\n")), smat, 1, 1); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(4, 11)
    if res := checkResult("New(strings, 4, 11)", New(strings.Split(smat, "\n")), smat, 4, 11); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(1, 11)
    if res := checkResult("New(strings, 1, 11)", New(strings.Split(smat, "\n")), smat, 1, 11); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(11, 1)
    if res := checkResult("New(strings, 11, 1)", New(strings.Split(smat, "\n")), smat, 11, 1); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(258, 64)
    if res := checkResult("New(strings, 258, 64)", New(strings.Split(smat, "\n")), smat, 258, 64); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(97, 1035)
    if res := checkResult("New(strings, 97, 1035)", New(strings.Split(smat, "\n")), smat, 97, 1035); res != "" {
        t.Errorf(res)
    }
    _, smat = makeRandomMatrix(213, 97)
    if res := checkResult("New(strings, 213, 97)", New(strings.Split(smat, "\n")), smat, 213, 97); res != "" {
        t.Errorf(res)
    }
}

//TestNewFromIntegerArray test function New(m, (u)int{, 8, 16, 32}) that create Matrix
//with m rows from integer array, len of this array must be divided to m
func TestNewFromIntegerArray(t *testing.T) {
    if res := checkResult("New([]uint8{})", New(0, []uint8{}), "", 0, 0); res != "" {
        t.Errorf(res)
    }
    mat, smat := makeRandomMatrix(1, 1)
    if res := checkResult("New(len(1))", New(1, mat), smat, 1, 1); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(4, 11)
    if res := checkResult("New([]uint8, 4, 11)", New(4, mat), smat, 4, 11); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(1, 11)
    if res := checkResult("New([]uint8, 1, 11)", New(1, mat), smat, 1, 11); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(11, 1)
    if res := checkResult("New([]uint8, 11, 1)", New(11, mat), smat, 11, 1); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(258, 64)
    if res := checkResult("New([]uint8, 258, 64)", New(258, mat), smat, 258, 64); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(97, 1035)
    if res := checkResult("New([]uint8, 97, 1035)", New(97, mat), smat, 97, 1035); res != "" {
        t.Errorf(res)
    }
    mat, smat = makeRandomMatrix(213, 97)
    if res := checkResult("New([]uint8, 213, 97)", New(213, mat), smat, 213, 97); res != "" {
        t.Errorf(res)
    }
}

//TestIter tests function for iterating over matrix's rows.
func TestIter(t *testing.T) {
    mat := New(6, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
    })
    body := mat.body
    i := 0
    for b := range mat.Iter() {
        if !body[i].Equal(b) {
            t.Errorf("matrix testing: range mat is incorrect, (mat[%d] = %d) != %d",
                i, body[i], b)
        }
        i++
    }
}

func checkResult(method string, mat *Matrix, res string, nrows, ncolumns int) string {
    if fmt.Sprint(mat) != res {
        return fmt.Sprintf("matrix constructor testing:  %v is incorrect, %s != %v)",
            method, mat, res)
    }
    if mat.Nrows() != nrows {
        return fmt.Sprintf("matrix constructor testing:  Nrows of %v is incorrect, %v != %v)",
            method, mat.Nrows(), nrows)
    }
    if mat.Ncolumns() != ncolumns {
        return fmt.Sprintf("matrix constructor testing:  Ncolumns of %v is incorrect, %v != %v)",
            method, mat.Ncolumns(), ncolumns)
    }
    return ""
}

func makeZeroMatrix(m, n int) string {
    mat := ""
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            mat += "0"
        }
        if i < m-1 {
            mat += "\n"
        }
    }
    return mat
}

func makeRandomMatrix(m, n int) ([]uint8, string) {
    mat := make([]uint8, m*n)
    smat := ""
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rand.Intn(2) == 0 {
                mat[i*n+j] = 0
                smat += "0"
            } else {
                mat[i*n+j] = 1
                smat += "1"
            }
        }
        if i < m-1 {
            smat += "\n"
        }
    }
    return mat, smat
}
