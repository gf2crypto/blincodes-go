package matrix_test

import (
	"crypto/rand"
	"fmt"
	"github.com/gf2crypto/blincodes-go/matrix"
	"github.com/gf2crypto/blincodes-go/vector"
	"strings"
	"testing"
)

func getZeroMatrix(m, n uint) [][]byte {
	if m == 0 || n == 0 {
		return make([][]byte, 0)
	}
	mat := make([][]byte, m)
	for i := uint(0); i < m; i++ {
		mat[i] = make([]byte, n)
	}
	return mat
}

func isEqualArrays(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func getRandMat(m, n uint) [][]byte {
	if m == 0 || n == 0 {
		return make([][]byte, 0)
	}
	b := make([][]byte, m)
	for i := uint(0); i < m; i++ {
		b[i] = make([]byte, n)
		_, err := rand.Read(b[i])
		if err != nil {
			panic(err)
		}
		for j := uint(0); j < n; j++ {
			b[i][j] &= 1
		}
	}
	return b
}

func bArrayToS(a [][]byte) string {
	s := ""
	if len(a) == 0 {
		return ""
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			s += string('0' + a[i][j])
		}
		if i != len(a)-1 {
			s += "\n"
		}
	}
	return s
}

//func checkResult(method string, mat *matrix.Matrix, res string, nRows, nColumns uint) string {
//   //fmt.Println("Matrix: ")
//   //fmt.Println(mat)
//   //fmt.Println(mat.NRows(), mat.NColumns())
//   //fmt.Println("RES:")
//   //fmt.Println(len(res))
//   if fmt.Sprint(mat) != res {
//       return fmt.Sprintf("matrix testing:  %s is incorrect, expected\n%s,\nbut got\n%s)",
//           method, res, mat)
//   }
//   if mat.NRows() != nRows {
//       return fmt.Sprintf("matrix testing:  NRows of %v is incorrect, %v != %v)",
//           method, mat.NRows(), nRows)
//   }
//   if mat.NColumns() != nColumns {
//       return fmt.Sprintf("matrix testing:  NColumns of %v is incorrect, %v != %v)",
//           method, mat.NColumns(), nColumns)
//   }
//   return ""
//}

func TestMatrix_New(t *testing.T) {
	m := matrix.New()
	if r := m.NRows(); r != 0 {
		msg := fmt.Sprintf("matrix.New() has %d rows, want 0", r)
		t.Errorf(msg)
	}
	if r := m.NColumns(); r != 0 {
		msg := fmt.Sprintf("matrix.New() has %d columns, want 0", r)
		t.Errorf(msg)
	}
}

func TestMatrix_SetZero(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0}, {1, 1}, {11, 11},
		{64, 64}, {97, 97}, {213, 213},
		{0, 1}, {110, 0}, {4, 11},
		{1, 11}, {11, 1}, {258, 64},
		{97, 1035}, {97, 213}, {1024, 512},
		{512, 1024},
	}
	for _, tCase := range testCases {
		v := new(matrix.Matrix).SetZero(tCase.nRows, tCase.nColumns)
		if s, want := v.Bits(), getZeroMatrix(tCase.nRows, tCase.nColumns); !isEqualArrays(s, want) {
			msg := fmt.Sprintf("matrix.SetZero(%d, %d)=\n%v, want\n%v",
				tCase.nRows, tCase.nColumns, s, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_SetUnit(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0}, {1, 1}, {11, 11},
		{64, 64}, {97, 97}, {213, 213},
		{0, 1}, {110, 0}, {4, 11},
		{1, 11}, {11, 1}, {258, 64},
		{97, 1035}, {97, 213}, {1024, 512},
		{512, 1024},
	}
	for _, tCase := range testCases {
		v := new(matrix.Matrix).SetUnit(tCase.nRows, tCase.nColumns)
		want := getZeroMatrix(tCase.nRows, tCase.nColumns)
		for i := 0; i < len(want); i++ {
			for j := 0; j < len(want[0]); j++ {
				want[i][j] = 1
			}
		}
		if s := v.Bits(); !isEqualArrays(s, want) {
			msg := fmt.Sprintf("matrix.SetUnit(%d, %d)=\n%v, want\n%v",
				tCase.nRows, tCase.nColumns, s, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_SetV(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
		nTests   uint
	}{
		{0, 0, 1}, {1, 1, 3},
		{11, 11, 10}, {64, 64, 100},
		{97, 97, 100}, {213, 213, 100},
		{0, 1, 1}, {110, 0, 1},
		{4, 11, 100}, {1, 11, 10},
		{11, 1, 10}, {258, 64, 100},
		{97, 1035, 10}, {97, 213, 10},
		{1024, 512, 10}, {512, 1024, 5},
	}
	for _, tCase := range testCases {
		for i := uint(0); i < tCase.nTests; i++ {
			mat := getRandMat(tCase.nRows, tCase.nColumns)
			vmat := make([]*vector.Vector, 0, tCase.nRows)
			for j := uint(0); j < tCase.nRows && len(mat) > 0; j++ {
				vmat = append(vmat, new(vector.Vector).SetBitArray(mat[j]))
			}
			m := new(matrix.Matrix).SetBits(mat)
			if get := m.Bits(); !isEqualArrays(get, mat) {
				msg := fmt.Sprintf("matrix.SetBits(%v)=\n%v,\nwant\n%v",
					mat, get, mat)
				t.Errorf(msg)
			}
		}
	}
}

func TestMatrix_SetBytes(t *testing.T) {
	b := [][]byte{
		{0x01, 0x02, 0x03},
		{0xAB, 0xBC, 0xFD},
		{},
		{0x76, 0x57, 0xAA, 0xED, 0xFF},
	}
	testCases := []struct {
		nColumns uint
		want     string
	}{
		{5, "00000\n" +
			"10101\n" +
			"00000\n" +
			"01110"},
		{24, "000000010000001000000011\n" +
			"101010111011110011111101\n" +
			"000000000000000000000000\n" +
			"011101100101011110101010"},
		{30, "000000010000001000000011000000\n" +
			"101010111011110011111101000000\n" +
			"000000000000000000000000000000\n" +
			"011101100101011110101010111011"},
	}

	for _, tCase := range testCases {
		a := new(matrix.Matrix).SetBytes(b, tCase.nColumns)
		if s := a.String(); s != tCase.want {
			msg := fmt.Sprintf("matrix.SetBytes(%v, %d)=\n%s,\nwant\n%s",
				b, tCase.nColumns, s, tCase.want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_SetBits(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
		nTests   uint
	}{
		{0, 0, 1},
		{1, 1, 1},
		{5, 5, 10},
		{27, 67, 10},
		{67, 27, 10},
		{254, 512, 10},
		{512, 1024, 5},
		{1024, 512, 5},
		{7123, 8761, 1},
		{8543, 1247, 1},
		{1, 1024, 10},
		{1024, 1, 10},
	}
	for _, tCase := range testCases {
		for i := uint(0); i < tCase.nTests; i++ {
			mat := getRandMat(tCase.nRows, tCase.nColumns)
			m := new(matrix.Matrix).SetBits(mat)
			if get := m.Bits(); !isEqualArrays(get, mat) {
				msg := fmt.Sprintf("matrix.SetBits(%b)=\n%v,\nwant\n%v",
					mat, get, mat)
				t.Errorf(msg)
			}
		}
	}
}

func TestMatrix_SetIdentity(t *testing.T) {
	testCases := []struct {
		dim uint
	}{
		{0}, {1}, {2}, {3}, {10}, {17}, {20}, {32},
		{121}, {276}, {333}, {512}, {1023}, {1024}, {1025},
	}
	for _, tCase := range testCases {
		m := new(matrix.Matrix).SetIdentity(tCase.dim)
		want := getZeroMatrix(tCase.dim, tCase.dim)
		for i := uint(0); i < tCase.dim; i++ {
			want[i][i] = 1
		}
		if get := m.Bits(); !isEqualArrays(get, want) {
			msg := fmt.Sprintf("matrix.SetIdentity(%d)=\n%v,\nwant\n%v",
				tCase.dim, get, want)
			t.Errorf(msg)
		}
	}
}

//
func TestMatrix_SetRandom(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
		nSamples uint
	}{
		{0, 0, 1}, {1, 1, 1},
		{5, 5, 5}, {100, 100, 10},
		{512, 512, 20}, {237, 237, 20},
		{1027, 1027, 5}, {3000, 3000, 5},
		{1, 2, 2}, {5, 10, 10},
		{100, 200, 100}, {512, 1024, 20},
		{237, 474, 5}, {1027, 2054, 10},
		{531, 1895, 100}, {2, 1, 2},
		{10, 5, 5}, {200, 100, 100},
		{1024, 512, 10}, {474, 237, 100},
		{2054, 1027, 5}, {1895, 531, 100},
		{1, 1000, 5}, {1000, 1, 5},
	}
	for _, tCase := range testCases {
		samples := make([]*matrix.Matrix, 0, tCase.nSamples)
		for i := uint(0); i < tCase.nSamples; i++ {
			samples = append(samples, new(matrix.Matrix).SetRandom(tCase.nRows, tCase.nColumns))
		}
		for i, m := range samples {
			if m.NRows() != tCase.nRows {
				msg := fmt.Sprintf("Random(%d, %d)=\n%s,\nNRows is incorrect, got %d, want %d)",
					tCase.nRows, tCase.nColumns, m, m.NRows(), tCase.nRows)
				t.Errorf(msg)
			}
			if m.NColumns() != tCase.nColumns {
				msg := fmt.Sprintf("Random(%d, %d)=\n%s,\nNColumns is incorrect, got %d, want %d)",
					tCase.nRows, tCase.nColumns, m, m.NColumns(), tCase.nColumns)
				t.Errorf(msg)
			}
			for j := i + 1; j < len(samples); j++ {
				if m.IsEqual(samples[j]) {
					msg := fmt.Sprintf("Random(%d, %d): found two coincide matrixes, "+
						"sample[%d]==sample[%d]==\n%s)",
						tCase.nRows, tCase.nColumns, i, j, m)
					t.Errorf(msg)
				}
			}
		}
	}
}

func TestMatrix_SetStrings(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
		nTests   uint
	}{
		{0, 0, 1}, {1, 1, 1},
		{5, 5, 5}, {27, 67, 5},
		{67, 27, 1}, {254, 512, 1},
		{512, 254, 1},
		{1, 1024, 1}, {1024, 1, 1},
	}
	for _, tCase := range testCases {
		for i := uint(0); i < tCase.nTests; i++ {
			mat := getRandMat(tCase.nRows, tCase.nColumns)
			s := strings.Split(bArrayToS(mat), "\n")
			m := new(matrix.Matrix).SetStrings(s)
			if get := m.Bits(); !isEqualArrays(get, mat) {
				msg := fmt.Sprintf("matrix.SetStrings(%v)=\n%v,\n want\n%v",
					s, get, mat)
				t.Errorf(msg)
			}
		}
	}
}

func TestMatrix_Parse(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
		nTests   uint
	}{
		{0, 0, 1}, {1, 1, 1},
		{5, 5, 5}, {27, 67, 5},
		{67, 27, 1}, {254, 512, 1},
		{512, 254, 1},
		{1, 1024, 1}, {1024, 1, 1},
	}
	for _, tCase := range testCases {
		for i := uint(0); i < tCase.nTests; i++ {
			mat := getRandMat(tCase.nRows, tCase.nColumns)
			s := bArrayToS(mat)
			m := new(matrix.Matrix).Parse(s)
			if get := m.Bits(); !isEqualArrays(get, mat) {
				msg := fmt.Sprintf("matrix.Parse(%v)=\n%v,\n want\n%v",
					s, get, mat)
				t.Errorf(msg)
			}
		}
	}
}

func TestMatrix_SetPerm(t *testing.T) {
	getPermMat := func(p []uint) [][]byte {
		b := make([][]byte, len(p))
		for i := 0; i < len(p); i++ {
			b[i] = make([]byte, len(p))
			for j := 0; j < len(p); j++ {
				b[i][j] = 0
			}
		}
		for i := 0; i < len(p); i++ {
			b[p[i]%uint(len(p))][i] = 1
		}
		return b
	}
	for i := 0; i < 1025; i++ {
		p := matrix.GetRandomPerm(uint(i))
		want := getPermMat(p)
		a := new(matrix.Matrix).SetPerm(p)
		if get := a.Bits(); !isEqualArrays(get, want) {
			msg := fmt.Sprintf("SetPerm(%v)=\n%s,\nwant\n%v",
				p, get, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_SetPermL(t *testing.T) {
	getPermMat := func(p []uint) [][]byte {
		b := make([][]byte, len(p))
		for i := 0; i < len(p); i++ {
			b[i] = make([]byte, len(p))
			for j := 0; j < len(p); j++ {
				b[i][j] = 0
			}
		}
		for i := 0; i < len(p); i++ {
			b[i][p[i]%uint(len(p))] = 1
		}
		return b
	}
	for i := 0; i < 1025; i++ {
		p := matrix.GetRandomPerm(uint(i))
		want := getPermMat(p)
		a := new(matrix.Matrix).SetPermL(p)
		if get := a.Bits(); !isEqualArrays(get, want) {
			msg := fmt.Sprintf("SetPermL(%v)=\n%s,\nwant\n%v",
				p, get, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Rank(t *testing.T) {
	evalRank := func(mat [][]byte) uint {
		rank := uint(0)
		m := uint(len(mat))
		if m == 0 {
			return 0
		}
		n := uint(len(mat[0]))
		for j := uint(0); j < n; j++ {
			k := rank
			for ; k < m; k++ {
				if mat[k][j] != 0 {
					break
				}
			}
			if k == m {
				continue
			}
			if k != rank {
				for l := uint(0); l < n; l++ {
					mat[rank][l] ^= mat[k][l]
				}
			}
			for k = rank + 1; k < m; k++ {
				if mat[k][j] == 0 {
					continue
				}
				for l := uint(0); l < n; l++ {
					mat[k][l] ^= mat[rank][l]
				}
			}
			rank++
		}
		return rank
	}
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0},
		{1, 1},
		{10, 0},
		{0, 10},
		{5, 5},
		{7, 4},
		{5, 7},
		{1, 2},
		{5, 10},
		{2, 1},
		{10, 5},
		{100, 100},
		{512, 512},
		{237, 237},
		{1027, 1027},
		{100, 200},
		{512, 1024},
		{237, 474},
		{531, 1895},
		{200, 100},
		{1024, 512},
		{474, 237},
		{1895, 531},
		{1, 1000},
		{1000, 1},
	}
	for _, tCase := range testCases {
		m := getRandMat(tCase.nRows, tCase.nColumns)
		if get, want := new(matrix.Matrix).SetBits(m).Rank(), evalRank(m); get != want {
			s := fmt.Sprintf("Rank(%v) == %d, but want %d",
				m, get, want)
			t.Errorf(s)
		}
	}
}

func TestMatrix_Echelon_Basic(t *testing.T) {
	testCases := []struct {
		mat  *matrix.Matrix
		want *matrix.Matrix
	}{
		{new(matrix.Matrix).SetZero(5, 7), new(matrix.Matrix).SetZero(5, 7)},
		{new(matrix.Matrix).SetZero(7, 5), new(matrix.Matrix).SetZero(7, 5)},
		{new(matrix.Matrix).SetZero(0, 0), new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 1, 1, 1},
			{0, 1, 0, 1, 1},
			{0, 0, 1, 0, 1},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 1, 1, 1},
			{0, 1, 0, 1, 1},
			{0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 1, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 1, 1, 1, 0},
			{0, 1, 0, 1, 1, 0},
			{0, 0, 1, 0, 1, 0},
			{0, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		})},
	}
	for _, tCase := range testCases {
		msg := fmt.Sprintf("Echelon(%s)=\n", tCase.mat)
		tCase.mat.Echelon(tCase.mat, nil)
		if !tCase.mat.IsEqual(tCase.want) {
			msg += fmt.Sprintf("%s,\nwant\n%s", tCase.mat, tCase.want)
		}
	}
}

func TestMatrix_Echelon(t *testing.T) {
	evalEchelon := func(mat [][]byte) [][]byte {
		rank := uint(0)
		m := uint(len(mat))
		if m == 0 {
			return mat
		}
		n := uint(len(mat[0]))
		for j := uint(0); j < n; j++ {
			k := rank
			for ; k < m; k++ {
				if mat[k][j] != 0 {
					break
				}
			}
			if k == m {
				continue
			}
			if k != rank {
				for l := uint(0); l < n; l++ {
					mat[rank][l] ^= mat[k][l]
				}
			}
			for k = rank + 1; k < m; k++ {
				if mat[k][j] == 0 {
					continue
				}
				for l := uint(0); l < n; l++ {
					mat[k][l] ^= mat[rank][l]
				}
			}
			rank++
		}
		return mat
	}
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0}, {1, 1}, {5, 5}, {7, 4},
		{5, 7}, {1, 2}, {5, 10},
		{2, 1}, {10, 5}, {100, 100},
		{512, 512}, {237, 237}, {1027, 1027},
		{100, 200}, {512, 1024}, {237, 474},
		{531, 1895}, {200, 100}, {1024, 512},
		{474, 237}, {2054, 1027}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	for _, tCase := range testCases {
		a := getRandMat(tCase.nRows, tCase.nColumns)
		m := new(matrix.Matrix).SetBits(a)
		msg := fmt.Sprintf("Echelon(%s)=\n", m)
		m.Echelon(m, nil)
		want := new(matrix.Matrix).SetBits(evalEchelon(a))
		if !m.IsEqual(want) {
			msg += fmt.Sprintf("%s,\nwant\n%s",
				m, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Echelon_Partial(t *testing.T) {
	testCases := []struct {
		cols []uint
		mat  *matrix.Matrix
		want *matrix.Matrix
	}{
		{[]uint{1, 2, 6}, new(matrix.Matrix).SetZero(5, 7),
			new(matrix.Matrix).SetZero(5, 7)},
		{[]uint{3}, new(matrix.Matrix).SetZero(7, 5),
			new(matrix.Matrix).SetZero(7, 5)},
		{[]uint{2}, new(matrix.Matrix).SetZero(0, 0),
			new(matrix.Matrix).SetZero(0, 0)},
		{[]uint{3},
			new(matrix.Matrix).SetBits([][]byte{
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 0, 1, 1},
				{0, 0, 0, 1},
			}), new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		})},
		{[]uint{3, 3}, new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		})},
		{[]uint{0, 1, 2, 2, 2}, new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{[]uint{0, 3}, new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 1, 0, 0},
			{0, 1, 1, 0},
		})},
		{[]uint{3, 1, 4}, new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{1, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
		})},
		{[]uint{2, 1}, new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 1, 0, 1, 1},
			{1, 0, 0, 1, 0},
			{1, 0, 0, 1, 0},
		})},
		{[]uint{0, 1, 2, 3, 4, 5}, new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 1, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 1, 1, 1, 0},
			{0, 1, 0, 1, 1, 0},
			{0, 0, 1, 0, 1, 0},
			{0, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		})},
	}
	for _, tCase := range testCases {
		msg := fmt.Sprintf("Echelon(%s, %v)=\n", tCase.mat, tCase.cols)
		tCase.mat.Echelon(tCase.mat, tCase.cols)
		if !tCase.mat.IsEqual(tCase.want) {
			msg += fmt.Sprintf("%s,\nwant\n%s",
				tCase.mat, tCase.want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Diagonal_Basic(t *testing.T) {
	testCases := []struct {
		mat  *matrix.Matrix
		want *matrix.Matrix
	}{
		{new(matrix.Matrix).SetZero(5, 7), new(matrix.Matrix).SetZero(5, 7)},
		{new(matrix.Matrix).SetZero(7, 5), new(matrix.Matrix).SetZero(7, 5)},
		{new(matrix.Matrix).SetZero(0, 0), new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 1, 0},
			{0, 1, 0, 1, 1},
			{0, 0, 1, 0, 1},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 1, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), new(matrix.Matrix).SetBits([][]byte{
			{1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		})},
	}
	for _, tCase := range testCases {
		msg := fmt.Sprintf("Diagonal(%s)=\n", tCase.mat)
		tCase.mat.Diagonal(tCase.mat, nil)
		if !tCase.mat.IsEqual(tCase.want) {
			msg += fmt.Sprintf("%s,\nwant\n%s",
				tCase.mat, tCase.want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Diagonal(t *testing.T) {
	evalDiagonal := func(mat [][]byte) [][]byte {
		rank := uint(0)
		m := uint(len(mat))
		if m == 0 {
			return mat
		}
		n := uint(len(mat[0]))
		for j := uint(0); j < n; j++ {
			k := rank
			for ; k < m; k++ {
				if mat[k][j] != 0 {
					break
				}
			}
			if k == m {
				continue
			}
			if k != rank {
				for l := uint(0); l < n; l++ {
					mat[rank][l] ^= mat[k][l]
				}
			}
			for k = 0; k < m; k++ {
				if mat[k][j] == 0 || k == rank {
					continue
				}
				for l := uint(0); l < n; l++ {
					mat[k][l] ^= mat[rank][l]
				}
			}
			rank++
		}
		return mat
	}
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0}, {1, 1}, {5, 5},
		{7, 4}, {5, 7}, {1, 2},
		{5, 10}, {2, 1}, {10, 5},
		{100, 100}, {512, 512}, {237, 237},
		{1027, 1027}, {100, 200}, {512, 1024},
		{237, 474}, {531, 1895}, {200, 100},
		{1024, 512}, {474, 237}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	for _, tCase := range testCases {
		a := getRandMat(tCase.nRows, tCase.nColumns)
		m := new(matrix.Matrix).SetBits(a)
		msg := fmt.Sprintf("Diagonal(%s)=\n", m)
		m.Diagonal(m, nil)
		want := new(matrix.Matrix).SetBits(evalDiagonal(a))
		if get := m.Bits(); !isEqualArrays(a, get) {
			msg += fmt.Sprintf("\n%s,\nwant\n%s",
				m, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Ort(t *testing.T) {
	testCases := []struct {
		mat *matrix.Matrix
	}{
		{new(matrix.Matrix).SetZero(5, 7)},
		{new(matrix.Matrix).SetZero(7, 5)},
		{new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 1, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		})},
	}
	for _, tCase := range testCases {
		msg := fmt.Sprintf("Ort(%s)=\n", tCase.mat)
		c := new(matrix.Matrix).SetM(tCase.mat)
		tCase.mat.Ort(tCase.mat)
		msg += tCase.mat.String()
		if tCase.mat.NColumns() != c.NColumns() {
			msg += fmt.Sprintf("\nort.NColumns == %d, want %d",
				tCase.mat.NColumns(), c.NColumns())
			t.Errorf(msg)
		}
		if want := c.NColumns() - c.Rank(); tCase.mat.NRows() != want && !tCase.mat.IsZero() {
			msg += fmt.Sprintf("\nort.NRows == %d, want %d (Ncolumns - rank)",
				tCase.mat.NRows(), want)
			t.Errorf(msg)
		}
		if mul := new(matrix.Matrix).Dot(c, tCase.mat.T(tCase.mat));
			!mul.IsZero() {
			msg += fmt.Sprintf("\nA*ort^T=\n%s,\nwant %s", mul,
				new(matrix.Matrix).SetZero(c.NRows(), tCase.mat.NColumns()))
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Ort_Random(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{5, 5}, {7, 4}, {5, 7},
		{1, 2}, {5, 10}, {2, 1},
		{10, 5}, {100, 100}, {512, 512},
		{237, 237}, {1027, 1027}, {100, 200},
		{512, 1024}, {237, 474}, {531, 1895},
		{200, 100}, {1024, 512}, {474, 237},
		{1895, 531}, {1, 1000}, {1000, 1},
	}
	for _, tCase := range testCases {
		m := new(matrix.Matrix).SetBits(getRandMat(tCase.nRows, tCase.nColumns))
		msg := fmt.Sprintf("Ort(%s)=\n", m)
		c := new(matrix.Matrix).SetM(m)
		m.Ort(m)
		msg += m.String()
		if m.NColumns() != c.NColumns() {
			msg += fmt.Sprintf("\nort.NColumns == %d, want %d",
				m.NColumns(), c.NColumns())
			t.Errorf(msg)
		}
		if want := c.NColumns() - c.Rank(); m.NRows() != want && !m.IsZero() {
			msg += fmt.Sprintf("\nort.NRows == %d, want %d (Ncolumns - rank)",
				m.NRows(), want)
			t.Errorf(msg)
		}
		if mul := new(matrix.Matrix).Dot(c, m.T(m));
			!mul.IsZero() {
			msg += fmt.Sprintf("\nA*ort^T=\n%s,\nwant %s", mul,
				new(matrix.Matrix).SetZero(c.NRows(), m.NColumns()))
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Inv(t *testing.T) {
	testCases := []struct {
		mat *matrix.Matrix
	}{
		{new(matrix.Matrix).SetZero(5, 7)},
		{new(matrix.Matrix).SetZero(7, 5)},
		{new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 1},
			{0, 0, 0, 1},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0},
			{0, 0, 1, 0, 1},
			{1, 1, 0, 0, 1},
			{1, 1, 1, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{0, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 1, 0},
			{1, 1, 0, 0, 1, 0},
			{1, 1, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 0, 1, 0, 0},
		})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		})},
	}
	for _, tCase := range testCases {
		msg := fmt.Sprintf("Inv(%s)=\n", tCase.mat)
		c := new(matrix.Matrix).SetM(tCase.mat)
		tCase.mat.Inv(tCase.mat)
		msg += tCase.mat.String()
		if tCase.mat.NColumns() != c.NRows() {
			msg += fmt.Sprintf("\ninv.NColumns == %d, want %d",
				tCase.mat.NColumns(), c.NRows())
			t.Errorf(msg)
		}
		if tCase.mat.NRows() != c.NRows() {
			msg += fmt.Sprintf("\ninv.NRows == %d, want %d",
				tCase.mat.NRows(), c.NRows())
			t.Errorf(msg)
		}
		if r := tCase.mat.Rank(); r != c.NRows() {
			msg += fmt.Sprintf("\ninv.Rank == %d, want %d",
				r, c.NRows())
			t.Errorf(msg)
		}
		mul := new(matrix.Matrix).Dot(tCase.mat, c)
		want := new(matrix.Matrix).Diagonal(c, nil)
		if !want.IsEqual(mul) {
			msg += fmt.Sprintf("\ninv * A=\n%s,\nwant %s", mul, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_Inv_Random(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{5, 5}, {7, 4}, {5, 7}, {1, 2},
		{5, 10}, {2, 1}, {10, 5},
		{100, 100}, {512, 512}, {237, 237},
		{1027, 1027}, {100, 200}, {512, 1024},
		{237, 474}, {531, 1895}, {200, 100},
		{1024, 512}, {474, 237}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	for _, tCase := range testCases {
		m := new(matrix.Matrix).SetBits(getRandMat(tCase.nRows, tCase.nColumns))
		msg := fmt.Sprintf("Inv(%s)=\n", m)
		c := new(matrix.Matrix).SetM(m)
		m.Inv(m)
		msg += m.String()
		if m.NColumns() != c.NRows() {
			msg += fmt.Sprintf("\ninv.NColumns == %d, want %d",
				m.NColumns(), c.NRows())
			t.Errorf(msg)
		}
		if m.NRows() != c.NRows() {
			msg += fmt.Sprintf("\ninv.NRows == %d, want %d",
				m.NRows(), c.NRows())
			t.Errorf(msg)
		}
		if r := m.Rank(); r != c.NRows() {
			msg += fmt.Sprintf("\ninv.Rank == %d, want %d",
				r, c.NRows())
			t.Errorf(msg)
		}
		mul := new(matrix.Matrix).Dot(m, c)
		want := new(matrix.Matrix).Diagonal(c, nil)
		if !want.IsEqual(mul) {
			msg += fmt.Sprintf("\ninv * A=\n%s,\nwant %s", mul, want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_SetM(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{5, 5}, {7, 4}, {5, 7}, {1, 2},
		{5, 10}, {2, 1}, {10, 5},
		{100, 100}, {512, 512}, {237, 237},
		{1027, 1027}, {100, 200}, {512, 1024},
		{237, 474}, {531, 1895}, {200, 100},
		{1024, 512}, {474, 237}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	for _, tCase := range testCases {
		bits := getRandMat(tCase.nRows, tCase.nColumns)
		m := new(matrix.Matrix).SetBits(bits)
		a := new(matrix.Matrix).SetM(m)
		msg := fmt.Sprintf("new.SetM(%s)=\n%s", m, a)
		if !m.IsEqual(a) {
			msg += fmt.Sprintf("\nwant\n%s", m)
			t.Errorf(msg)
		}
		m.SetM(m)
		msg = fmt.Sprintf("self.SetM(%s)=\n%s", a, m)
		if !a.IsEqual(m) {
			msg += fmt.Sprintf("\nwant\n%s", m)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_T(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{5, 5}, {7, 4}, {5, 7}, {1, 2},
		{5, 10}, {2, 1}, {10, 5},
		{100, 100}, {512, 512}, {237, 237},
		{1027, 1027}, {100, 200}, {512, 1024},
		{237, 474}, {531, 1895}, {200, 100},
		{1024, 512}, {474, 237}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	nTests := 10
	evalT := func(mat [][]byte) *matrix.Matrix {
		m := uint(len(mat))
		if m == 0 {
			return new(matrix.Matrix).SetZero(0, 0)
		}
		n := uint(len(mat[0]))
		matT := make([][]byte, n)
		for i := uint(0); i < n; i++ {
			matT[i] = make([]byte, m)
			for j := uint(0); j < m; j++ {
				matT[i][j] = mat[j][i]
			}

		}
		return new(matrix.Matrix).SetBits(matT)
	}
	for _, tCase := range testCases {
		for i := 0; i < nTests; i++ {
			bits := getRandMat(tCase.nRows, tCase.nColumns)
			a := new(matrix.Matrix).SetBits(bits)
			msg := fmt.Sprintf("T(%s)=", a)
			a.T(a)
			if want := evalT(bits); !a.IsEqual(want) {
				msg += fmt.Sprintf("\n%s,\nwant %s", a, want)
				t.Errorf(msg)
			}
		}
	}
}

func TestMatrix_SubMatrix(t *testing.T) {
	testCases := []struct {
		mat  *matrix.Matrix
		rows []uint
		cols []uint
		want *matrix.Matrix
	}{
		{new(matrix.Matrix).SetZero(0, 0), []uint{}, []uint{},
			new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetZero(0, 0), []uint{1, 2, 4}, []uint{5, 9, 11, 15},
			new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{1, 2, 4}, []uint{5, 9, 11, 15},
			new(matrix.Matrix).SetBits([][]byte{
				{0, 1, 1, 1},
				{1, 0, 0, 1},
				{1, 1, 1, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{1, 2, 4}, []uint{},
			new(matrix.Matrix).SetBits([][]byte{
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{4, 2, 1, 35}, []uint{},
			new(matrix.Matrix).SetBits([][]byte{
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{}, []uint{5, 9, 11, 15},
			new(matrix.Matrix).SetBits([][]byte{
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 0, 1, 1},
				{1, 1, 1, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{}, []uint{11, 9, 5, 15},
			new(matrix.Matrix).SetBits([][]byte{
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{1, 0, 0, 1},
				{0, 0, 1, 1},
				{1, 1, 1, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{}, []uint{},
			new(matrix.Matrix).SetBits([][]byte{
				{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
				{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			})},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{7, 12, 9, 10}, []uint{2, 6, 7},
			new(matrix.Matrix).SetZero(0, 0)},
		{new(matrix.Matrix).SetBits([][]byte{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		}), []uint{1, 2, 4}, []uint{34, 67, 30},
			new(matrix.Matrix).SetZero(0, 0)},
	}
	for _, tCase := range testCases {
		a := new(matrix.Matrix).SetM(tCase.mat)
		msg := fmt.Sprintf("SubMatrix(%s, %v, %v)=\n", a, tCase.rows, tCase.cols)
		a.SubMatrix(a, tCase.rows, tCase.cols)
		if !a.IsEqual(tCase.want) {
			msg += fmt.Sprintf("%s,\nwant\n%s", a, tCase.want)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_ConRows(t *testing.T) {
	testCases := []struct {
		nRows    uint
		nColumns uint
	}{
		{0, 0}, {5, 5}, {7, 4}, {5, 7},
		{1, 2}, {5, 10}, {2, 1}, {10, 5},
		{100, 100}, {512, 512}, {237, 237},
		{1027, 1027}, {100, 200}, {512, 1024},
		{237, 474}, {531, 1895}, {200, 100},
		{1024, 512}, {474, 237}, {1895, 531},
		{1, 1000}, {1000, 1},
	}
	evalConRows := func(mat1, mat2 [][]byte) *matrix.Matrix {
		m1 := uint(len(mat1))
		if m1 == 0 {
			return new(matrix.Matrix).SetBits(mat2)
		}
		m2 := uint(len(mat2))
		if m2 == 0 {
			return new(matrix.Matrix).SetBits(mat1)
		}
		n := uint(len(mat1[0]))
		mat := make([][]byte, m1+m2)
		for i := uint(0); i < m1+m2; i++ {
			mat[i] = make([]byte, n)
			if i < m1 {
				copy(mat[i], mat1[i])
			} else {
				copy(mat[i], mat2[i-m1])
			}
		}
		return new(matrix.Matrix).SetBits(mat)
	}
	for _, tCase := range testCases {
		bits1 := getRandMat(tCase.nRows, tCase.nColumns)
		bits2 := getRandMat(tCase.nRows+tCase.nColumns, tCase.nColumns)
		a1 := new(matrix.Matrix).SetBits(bits1)
		a2 := new(matrix.Matrix).SetBits(bits2)
		msg := fmt.Sprintf("ConRows(%s, %s)=\n", a1, a2)
		a1.ConRows(a1, a2)
		b := evalConRows(bits1, bits2)
		if !a1.IsEqual(b) {
			msg += fmt.Sprintf("%s,\nwant\n%s", a1, b)
			t.Errorf(msg)
		}
	}
}

func TestMatrix_ConCols(t *testing.T) {
   testCases := []struct{
       nRows uint
       nColumns uint
   }{
       {5, 5}, {7, 4}, {5, 7}, {1,2},
       {5, 10}, {2,1}, {10, 5}, {100, 100},
       {512, 512}, {237, 237}, {1027, 1027},
       {100, 200}, {512, 1024}, {237, 474},
       {531, 1895}, {200, 100}, {1024, 512},
       {474, 237}, {1895, 531}, {1, 1000},
       {1000, 1},
   }
   evalConCols := func(mat1, mat2 [][]byte) *matrix.Matrix {
       m := uint(len(mat1))
       if m == 0 {
           return new(matrix.Matrix).SetBits(mat2)
       }
       n1 := uint(len(mat1[0]))
       if n1 == 0 {
           return new(matrix.Matrix).SetBits(mat2)
       }
       n2 := uint(len(mat2[0]))
       if n2 == 0 {
           return new(matrix.Matrix).SetBits(mat1)
       }
       mat := make([][]byte, m)
       for i:=uint(0); i < m; i++ {
           mat[i] = make([]byte, n1 + n2)
       }
       for j:=uint(0); j < n1+n2; j++ {
           for i:=uint(0); i < m; i++ {
               if j < n1 {
                   mat[i][j] = mat1[i][j]
               } else {
                   mat[i][j] = mat2[i][j-n1]
               }
           }
       }
       return new(matrix.Matrix).SetBits(mat)
   }
   for _, tCase := range testCases{
       bits1 := getRandMat(tCase.nRows, tCase.nColumns)
       bits2 := getRandMat(tCase.nRows, tCase.nColumns+tCase.nRows)
       a1 := new(matrix.Matrix).SetBits(bits1)
       a2 := new(matrix.Matrix).SetBits(bits2)
       msg := fmt.Sprintf("ConCols(%s, %s)=\n", a1, a2)
       a1.ConCols(a1, a2)
       b := evalConCols(bits1, bits2)
       if !a1.IsEqual(b) {
           msg += fmt.Sprintf("%s,\nwant\n%s", a1, b)
           t.Errorf(msg)
       }
   }
}

func TestMatrix_Add(t *testing.T) {
   testCases := []struct{
       nRows uint
       nColumns uint
   }{
       {5, 5}, {7, 4}, {5, 7}, {1,2},
       {5, 10}, {2,1}, {10, 5},
       {100, 100}, {512, 512}, {237, 237},
       {1027, 1027}, {100, 200}, {512, 1024},
       {237, 474}, {531, 1895}, {200, 100},
       {1024, 512}, {474, 237}, {1895, 531},
       {1, 1000}, {1000, 1},
   }
   evalAdd := func(mat1, mat2 [][]byte) *matrix.Matrix {
       m := uint(len(mat1))
       if m==0 {
           return new(matrix.Matrix).SetZero(0,0)
       }
       n := uint(len(mat1[0]))
       mat := make([][]byte, m)
       for i:=uint(0); i < m; i++ {
           mat[i] = make([]byte, n)
           for j:=uint(0); j < n; j++ {
               mat[i][j] = mat1[i][j] ^ mat2[i][j]
           }
       }
       return new(matrix.Matrix).SetBits(mat)
   }
   for _, tCase := range testCases{
       bits1 := getRandMat(tCase.nRows, tCase.nColumns)
       bits2 := getRandMat(tCase.nRows, tCase.nColumns)
       a1 := new(matrix.Matrix).SetBits(bits1)
       a2 := new(matrix.Matrix).SetBits(bits2)
       msg := fmt.Sprintf("Add(%s, %s)=\n", a1, a2)
       a1.Add(a1, a2)
       b := evalAdd(bits1, bits2)
       if !a1.IsEqual(b) {
           msg += fmt.Sprintf("%s,\nwant\n%s", a1, b)
           t.Errorf(msg)
       }
   }
}

func TestMatrix_Dot(t *testing.T) {
    testCases := []struct{
        nRows uint
        nColumns uint
    }{
        {5, 5}, {7, 4}, {5, 7}, {1,2},
        {5, 10}, {2,1}, {10, 5},
        {100, 100}, {512, 512}, {237, 237},
        {1027, 1027}, {100, 200}, {512, 1024},
        {237, 474}, {531, 1895}, {200, 100},
        {1024, 512}, {474, 237}, {1895, 531},
        {1, 1000}, {1000, 1},
    }
   evalDot := func(mat1, mat2 [][]byte) *matrix.Matrix {
       m1 := uint(len(mat1))
       if m1 == 0 {
           return new(matrix.Matrix).SetZero(0,0)
       }
       n1 := uint(len(mat1[0]))
       if n1 == 0 {
           return new(matrix.Matrix).SetZero(0,0)
       }
       n2 := uint(len(mat2[0]))
       mat := make([][]byte, m1)
       for i:=uint(0); i < m1; i++ {
           mat[i] = make([]byte, n2)
           for j:=uint(0); j < n2; j++ {
               mat[i][j] = 0
               for p:=uint(0); p < n1; p++ {
                   mat[i][j] ^= mat1[i][p] * mat2[p][j]
               }
           }
       }
       return new(matrix.Matrix).SetBits(mat)
   }
   for _, tCase := range testCases{
       bits1 := getRandMat(tCase.nRows, tCase.nColumns)
       bits2 := getRandMat(tCase.nColumns, tCase.nRows + tCase.nColumns)
       a1 := new(matrix.Matrix).SetBits(bits1)
       a2 := new(matrix.Matrix).SetBits(bits2)
       msg := fmt.Sprintf("Dot(%s, %s)=\n", a1, a2)
       a1.Dot(a1, a2)
       b := evalDot(bits1, bits2)
       if !a1.IsEqual(b) {
           msg += fmt.Sprintf("%s,\nwant\n%s", a1, b)
           t.Errorf(msg)
       }
   }
}

func TestMatrix_GetCol(t *testing.T) {
   mat := new(matrix.Matrix).SetBits([][]byte{
       {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
       {0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
       {0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
       {0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
       {0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
   })
   tCases := []struct{
       mat *matrix.Matrix
       num uint
       want *vector.Vector
   }{
       {new(matrix.Matrix).SetZero(0,0), 10, new(vector.Vector).SetZero(0)},
       {mat,0, new(vector.Vector).SetBitArray([]byte{1,0,0,0,0})},
       {mat,1, new(vector.Vector).SetBitArray([]byte{1,0,0,0,1})},
       {mat,11, new(vector.Vector).SetBitArray([]byte{1,1,0,1,1})},
       {mat,15, new(vector.Vector).SetBitArray([]byte{1,1,1,1,1})},
       {mat,29, new(vector.Vector).SetBitArray([]byte{1,1,1,0,1})},
   }
   for _, tCase := range tCases {
       v := tCase.mat.GetCol(tCase.num)
       if v.Cmp(tCase.want) != 0 {
           msg := fmt.Sprintf("%s.GetCol(%d)=%s, want %s",
               tCase.mat, tCase.num, v, tCase.want)
           t.Errorf(msg)
       }
   }
}

func TestMatrix_GetRow(t *testing.T) {
   mat := new(matrix.Matrix).SetBits([][]byte{
       {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
       {0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
       {0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
       {0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
       {0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
   })
   tCases := []struct{
       mat *matrix.Matrix
       num uint
       want *vector.Vector
   }{
       {new(matrix.Matrix).SetZero(0,0), 5, new(vector.Vector).SetZero(0)},
       {mat,0, new(vector.Vector).SetBitArray([]byte{
           1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})},
       {mat,1, new(vector.Vector).SetBitArray([]byte{
           0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1})},
       {mat,2, new(vector.Vector).SetBitArray([]byte{
           0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1})},
       {mat,3, new(vector.Vector).SetBitArray([]byte{
           0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1})},
       {mat,4, new(vector.Vector).SetBitArray([]byte{
           0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1})},
       {mat,24, new(vector.Vector).SetBitArray([]byte{
           0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1})},
   }
   for _, tCase := range tCases {
       v := tCase.mat.GetRow(tCase.num)
       if v.Cmp(tCase.want) != 0 {
           msg := fmt.Sprintf("%s.GetRow(%d)=%s, want %s",
               tCase.mat, tCase.num, v, tCase.want)
           t.Errorf(msg)
       }
   }
}

func TestMatrix_Solve(t *testing.T) {
   tCases := []struct{
       mat *matrix.Matrix
       vec *vector.Vector
       nSolutions uint // 0, 1, >2
   }{
       {new(matrix.Matrix).SetBits([][]byte{
           {1, 1, 1, 1},
           {0, 1, 1, 1},
           {0, 0, 1, 1},
           {0, 0, 0, 1},
       }), new(vector.Vector).SetBitArray([]byte{1,0,1,0}),  1},
       {new(matrix.Matrix).SetBits([][]byte{
           {1, 1, 1, 1},
           {0, 1, 1, 1},
           {0, 0, 1, 1},
           {0, 0, 0, 1},
       }), new(vector.Vector).SetBitArray([]byte{0,0,0,0}),  1},
       {new(matrix.Matrix).SetBits([][]byte{
           {0, 1, 1, 1, 0},
           {0, 0, 1, 0, 1},
           {1, 1, 0, 0, 1},
           {1, 1, 1, 0, 0},
       }), new(vector.Vector).SetBitArray([]byte{1,0,1,0}), 0},
       {new(matrix.Matrix).SetBits([][]byte{
           {0, 1, 1, 1, 0},
           {0, 0, 1, 0, 1},
           {1, 1, 0, 0, 1},
           {1, 1, 1, 0, 0},
       }), new(vector.Vector).SetBitArray([]byte{1,1,1,0}), 2},
   }
   for _, tCase := range tCases {
       fund, sol := tCase.mat.Solve(tCase.vec)
       msg := fmt.Sprintf("%s.Solve(%s)=\nsolution:\n%s\nfundamental system:\n%v\n",
           tCase.mat, tCase.vec, sol, fund)
       if want:=tCase.mat.NColumns() - tCase.mat.Rank(); uint(len(fund)) != want {
           msg += fmt.Sprintf("size of fundamental system eqauls %d, " +
               "want %d (mat.Length - mat.Rank(), %d - %d)",
               len(fund), want, tCase.mat.NColumns(), tCase.mat.Rank())
           t.Errorf(msg)
       }
       for _, v := range fund {
           vMat := new(matrix.Matrix).SetV([]*vector.Vector{v})
           vMat.T(vMat)
           get := new(matrix.Matrix).Dot(tCase.mat, vMat)
           get.T(get)
           if want := new(matrix.Matrix).SetZero(1, tCase.mat.NRows()); !want.IsEqual(get) {
               msg += fmt.Sprintf("%s * (%s)^T = (%s)^T, want (%s)^T",
                   tCase.mat, v, get, want)
               t.Errorf(msg)
           }
       }
       switch tCase.nSolutions {
       case 0:
           if want := new(vector.Vector).SetZero(0); sol.Cmp(want) != 0 {
               msg += fmt.Sprintf("solution = %s, want %s",
                   sol, want)
               t.Errorf(msg)
           }
       default:
           vMat := new(matrix.Matrix).SetV([]*vector.Vector{sol})
           vMat.T(vMat)
           get := new(matrix.Matrix).Dot(tCase.mat, vMat)
           get.T(get)
           if  want := new(matrix.Matrix).SetV([]*vector.Vector{tCase.vec}); !want.IsEqual(get) {
               msg += fmt.Sprintf("%s * (%s)^T = (%s)^T, want (%s)^T",
                   tCase.mat, sol, get, want)
               t.Errorf(msg)
           }
       }
   }
}

func TestMatrix_NonSing(t *testing.T) {
   tests := []struct{
       dim uint
       sampleSize uint
   }{
       {10, 100}, {20, 100}, {30, 100},
       {100, 100}, {200, 100}, {375, 100},
       {1024, 10},
   }
   for _, test := range tests{
       samples := make([]*matrix.Matrix, 0, test.sampleSize)
       for i:=uint(0); i < test.sampleSize; i++ {
           a := new(matrix.Matrix).NonSing(test.dim)
           samples = append(samples, a)
       }
       // Rank test
       for _, s := range samples {
           if r := s.Rank(); r != s.NRows() {
               msg := fmt.Sprintf("matrix %s has rank %d, want %d", s, r, s.NRows())
               t.Errorf(msg)
           }
       }
       // Unique test
       for i:=0; i < len(samples); i++ {
           for j:=i+1; j < len(samples); j++ {
               if samples[i].IsEqual(samples[j]) {
                   msg := fmt.Sprintf("sample[%d] is equal to sample[%d] is %s, "+
                       "want the every matrix in the sample be unique",
                       i, j, samples[i])
                   t.Errorf(msg)
               }
           }
       }
   }
}

func TestMatrix_RandomMaxRank(t *testing.T) {
   tests := []struct{
       nRows uint
       nCols uint
       sampleSize uint
   }{
       {10, 20,100}, {20, 30,100},
       {30, 10,100}, {100, 50,100},
       {200, 301,100}, {375, 271,10},
       {401, 402, 10}, {1024, 2048,5},
   }
   for _, test := range tests{
       samples := make([]*matrix.Matrix, 0, test.sampleSize)
       for i:=uint(0); i < test.sampleSize; i++ {
           a := new(matrix.Matrix).RandomMaxRank(test.nRows, test.nCols)
           samples = append(samples, a)
       }
       // Rank test
       for _, s := range samples {
           rank := s.NRows()
           if r:=s.NColumns(); r < rank {
               rank = r
           }
           if r := s.Rank(); r != rank {
               msg := fmt.Sprintf("matrix %s has rank %d, want %d", s, r, rank)
               t.Errorf(msg)
           }
       }
       // Unique test
       for i:=0; i < len(samples); i++ {
           for j:=i+1; j < len(samples); j++ {
               if samples[i].IsEqual(samples[j]) {
                   msg := fmt.Sprintf("sample[%d] is equal to sample[%d] is %s, "+
                       "want the every matrix in the sample be unique",
                       i, j, samples[i])
                   t.Errorf(msg)
               }
           }
       }
   }
}
