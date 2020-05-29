package vector

import (
    "testing"
)

//TestNotNil tests function NOT for nil vector
func TestNotNil(t *testing.T) {
    v := New(nil)
    res := New(nil)
    v = v.Not()
    if !v.Equal(res) {
        t.Errorf("vector testing: NOT nil is incorrect, NOT nil != nil, but %v",
            v)
    }
}

//TestNotEmpty tests function Not of empty vector
func TestNotEmpty(t *testing.T) {
    v := New([]uint8{})
    res := New(nil)
    v = v.Not()
    if !v.Equal(res) {
        t.Errorf("vector testing: Not empty is incorrect, Not empty != nil, but %v",
            v)
    }
}

//TestNotLess64 tests function Not for vectors of length less than 64
func TestNotLess64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    v := u.Not()
    if !v.Equal(res) {
        t.Errorf("vector testing: Not is incorrect, Not %v = %v != %v",
            u, v, res)
    }
}

//TestNot64 tests function Not for vectors of length 64
func TestNot64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    res := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0,
        1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
    })
    v := u.Not()
    if !v.Equal(res) {
        t.Errorf("vector testing: Not is incorrect, Not %v = %v != %v",
            u, v, res)
    }
}

//TestNotMore64 tests function Not for vectors of length more than 64
func TestNotMore64(t *testing.T) {
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0,
        1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
        1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    v := u.Not()
    if !v.Equal(res) {
        t.Errorf("vector testing: Not is incorrect, Not %v = %v != %v",
            u, v, res)
    }
}
