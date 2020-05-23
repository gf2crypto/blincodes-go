package vector

import (
    "testing"
)

//TestNilRotate tests functions RotateLeft and RotateRight for nil vectors
func TestNilRotate(t *testing.T) {
    v, _ := New(nil)
    resL := v.RotateLeft(10)
    if !resL.Equal(v) {
        t.Errorf("vector testing: nil.RotateLeft(10) is incorrect, nil.RotateLeft(10) != nil, but %v",
            v)
    }
    resR := v.RotateRight(10)
    if !resR.Equal(v) {
        t.Errorf("vector testing: nil.RotateRight(10) is incorrect, nil.RotateRight(10) != nil, but %v",
            v)
    }
}

//TestRotateLen1 tests functions RotateLeft and RotateRight for vector of length 1
func TestRotateLen1(t *testing.T) {
    v0, _ := New([]uint8{0})
    v1, _ := New([]uint8{1})
    v := v0.RotateLeft(0)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateLeft(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateLeft(0)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateLeft(0) is incorrect, is %v, but expected [1]",
            v)
    }
    v = v0.RotateLeft(1)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateLeft(1) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateLeft(1)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateLeft(1) is incorrect, is %v, but expected [1]",
            v)
    }
    v = v0.RotateLeft(3)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateLeft(3) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateLeft(3)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateLeft(3) is incorrect, is %v, but expected [1]",
            v)
    }

    v = v0.RotateRight(0)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateRight(0)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateRight(0) is incorrect, is %v, but expected [1]",
            v)
    }
    v = v0.RotateRight(1)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateRight(1)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateRight(1) is incorrect, is %v, but expected [1]",
            v)
    }
    v = v0.RotateRight(3)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].RotateRight(3) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.RotateLeft(3)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].RotateLeft(3) is incorrect, is %v, but expected [1]",
            v)
    }
}

//TestRotateLess64 tests functions RotateLeft and RotateRight for vectors of length less than 64
func TestRotateLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0,
    })
    resLeft1, _ := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1,
    })
    resRight1, _ := New([]uint8{
        0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1,
    })
    resLeft21, _ := New([]uint8{
        1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1,
    })
    resRight21, _ := New([]uint8{
        1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1,
    })
    res := v.RotateLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.RotateLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.RotateLeft(21)
    if !res.Equal(resLeft21) {
        t.Errorf("vector testing: %v.RotateLeft(21) is incorrect, is %v, but expected %v",
            v, res, resLeft21)
    }
    res = v.RotateLeft(40)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateLeft(40) is incorrect, is %v, but expected %v",
            v, res, v)
    }

    res = v.RotateRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.RotateRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.RotateRight(21)
    if !res.Equal(resRight21) {
        t.Errorf("vector testing: %v.RotateRight(21) is incorrect, is %v, but expected %v",
            v, res, resRight21)
    }
    res = v.RotateRight(40)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateRight(40) is incorrect, is %v, but expected %v",
            v, res, v)
    }
}

//TestRotate64 tests functions RotateLeft and RotateRight for vectors of length 64
func TestRotate64(t *testing.T) {
    v, _ := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
    })
    resLeft1, _ := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1,
    })
    resRight1, _ := New([]uint8{
        0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1,
    })
    resLeft28, _ := New([]uint8{
        1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0,
    })
    resRight28, _ := New([]uint8{
        1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1,
    })
    res := v.RotateLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.RotateLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.RotateLeft(28)
    if !res.Equal(resLeft28) {
        t.Errorf("vector testing: %v.RotateLeft(28) is incorrect, is %v, but expected %v",
            v, res, resLeft28)
    }
    res = v.RotateLeft(64)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateLeft(64) is incorrect, is %v, but expected %v",
            v, res, v)
    }

    res = v.RotateRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.RotateRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.RotateRight(28)
    if !res.Equal(resRight28) {
        t.Errorf("vector testing: %v.RotateRight(28) is incorrect, is %v, but expected %v",
            v, res, resRight28)
    }
    res = v.RotateRight(64)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateRight(64) is incorrect, is %v, but expected %v",
            v, res, v)
    }
}

//TestRotateMore64 tests functions RotateLeft and RotateRight for vectors of length more than 64
func TestRotateMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1,
    })
    resLeft1, _ := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1, 1,
    })
    resRight1, _ := New([]uint8{
        1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1,
    })
    resLeft28, _ := New([]uint8{
        1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0,
    })
    resRight28, _ := New([]uint8{
        1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0,
    })
    resLeft67, _ := New([]uint8{
        0, 1, 0, 0, 1, 1,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1,
    })
    resRight67, _ := New([]uint8{
        1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1,
        1, 0, 0, 1, 0, 1,
    })
    resLeft80, _ := New([]uint8{
        1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1,
    })
    resRight80, _ := New([]uint8{
        1, 0, 1, 0, 0, 1, 1,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1,
    })
    res := v.RotateLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.RotateLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.RotateLeft(28)
    if !res.Equal(resLeft28) {
        t.Errorf("vector testing: %v.RotateLeft(28) is incorrect, is %v, but expected %v",
            v, res, resLeft28)
    }
    res = v.RotateLeft(67)
    if !res.Equal(resLeft67) {
        t.Errorf("vector testing: %v.RotateLeft(67) is incorrect, is %v, but expected %v",
            v, res, resLeft67)
    }
    res = v.RotateLeft(80)
    if !res.Equal(resLeft80) {
        t.Errorf("vector testing: %v.RotateLeft(80) is incorrect, is %v, but expected %v",
            v, res, resLeft80)
    }

    res = v.RotateRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.RotateRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.RotateRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.RotateRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.RotateRight(28)
    if !res.Equal(resRight28) {
        t.Errorf("vector testing: %v.RotateRight(28) is incorrect, is %v, but expected %v",
            v, res, resRight28)
    }
    res = v.RotateRight(67)
    if !res.Equal(resRight67) {
        t.Errorf("vector testing: %v.RotateRight(67) is incorrect, is %v, but expected %v",
            v, res, resRight67)
    }
    res = v.RotateRight(80)
    if !res.Equal(resRight80) {
        t.Errorf("vector testing: %v.RotateRight(80) is incorrect, is %v, but expected %v",
            v, res, resRight80)
    }
}
