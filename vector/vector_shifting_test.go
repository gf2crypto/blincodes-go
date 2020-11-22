package vector

import (
    "testing"
)

//TestNilShift tests functions ShiftLeft and ShiftRight for nil vectors
func TestNilShift(t *testing.T) {
    v := New(nil)
    resL := v.ShiftLeft(10)
    if !resL.Equal(v) {
        t.Errorf("vector testing: nil.ShiftLeft(10) is incorrect, nil.ShiftLeft(10) != nil, but %v",
            v)
    }
    resR := v.ShiftRight(10)
    if !resR.Equal(v) {
        t.Errorf("vector testing: nil.ShiftRight(10) is incorrect, nil.ShiftRight(10) != nil, but %v",
            v)
    }
}

//TestShiftLen1 tests functions ShiftLeft and ShiftRight for vector of length 1
func TestShiftLen1(t *testing.T) {
    v0 := New([]uint8{0})
    v1 := New([]uint8{1})
    res := New([]uint8{0})
    v := v0.ShiftLeft(0)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].ShiftLeft(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftLeft(0)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [1].ShiftLeft(0) is incorrect, is %v, but expected [1]",
            v)
    }
    v = v0.ShiftLeft(1)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftLeft(1) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftLeft(1)
    if !v.Equal(res) {
        t.Errorf("vector testing: [1].ShiftLeft(1) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v0.ShiftLeft(3)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftLeft(3) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftLeft(3)
    if !v.Equal(res) {
        t.Errorf("vector testing: [1].ShiftLeft(3) is incorrect, is %v, but expected [0]",
            v)
    }

    v = v0.ShiftRight(0)
    if !v.Equal(v0) {
        t.Errorf("vector testing: [0].ShiftRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftRight(0)
    if !v.Equal(v1) {
        t.Errorf("vector testing: [0].ShiftRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v0.ShiftRight(1)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftRight(1)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v0.ShiftRight(3)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftRight(0) is incorrect, is %v, but expected [0]",
            v)
    }
    v = v1.ShiftLeft(3)
    if !v.Equal(res) {
        t.Errorf("vector testing: [0].ShiftLeft(0) is incorrect, is %v, but expected [0]",
            v)
    }
}

//TestShiftLess64 tests functions ShiftLeft and ShiftRight for vectors of length less than 64
func TestShiftLess64(t *testing.T) {
    v := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0,
    })
    resLeft1 := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 0,
    })
    resRight1 := New([]uint8{
        0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1,
    })
    resLeft21 := New([]uint8{
        1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0,
    })
    resRight21 := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1,
    })
    resZero := New(40)
    res := v.ShiftLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.ShiftLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.ShiftLeft(21)
    if !res.Equal(resLeft21) {
        t.Errorf("vector testing: %v.ShiftLeft(21) is incorrect, is %v, but expected %v",
            v, res, resLeft21)
    }
    res = v.ShiftLeft(40)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftLeft(40) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }

    res = v.ShiftRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.ShiftRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.ShiftRight(21)
    if !res.Equal(resRight21) {
        t.Errorf("vector testing: %v.ShiftRight(21) is incorrect, is %v, but expected %v",
            v, res, resRight21)
    }
    res = v.ShiftRight(40)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftRight(40) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }
}

//TestShift64 tests functions ShiftLeft and ShiftRight for vectors of length 64
func TestShift64(t *testing.T) {
    v := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
    })
    resLeft1 := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0,
    })
    resRight1 := New([]uint8{
        0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1,
    })
    resLeft28 := New([]uint8{
        1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1,
        0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1,
        0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    resRight28 := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1,
        0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1,
        1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1,
    })
    resZero := New(v.Len())
    res := v.ShiftLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.ShiftLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.ShiftLeft(28)
    if !res.Equal(resLeft28) {
        t.Errorf("vector testing: %v.ShiftLeft(28) is incorrect, is %v, but expected %v",
            v, res, resLeft28)
    }
    res = v.ShiftLeft(64)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftLeft(64) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }

    res = v.ShiftRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.ShiftRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.ShiftRight(28)
    if !res.Equal(resRight28) {
        t.Errorf("vector testing: %v.ShiftRight(28) is incorrect, is %v, but expected %v",
            v, res, resRight28)
    }
    res = v.ShiftRight(64)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftRight(64) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }
}

//TestShiftMore64 tests functions ShiftLeft and ShiftRight for vectors of length more than 64
func TestShiftMore64(t *testing.T) {
    v := New([]uint8{
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1,
    })
    resLeft1 := New([]uint8{
        0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1, 1, 0,
    })
    resRight1 := New([]uint8{
        0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
        1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0,
        0, 1, 1, 0, 1, 0, 0, 1,
    }) //00000000000000000000000000000000000000
    //0000000000000000000000000000
    resLeft28 := New([]uint8{
        1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1,
        0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1,
        0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    resRight28 := New([]uint8{
        0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1,
        1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0,
    })
    resLeft67 := New([]uint8{
        0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0,
    })
    resRight67 := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 0, 1, 0, 1,
    })
    resZero := New(v.Len())
    res := v.ShiftLeft(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftLeft(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftLeft(1)
    if !res.Equal(resLeft1) {
        t.Errorf("vector testing: %v.ShiftLeft(1) is incorrect, is %v, but expected %v",
            v, res, resLeft1)
    }
    res = v.ShiftLeft(28)
    if !res.Equal(resLeft28) {
        t.Errorf("vector testing: %v.ShiftLeft(28) is incorrect, is %v, but expected %v",
            v, res, resLeft28)
    }
    res = v.ShiftLeft(67)
    if !res.Equal(resLeft67) {
        t.Errorf("vector testing: %v.ShiftLeft(67) is incorrect, is %v, but expected %v",
            v, res, resLeft67)
    }
    res = v.ShiftLeft(80)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftLeft(80) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }

    res = v.ShiftRight(0)
    if !res.Equal(v) {
        t.Errorf("vector testing: %v.ShiftRight(0) is incorrect, is %v, but expected %v",
            v, res, v)
    }
    res = v.ShiftRight(1)
    if !res.Equal(resRight1) {
        t.Errorf("vector testing: %v.ShiftRight(1) is incorrect, is %v, but expected %v",
            v, res, resRight1)
    }
    res = v.ShiftRight(28)
    if !res.Equal(resRight28) {
        t.Errorf("vector testing: %v.ShiftRight(28) is incorrect, is %v, but expected %v",
            v, res, resRight28)
    }
    res = v.ShiftRight(67)
    if !res.Equal(resRight67) {
        t.Errorf("vector testing: %v.ShiftRight(67) is incorrect, is %v, but expected %v",
            v, res, resRight67)
    }
    res = v.ShiftRight(80)
    if !res.Equal(resZero) {
        t.Errorf("vector testing: %v.ShiftRight(80) is incorrect, is %v, but expected %v",
            v, res, resZero)
    }
}
