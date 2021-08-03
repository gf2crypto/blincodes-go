package lincode

import (
	"bytes"
	"fmt"
	"github.com/gf2crypto/blincodes-go/matrix"
	"github.com/gf2crypto/blincodes-go/vector"
)

//LinearCode defines the abstract interface of linear block code
type LinearCode interface {
	GetBasis() []*vector.Vector
	N() uint
	K() uint
}

func checkLength(a, b LinearCode, name string) {
	if a.N() != b.N() {
		msg := "%s error: expected linear codes have the same length, " +
			"but a.N()=%d, b.N()=%d, %d != %d"
		panic(fmt.Errorf(msg, name, a.N(), b.N(), a.N(), b.N()))
	}
}

func String(c LinearCode) string {
	var buf bytes.Buffer
	_, err := fmt.Fprint(&buf, "[\n")
	if err != nil {
		panic(err)
	}
	b := c.GetBasis()
	if len(b) == 0 {
		_, err = fmt.Fprintf(&buf, "%s\n", new(vector.Vector).SetZero(c.N()))
		if err != nil {
			panic(err)
		}
	} else {
		for _, v := range b {
			_, err = fmt.Fprintf(&buf, "%s\n", v)
			if err != nil {
				panic(err)
			}
		}
	}
	_, err = fmt.Fprint(&buf, "]")
	if err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}

//HadamardProduct returns component-wise (Shur or Hadamard) product of codes a and b
func HadamardProduct(a, b LinearCode) LinearCode {
	checkLength(a, b, "hadamard product")
	c := new(GenericLinCode).SetZero(a.N())
	for _, v := range a.GetBasis() {
		for _, w := range b.GetBasis() {
			c.Append(c, []*vector.Vector{new(vector.Vector).And(v, w)})
		}
	}
	return c
}

//Dual returns dual of code a
func Dual(a LinearCode) LinearCode {
	g := new(matrix.Matrix).SetV(a.GetBasis())
	g.Ort(g)
	return new(GenericLinCode).SetM(g)
}

//Intersect returns intersection of code a and b
func Intersect(a, b LinearCode) LinearCode {
	checkLength(a, b, "intersect")
	c := new(GenericLinCode).SetZero(a.N())
	c.Append(c, Dual(a).GetBasis())
	c.Append(c, Dual(b).GetBasis())
	return Dual(c)
}

//Sum returns sum of codes a and b: a+b
func Sum(a, b LinearCode) LinearCode {
	checkLength(a, b, "sum")
	c := new(GenericLinCode).SetZero(a.N())
	c.Append(c, a.GetBasis())
	c.Append(c, b.GetBasis())
	return c
}

//Hull returns hull of code a
//Hull is a intersection of Dual(a) and a
func Hull(a LinearCode) LinearCode {
	c := new(GenericLinCode).SetZero(a.N())
	c.Append(c, a.GetBasis())
	c.Append(c, Dual(a).GetBasis())
	return Dual(c)
}

//Puncture evaluates of puncture code.
// Punctured code is code obtaining by set the positions
// with indexes from `ncolumns` of every codeword to zero.
// Punctured code is NOT subcode of original code!
func Puncture(a LinearCode, columns []uint) LinearCode {
	mask := new(vector.Vector).SetSupport(a.N(), columns)
	mask.Not(mask)
	c := new(GenericLinCode).SetZero(a.N())
	for _, v := range a.GetBasis() {
		c.Append(c, []*vector.Vector{new(vector.Vector).And(v, mask)})
	}
	return c
}

//Truncate evaluates of truncated code.
//Truncated code is code obtaining by choose codewords which
//have coordinates with indexes from `columns` is zero.
//Unlike the punctured code truncated code is a subcode of original code.
func Truncate(a LinearCode, columns []uint) LinearCode {
	gen := new(matrix.Matrix).SetV(a.GetBasis())
	gen.Diagonal(gen, columns)
	c := new(GenericLinCode).SetZero(a.N())
	for i := uint(0); i < gen.NRows(); i++ {
		flag := true
		v := gen.GetRow(i)
		for _, j := range columns {
			if j < v.Len() && v.Get(j) != 0 {
				flag = false
				break
			}
		}
		if flag {
			c.Append(c, []*vector.Vector{v})
		}
	}
	return c
}

//Encode encodes the message
func Encode(a LinearCode, v *vector.Vector) *vector.Vector {
	if v.Len() != a.K() {
		msg := "encode error: expected length of vector = %d (== dimension of code), got %d"
		panic(fmt.Errorf(msg, a.K(), v.Len()))
	}
	w := new(vector.Vector).SetZero(a.N())
	aBasis := a.GetBasis()
	for i := uint(0); i < v.Len(); i++ {
		if v.Get(i) == 1 {
			w.Xor(w, aBasis[i])
		}
	}
	return w
}

//IterWords iterates over code words of code a
func IterWords(a LinearCode) <-chan *vector.Vector {
	ch := make(chan *vector.Vector)
	go func() {
		defer close(ch)
		v := new(vector.Vector).SetZero(a.K())
		if v.Len() == 0 {
			return
		}

		for {
			ch <- Encode(a, v)
			flag := false
			for i := uint(0); i < v.Len() && !flag; i++ {
				switch v.Get(i) {
				case 0:
					v.SetBit(v, i, 1)
					flag = true
				case 1:
					v.SetBit(v, i, 0)
				}
			}
			if !flag {
				break
			}
		}
	}()
	return ch
}

//Spectrum returns weight spectrum of code a
func Spectrum(a LinearCode) *map[uint]uint {
	spec := make(map[uint]uint)
	for v := range IterWords(a) {
		spec[v.Wt()]++
	}
	return &spec
}

//D returns code distance of code
func D(a LinearCode) uint {
	d := a.N()
	for v := range IterWords(a) {
		if wt := v.Wt(); wt != 0 && wt < d {
			d = wt
		}
	}
	return d
}

//ParityChecks returns parity-check matrix of code a
func ParityChecks(a LinearCode) *matrix.Matrix {
	return Generator(Dual(a))
}

//Generator returns generator matrix of a
func Generator(a LinearCode) *matrix.Matrix {
	return new(matrix.Matrix).SetV(a.GetBasis())
}

//IsSubset tests if code a is subset of code b
func IsSubset(a, b LinearCode) bool {
	if a.N() != b.N() || a.K() > b.K() {
		return false
	}
	gen := Generator(a)
	pc := ParityChecks(a)
	pc.T(pc)
	if res := new(matrix.Matrix).Dot(gen, pc); res.IsZero() {
		return true
	}
	return false
}

//IsEqual tests if a is equal b or not
func IsEqual(a, b LinearCode) bool {
	if a.K() == b.K() && IsSubset(a, b) {
		return true
	}
	return false
}
