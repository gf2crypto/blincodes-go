package vector

import (
    "testing"
)

//TestNilAndNil tests function Or for two nil vectors
func TestNilAndNil(t *testing.T) {
    var w Vector
    v := New(nil)
    res := v.And(&w)
    if !res.Equal(&w) {
        t.Errorf("vector testing: nil AND nil is incorrect, nil AND nil != nil, but %v",
            res)
    }
}

//TestNilAndEmpty tests function And for nil vector and empty vector
func TestNilAndEmpty(t *testing.T) {
    w := New([]uint8{})
    v := New(nil)
    res := v.And(w)
    if !res.Equal(w) {
        t.Errorf("vector testing: nil And empty is incorrect, nil And empty != nil, but %v",
            res)
    }
}

//TestEmptyAndNil tests function And for empty vector and nil vector
func TestEmptyAndNil(t *testing.T) {
    v := New([]uint8{})
    w := New(nil)
    res := v.And(w)
    if !res.Equal(w) {
        t.Errorf("vector testing: empty And nil is incorrect, empty And nil != nil, but %v",
            res)
    }
}

//TestAndLess64 tests function And for vectors of length less than 64
func TestAndLess64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v := u.And(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: And is incorrect, %v And %v = %v != %v",
            u, w, v, res)
    }
}

//TestAnd64 tests function And for vectors of length 64
func TestAnd64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    w := New([]uint8{
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    v := u.And(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: And is incorrect, %v And %v = %v != %v",
            u, w, v, res)
    }
}

//TestAndMore64 tests function And for vectors of length more than 64
func TestAndMore64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
        1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1,
        1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := New([]uint8{
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1,
        0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1,
        0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v := u.And(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: And is incorrect, %v And %v = %v != %v",
            u, w, v, res)
    }
}
