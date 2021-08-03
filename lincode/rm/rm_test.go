package rm

import (
	"fmt"
	"github.com/gf2crypto/blincodes-go/lincode"
	"github.com/gf2crypto/blincodes-go/vector"
	"testing"
)

func TestGeneration(t *testing.T) {
	tests := []struct {
		r, m uint
		want struct {
			n, k, d uint
			basis   []*vector.Vector
		}
	}{
		{0, 0, struct {
			n, k, d uint
			basis   []*vector.Vector
		}{1, 1, 1, []*vector.Vector{new(vector.Vector).SetUnits(1)}}},
		{2, 0, struct {
			n, k, d uint
			basis   []*vector.Vector
		}{1, 1, 1, []*vector.Vector{new(vector.Vector).SetUnits(1)}}},
		{0, 4, struct {
			n, k, d uint
			basis   []*vector.Vector
		}{16, 1, 16, []*vector.Vector{new(vector.Vector).SetUnits(16)}}},
		{1, 4, struct {
			n, k, d uint
			basis   []*vector.Vector
		}{16, 5, 8, []*vector.Vector{new(vector.Vector).SetBitArray([]byte{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		})}}},
		{2, 4, struct {
			n, k, d uint
			basis   []*vector.Vector
		}{16, 11, 4, []*vector.Vector{new(vector.Vector).SetBitArray([]byte{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
		}),
		}}},
	}
	cmpArray := func(a, b []*vector.Vector) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v.Cmp(b[i]) != 0 {
				fmt.Println(v, b[i])
				return false
			}
		}
		return true
	}
	vToStr := func(v []*vector.Vector) string {
		s := "[\n"
		for _, w := range v {
			s += fmt.Sprintf("    [ %s ],\n", w)
		}
		s += "]"
		return s
	}
	for _, test := range tests {
		c := new(Code).Set(test.r, test.m)
		if n := c.N(); n != test.want.n {
			msg := "RM(%d, %d).N() = %d, want %d"
			t.Errorf(msg, test.r, test.m, n, test.want.n)
		}
		if k := c.K(); k != test.want.k {
			msg := "RM(%d, %d).K() = %d, want %d"
			t.Errorf(msg, test.r, test.m, k, test.want.k)
		}
		if d := c.D(); d != test.want.d {
			msg := "RM(%d, %d).N() = %d, want %d"
			t.Errorf(msg, test.r, test.m, d, test.want.d)
		}
		if basis := c.GetBasis(); !cmpArray(basis, test.want.basis) {
			msg := "RM(%d, %d).GetBasis() = \n%v,\nwant\n%v"
			t.Errorf(msg, test.r, test.m, vToStr(basis), vToStr(test.want.basis))
		}
	}
}

func TestDuality(t *testing.T) {
	tests := []struct {
		r, m uint
		want lincode.LinearCode
	}{
		{1, 2, new(Code).Set(0, 2)},
		{1, 4, new(Code).Set(2, 4)},
		{2, 5, new(Code).Set(2, 5)},
		{2, 7, new(Code).Set(4, 7)},
		{3, 8, new(Code).Set(4, 8)},
		{3, 9, new(Code).Set(5, 9)},
		{3, 10, new(Code).Set(6, 10)},
		{4, 11, new(Code).Set(6, 11)},
		{5, 5, new(lincode.GenericLinCode).SetZero(32)},
		{9, 10, new(Code).Set(0, 10)},
	}
	for _, test := range tests {
		dual := lincode.Dual(new(Code).Set(test.r, test.m))
		if !lincode.IsEqual(dual, test.want) {
			msg := "Dual(RM(%d, %d)) =\n%s,\nwant\n%s"
			t.Errorf(msg, test.r, test.m, dual, lincode.String(test.want))
		}
	}
}

func TestHadamardProduct(t *testing.T) {
	tests := []struct {
		code1 struct {
			r, m uint
		}
		code2 struct {
			r, m uint
		}
		want lincode.LinearCode
	}{
		{struct {
			r, m uint
		}{1, 4}, struct {
			r, m uint
		}{2, 4}, new(Code).Set(3, 4)},
		{struct {
			r, m uint
		}{3, 8}, struct {
			r, m uint
		}{3, 8}, new(Code).Set(6, 8)},
		{struct {
			r, m uint
		}{3, 10}, struct {
			r, m uint
		}{2, 10}, new(Code).Set(5, 10)},
		{struct {
			r, m uint
		}{4, 10}, struct {
			r, m uint
		}{4, 10}, new(Code).Set(8, 10)},
	}
	for _, test := range tests {
		hp := lincode.HadamardProduct(new(Code).Set(test.code1.r, test.code1.m),
			new(Code).Set(test.code2.r, test.code2.m))
		if !lincode.IsEqual(hp, test.want) {
			msg := "RM(%d, %d) o RM(%d, %d) =\n%s,\nwant\n%s"
			t.Errorf(msg, test.code1.r, test.code1.m,
				test.code2.r, test.code2.m, hp, lincode.String(test.want))
		}
	}
}
