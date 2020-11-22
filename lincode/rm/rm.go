package rm

// import "fmt"
import "container/list"
import "github.com/gf2crypto/blincodes-go/matrix"
import "github.com/gf2crypto/blincodes-go/vector"

//LinearCode defines Reed-Muller Linear Code object
type LinearCode struct {
    R           int
    M           int
    generator   *matrix.Matrix
    parityCheck *matrix.Matrix
}

//New creates the RM(r,m) code
func New(r, m int) *LinearCode {
    if r < 0 {
        r = 0
    }
    if m < 0 {
        m = 0
    }
    if r > m {
        r = m
    }
    return &LinearCode{R: r, M: m, generator: nil, parityCheck: nil}
}

//D returns the code distance of RM code
func (rm *LinearCode) D() int {
    return 1 << (rm.M - rm.R)
}

//K return dimension of linear code
func (rm *LinearCode) K() int {
    return rm.Gen().Nrows()
}

//N return length of linear code
func (rm *LinearCode) N() int {
    return rm.Gen().Ncolumns()
}

//Gen returns generator matrix of Reed-Muller code
func (rm *LinearCode) Gen() *matrix.Matrix {
    if rm.generator != nil {
        return rm.generator
    }
    rm.generator = genrm(rm.R, rm.M)
    return rm.generator
}

//ParityCheck returns parity check matrix of Reed-Muller code
func (rm *LinearCode) ParityCheck() *matrix.Matrix {
    if rm.parityCheck != nil {
        return rm.parityCheck
    }
    rm.parityCheck = genrm(rm.M-rm.R-1, rm.M)
    return rm.parityCheck
}

//genrm constructs generator matrix of Reed-Muller code
func genrm(r, m int) *matrix.Matrix {
    if r < 0 && m >= 0 {
        return matrix.New(1, 1<<m)
    }
    a := make([]uint8, 1<<m)
    for i := 0; i < len(a); i++ {
        a[i] = 1
    }
    rows := list.New()
    rows.PushBack([2]interface{}{-1, vector.New(a)})
    if r >= 1 {
        monoms := make([](*vector.Vector), 0, m)
        for i := 0; i < m; i++ {
            r := 0
            b := 0
            for j := 0; j < 1<<m; j++ {
                if (j >> (m - i - 1)) != r {
                    r++
                    b = (b + 1) % 2
                }
                a[j] = uint8(b)
            }
            v := vector.New(a)
            monoms = append(monoms, v)
            rows.PushBack([2]interface{}{i, v})
        }
        lStart, lEnd := rows.Front(), rows.Back()
        for i := 1; i < r; i++ {
            for e := lStart.Next(); e != lEnd; e = e.Next() {
                for j := e.Value.([2]interface{})[0].(int) + 1; j < m; j++ {
                    rows.PushBack([2]interface{}{j, monoms[j].And(e.Value.([2]interface{})[1].(*vector.Vector))})
                }
            }
            for j := lEnd.Value.([2]interface{})[0].(int) + 1; j < m; j++ {
                rows.PushBack([2]interface{}{j, monoms[j].And(lEnd.Value.([2]interface{})[1].(*vector.Vector))})
            }
            lStart = lEnd
            lEnd = rows.Back()
        }
    }
    body := make([](*vector.Vector), 0, rows.Len())
    for e := rows.Front(); e != nil; e = e.Next() {
        body = append(body, e.Value.([2]interface{})[1].(*vector.Vector))
    }
    return matrix.New(body)
}
