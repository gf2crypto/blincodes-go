package rm

// import "fmt"
import "container/list"
import "github.com/gf2crypto/blincodes-go/vector"

//Code defines Reed-Muller Linear Code object
type Code struct {
	r     uint
	m     uint
	basis []*vector.Vector
}

//Set sets c to the RM(r,m) code
func (c *Code) Set(r, m uint) *Code {
	if r > m {
		r = m
	}
	return &Code{r: r, m: m, basis: genrm(r, m)}
}

//D returns the code distance of RM code
func (c *Code) D() uint {
	return 1 << (c.m - c.r)
}

//K returns dimension of RM code
func (c *Code) K() uint {
	return uint(len(c.basis))
}

//N returns length of RM code
func (c *Code) N() uint {
	return c.basis[0].Len()
}

//GetBasis returns basis of Reed-Muller code
func (c *Code) GetBasis() []*vector.Vector {
	return c.basis
}

//genrm constructs basis of Reed-Muller code
func genrm(r, m uint) []*vector.Vector {
	a := make([]uint8, 1<<m)
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	rows := list.New()
	rows.PushBack([2]interface{}{-1, new(vector.Vector).SetBitArray(a)})
	if r >= 1 {
		monoms := make([]*vector.Vector, 0, m)
		for i := uint(0); i < m; i++ {
			r := 0
			b := 0
			for j := 0; j < 1<<m; j++ {
				if (j >> (m - i - 1)) != r {
					r++
					b = (b + 1) % 2
				}
				a[j] = uint8(b)
			}
			v := new(vector.Vector).SetBitArray(a)
			monoms = append(monoms, v)
			rows.PushBack([2]interface{}{i, v})
		}
		lStart, lEnd := rows.Front(), rows.Back()
		for i := uint(1); i < r; i++ {
			for e := lStart.Next(); e != lEnd; e = e.Next() {
				for j := e.Value.([2]interface{})[0].(uint) + 1; j < m; j++ {
					rows.PushBack([2]interface{}{
						j, new(vector.Vector).And(monoms[j],
							e.Value.([2]interface{})[1].(*vector.Vector))})
				}
			}
			for j := lEnd.Value.([2]interface{})[0].(uint) + 1; j < m; j++ {
				rows.PushBack([2]interface{}{
					j, new(vector.Vector).And(monoms[j],
						lEnd.Value.([2]interface{})[1].(*vector.Vector))})
			}
			lStart = lEnd
			lEnd = rows.Back()
		}
	}
	basis := make([]*vector.Vector, 0, rows.Len())
	for e := rows.Front(); e != nil; e = e.Next() {
		basis = append(basis, e.Value.([2]interface{})[1].(*vector.Vector))
	}
	return basis
}
