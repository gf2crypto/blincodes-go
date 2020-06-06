package rm

// import "fmt"
import "testing"
import "github.com/gf2crypto/blincodes-go/matrix"

//TestRMGenerator00 test the evaluation of the generator matrix of RM(0, 0) code
func TestRMGenerator00(t *testing.T) {
    rm := New(0, 0)
    if rm.N() != 1 {
        t.Errorf("expected rm.N == 1, but rm.N=%v", rm.N())
    }
    if rm.K() != 1 {
        t.Errorf("expected rm.K == 1, but rm.K=%v", rm.K())
    }
    if g := rm.Gen(); !g.Equal(matrix.New(1, []uint8{1})) {
        t.Errorf("expected rm.G == [1], but rm.G=%v", g)
    }
    if h := rm.ParityCheck(); !h.Equal(matrix.New(1)) {
        t.Errorf("expected rm.H == [0], but rm.H=%v", h)
    }
}

//TestRMGenerator04 test the evaluation of the generator matrix of RM(0, 4) code
func TestRMGenerator04(t *testing.T) {
    rm := New(0, 4)
    if rm.N() != 16 {
        t.Errorf("expected rm.N == 16, but rm.N=%v", rm.N())
    }
    if rm.K() != 1 {
        t.Errorf("expected rm.K == 1, but rm.K=%v", rm.K())
    }
    gen := matrix.New(1, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if g := rm.Gen(); !g.Equal(gen) {
        t.Errorf("expected rm.G == %v, but rm.G=%v", gen, g)
    }
    if h := rm.ParityCheck(); !gen.Mul(h.T()).Equal(matrix.New(1, 15)) {
        t.Errorf("expected rm.G*(rm.H)^T == 0, rm.H==\n%v", h)
    }
}

//TestRMGenerator14 test the evaluation of the generator matrix of RM(1, 4) code
func TestRMGenerator14(t *testing.T) {
    rm := New(1, 4)
    if rm.N() != 16 {
        t.Errorf("expected rm.N == 16, but rm.N=%v", rm.N())
    }
    if rm.K() != 5 {
        t.Errorf("expected rm.K == 5, but rm.K=%v", rm.K())
    }
    gen := matrix.New(5, []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
    })
    if g := rm.Gen(); !g.Equal(gen) {
        t.Errorf("expected rm.G == %v, but rm.G=\n%v", gen, g)
    }
    if h := rm.ParityCheck(); !h.Equal(New(2, 4).Gen()) {
        t.Errorf("expected rm.H == rm.G(2,4), rm.H==\n%v", h)
    }
    if h := rm.ParityCheck(); !rm.Gen().Mul(h.T()).Equal(matrix.New(5, 11)) {
        t.Errorf("expected rm.G*(rm.H)^T == 0, rm.H==\n%v", h)
    }
}

//TestRMGenerator test the evaluation of the generator matrix of RM code
func TestRMGenerator(t *testing.T) {
    if g, h := New(3, 8).Gen(), New(4, 8).Gen(); !g.Mul(h.T()).Equal(matrix.New(g.Nrows(), h.Nrows())) {
        t.Errorf("expected rm.G(3, 8)*(rm.G(4, 8))^T == 0, rm.G(3,8)==\n%v\nrm.G(4,8)==\n%v", g, h)
    }
    if g, h := New(3, 8).ParityCheck(), New(4, 8).ParityCheck(); !g.Mul(h.T()).Equal(matrix.New(g.Nrows(), h.Nrows())) {
        t.Errorf("expected rm.H(3, 8)*(rm.H(4, 8))^T == 0, rm.H(3,8)==\n%v\nrm.H(4,8)==\n%v", g, h)
    }
}

//TestRMGenerator310 test the evaluation of the generator matrix of RM(3, 10) code
func TestRMGenerator310(t *testing.T) {
    if g, h := New(3, 10).Gen(), New(6, 10).Gen(); !g.Mul(h.T()).Equal(matrix.New(g.Nrows(), h.Nrows())) {
        t.Errorf("expected rm.G(3, 10)*(rm.G(6, 10))^T == 0, rm.G(3,10)==\n%v\nrm.G(6,10)==\n%v", g, h)
    }
    if g, h := New(3, 10).ParityCheck(), New(6, 10).ParityCheck(); !g.Mul(h.T()).Equal(matrix.New(g.Nrows(), h.Nrows())) {
        t.Errorf("expected rm.H(3, 10)*(rm.H(6, 10))^T == 0, rm.H(3,10)==\n%v\nrm.H(6,10)==\n%v", g, h)
    }
}
