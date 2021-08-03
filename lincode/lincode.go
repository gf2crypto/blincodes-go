package lincode

import (
	"bytes"
	"fmt"
	"github.com/gf2crypto/blincodes-go/matrix"
	"github.com/gf2crypto/blincodes-go/vector"
	"math"
	"sort"
)

type GenericLinCode struct {
	basis map[uint]*vector.Vector
	n     uint
}

func (c *GenericLinCode) String() string {
	var buf bytes.Buffer
	if len(c.basis) == 0 {
		_, err := fmt.Fprintf(&buf, "-1: %s", new(vector.Vector).SetZero(c.n))
		if err != nil {
			panic(err)
		}
		return buf.String()
	}
	lenSep := uint(math.Log10(float64(c.n))) + 1
	fs := fmt.Sprintf("%%0%dd: %%s\n", lenSep)
	keys := make([]int, 0, len(c.basis))
	for k := range c.basis {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		_, err := fmt.Fprintf(&buf, fs, k, c.basis[uint(k)])
		if err != nil {
			panic(err)
		}
	}
	s := buf.String()
	return s[:len(s)-1]
}

func (c *GenericLinCode) GetBasis() []*vector.Vector {
	basis := make([]*vector.Vector, 0, len(c.basis))
	if len(c.basis) == 0 {
		basis = append(basis, new(vector.Vector).SetZero(c.n))
		return basis
	}
	keys := make([]int, 0, len(c.basis))
	for k := range c.basis {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		basis = append(basis, c.basis[uint(k)])
	}
	return basis
}

func (c *GenericLinCode) N() uint {
	return c.n
}

func (c *GenericLinCode) K() uint {
	return uint(len(c.basis))
}

func (c *GenericLinCode) SetZero(n uint) *GenericLinCode {
	c.basis = make(map[uint]*vector.Vector)
	c.n = n
	return c
}

func (c *GenericLinCode) SetWholeSpace(n uint) *GenericLinCode {
	if n == 0 {
		return c.SetZero(0)
	}
	c.basis = make(map[uint]*vector.Vector)
	c.n = n
	for i := uint(0); i < n; i++ {
		c.basis[i] = new(vector.Vector).SetZero(n)
		c.basis[i] = c.basis[i].SetBit(c.basis[i], i, 1)
	}
	return c
}

func (c *GenericLinCode) SetV(words []*vector.Vector) *GenericLinCode {
	c.basis = make(map[uint]*vector.Vector)
	c.n = 0
	for i, w := range words {
		if i == 0 {
			c.n = w.Len()
		}
		c.appendWord(w)
	}
	return c
}

func (c *GenericLinCode) SetC(a *GenericLinCode) *GenericLinCode {
	if c == a {
		return c
	}
	basis := make(map[uint]*vector.Vector)
	for i, v := range a.basis {
		basis[i] = new(vector.Vector).SetV(v)
	}
	c.n, c.basis = a.n, basis
	return c
}

func (c *GenericLinCode) SetM(m *matrix.Matrix) *GenericLinCode {
	c.n = m.NColumns()
	c.basis = make(map[uint]*vector.Vector)
	for i := uint(0); i < m.NRows(); i++ {
		c.appendWord(m.GetRow(i))
	}
	return c
}

func (c *GenericLinCode) SetRandom(n, k uint) *GenericLinCode {
	return c.SetM(new(matrix.Matrix).SetRandom(k, n))
}

func (c *GenericLinCode) Append(a *GenericLinCode, words []*vector.Vector) *GenericLinCode {
	c.SetC(a)
	for _, w := range words {
		c.appendWord(w)
	}
	return c
}

func (c *GenericLinCode) AppendM(a *GenericLinCode, m *matrix.Matrix) *GenericLinCode {
	c.SetM(m)
	for _, v := range a.GetBasis() {
		c.appendWord(v)
	}
	return c
}

func (c *GenericLinCode) appendWord(word *vector.Vector) *GenericLinCode {
	if c.n != word.Len() {
		msg := "error in lincode.append: expected length of code %s and length of word %s to be the same, " +
			" but got code length is %d, word length is %d"
		panic(fmt.Errorf(msg, c, word, c.n, word.Len()))
	}
	tmp := new(vector.Vector).SetV(word)
	for i, w := range c.basis {
		if word.Get(i) != 0 {
			tmp.Xor(tmp, w)
		}
	}
	j := uint(0) // index of the first 1 in word
	for ; j < word.Len(); j++ {
		if tmp.Get(j) != 0 {
			break
		}
	}
	if j == word.Len() {
		// tmp == 0
		return c
	}
	// tmp != 0
	c.basis[j] = tmp
	for i, w := range c.basis {
		if i == j {
			continue
		}
		if w.Get(j) != 0 {
			c.basis[i] = w.Xor(w, c.basis[j])
		}
	}
	return c
}
