package lincode

import (
	"github.com/gf2crypto/blincodes-go/vector"
	"testing"
)

//TestHadamardProd tests evaluation of Hadamard product of codes.
func TestHadamardProd(t *testing.T) {
	tests := []struct {
		a, b, expected *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			})}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1,
			}),
		})},
	}
	for _, test := range tests {
		c := HadamardProduct(test.a, test.b)
		if !IsEqual(c, test.expected) {
			msg := "HadamardProduct(\n%s,\n%s)\n=\n%s,\nexpected\n%s\n"
			t.Errorf(msg, test.a, test.b, c, test.expected)
		}
	}
}

//TestIntersect tests evaluation of code intersection.
func TestIntersect(t *testing.T) {
	tests := []struct {
		a, b, expected *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			})}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		})},
	}
	for _, test := range tests {
		c := Intersect(test.a, test.b)
		if !IsEqual(c, test.expected) {
			msg := "Intersect(\n%s,\n%s)\n=\n%s,\nexpected\n%s\n"
			t.Errorf(msg, test.a, test.b, c, test.expected)
		}
	}
}

//TestSum tests evaluation of code sum.
func TestSum(t *testing.T) {
	tests := []struct {
		a, b, expected *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			})}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		})},
	}
	for _, test := range tests {
		c := Sum(test.a, test.b)
		if !IsEqual(c, test.expected) {
			msg := "Sum(\n%s,\n%s)\n=\n%s,\nexpected\n%s\n"
			t.Errorf(msg, test.a, test.b, c, test.expected)
		}
	}
}

//TestHull tests evaluation of code's hull.
func TestHull(t *testing.T) {
	tests := []struct {
		a, expected *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		})},
	}
	for _, test := range tests {
		c := Hull(test.a)
		if !IsEqual(c, test.expected) {
			msg := "Hull(\n%s)\n=\n%s,\nexpected\n%s\n"
			t.Errorf(msg, test.a, c, test.expected)
		}
	}
}

//TestPuncture tests evaluation of punctured code.
func TestPuncture(t *testing.T) {
	tests := []struct {
		a       *GenericLinCode
		columns []uint
		want    *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), []uint{0, 4, 8, 9, 15}, new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0,
			})})},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), []uint{0, 1, 2, 3, 4, 5, 8, 9, 14, 15},
			new(GenericLinCode).SetV([]*vector.Vector{
				new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
				}), new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0,
				}), new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0,
				}), new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0,
				})})},
	}
	for _, test := range tests {
		c := Puncture(test.a, test.columns)
		if !IsEqual(c, test.want) {
			msg := "Puncture(\n%s,\n%v)=\n%s,\nwant\n%s"
			t.Errorf(msg, test.a, test.columns, c, test.want)
		}
	}
}

//TestTruncate tests evaluation of truncated code.
func TestTruncate(t *testing.T) {
	tests := []struct {
		a       *GenericLinCode
		columns []uint
		want    *GenericLinCode
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), []uint{0, 4, 8, 9, 15}, new(GenericLinCode).SetZero(16)},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), []uint{0, 1, 2, 3},
			new(GenericLinCode).SetV([]*vector.Vector{
				new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
				}), new(vector.Vector).SetBitArray([]byte{
					0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
				})})},
	}
	for _, test := range tests {
		c := Truncate(test.a, test.columns)
		if !IsEqual(c, test.want) {
			msg := "Truncate(\n%s,\n%v)=\n%s,\nwant\n%s"
			t.Errorf(msg, test.a, test.columns, c, test.want)
		}
	}
}

//TestEncode tests to encode of message.
func TestEncode(t *testing.T) {
	tests := []struct {
		a       *GenericLinCode
		m, want *vector.Vector
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), new(vector.Vector).SetBitArray([]byte{
			1, 1, 1, 1, 1,
		}), new(vector.Vector).SetBitArray([]byte{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}),
		},
	}
	for _, test := range tests {
		v := Encode(test.a, test.m)
		if v.Cmp(test.want) != 0 {
			msg := "Encode(\n%s,\n%s)=\n%s,\nwant %s"
			t.Errorf(msg, test.a, test.m, v, test.want)
		}
	}
}

func TestIterWords(t *testing.T) {
	tests := []struct {
		a    *GenericLinCode
		want []*vector.Vector
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), []*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}),
		}},
	}
	for _, test := range tests {
		i := 0
		for v := range IterWords(test.a) {
			if v.Cmp(test.want[i]) != 0 {
				msg := "IterWords(\n%s): word[%d] = %s, want %s"
				t.Errorf(msg, test.a, i, v, test.want[i])
			}
			i++
		}
	}
}

//TestSpectrum tests to evaluate of code spectrum.
func TestSpectrum(t *testing.T) {
	tests := []struct {
		a    *GenericLinCode
		want map[uint]uint
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), map[uint]uint{
			0: 1, 8: 30, 16: 1}},
	}
	cmpArrays := func(a, b map[uint]uint) bool {
		if len(a) != len(b) {
			return false
		}
		for i, el := range a {
			if b[i] != el {
				return false
			}
		}
		return true
	}
	for _, test := range tests {
		if sp := Spectrum(test.a); !cmpArrays(*sp, test.want) {
			msg := "Spectrum(\n%s)=\n=%v, want %v"
			t.Errorf(msg, test.a, *sp, test.want)
		}
	}
}

//TestD tests to evaluate of code distance.
func TestD(t *testing.T) {
	tests := []struct {
		a    *GenericLinCode
		want uint
	}{
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
				1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
			}),
		}), 8},
		{new(GenericLinCode).SetV([]*vector.Vector{
			new(vector.Vector).SetBitArray([]byte{
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
				0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1,
			}), new(vector.Vector).SetBitArray([]byte{
				0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1,
			}),
		}), 4},
	}

	for _, test := range tests {
		if d := D(test.a); d != test.want {
			msg := "D(\n%s)=%d, want %d"
			t.Errorf(msg, test.a, d, test.want)
		}
	}
}
