package vector

import (
    "testing"
)

//TestNilOrNil tests function Or for two nil vectors
func TestNilOrNil(t *testing.T) {
    var w Vector
    v, _ := New(nil)
    v = v.Or(&w)
    if !v.Equal(&w) {
        t.Errorf("vector testing: nil Or nil is incorrect, nil Or nil != nil, but %v",
            v)
    }
}

//TestNilOrEmpty tests function Or for nil vector and empty vector
func TestNilOrEmpty(t *testing.T) {
    w, _ := New([]uint8{})
    v, _ := New(nil)
    v = v.Or(w)
    if !v.Equal(w) {
        t.Errorf("vector testing: nil Or empty is incorrect, nil Or empty != nil, but %v",
            v)
    }
}

//TestEmptyOrNil tests function Or for empty vector and nil vector
func TestEmptyOrNil(t *testing.T) {
    v, _ := New([]uint8{})
    w, _ := New(nil)
    v = v.Or(w)
    if !v.Equal(w) {
        t.Errorf("vector testing: empty Or nil is incorrect, empty Or nil != nil, but %v",
            v)
    }
}

//TestOrLess64 tests function Or for vectors of length less than 64
func TestOrLess64(t *testing.T) {
    u, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w, _ := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v := u.Or(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Or is incorrect, %v Or %v = %v != %v",
            u, w, v, res)
    }
}

//TestOr64 tests function Or for vectors of length 64
func TestOr64(t *testing.T) {
    u, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    w, _ := New([]uint8{
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v := u.Or(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Or is incorrect, %v Or %v = %v != %v",
            u, w, v, res)
    }
}

//TestOrMore64 tests function Or for vectors of length more than 64
func TestOrMore64(t *testing.T) {
    u, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
        1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1,
        1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0,
        0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
        1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w, _ := New([]uint8{
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1,
        1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1,
        0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v := u.Or(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Or is incorrect, %v Or %v = %v != %v",
            u, w, v, res)
    }
}
