package vector

import (
    "testing"
)

//TestNilConcatenateNil tests function Concatenate of two nil vectors
func TestNilConcatenateNil(t *testing.T) {
    var w Vector
    v := New(nil)
    res := v.Concatenate(&w)
    if !res.Equal(&w) {
        t.Errorf("vector testing: nil.Concatenate(nil) is incorrect, nil.Concatenate(nil) != nil, but %v",
            res)
    }
}

//TestNilConcatenateEmpty tests function Concatenate of nil vector and empty vector
func TestNilConcatenateEmpty(t *testing.T) {
    v := New([]uint8{})
    w := New(nil)
    res := v.Concatenate(w)
    if !res.Equal(w) {
        t.Errorf("vector testing: empty.Concatenate(nil) is incorrect, empty.Concatenate(nil) != nil, but %v",
            res)
    }
}

//TestEmptyConcatenateLen1 tests function Concatenate of empty vector and vector of length 1
func TestEmptyConcatenateLen1(t *testing.T) {
    v := New(nil)
    v0 := New([]uint8{0})
    v1 := New([]uint8{1})
    res := New([]uint8{0, 1})
    u := v.Concatenate(v0).Concatenate(v1)
    if !u.Equal(res) {
        t.Errorf("vector testing: empty.Concatenate([0], [1]) is incorrect, is %v, but expected 01",
            v)
    }
}

//TestConcatenateLess64 tests function Concatenate for vectors of length less than 64
func TestConcatenateLess64(t *testing.T) {
    v := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0,
    })
    w := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
    })
    res := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 0,
    })
    v = v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) < 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenate64 tests function Concatenate for vectors of length 64
func TestConcatenate64(t *testing.T) {
    v := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0,
    })
    w := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
    })
    res := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
    })
    v = v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) == 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenateMore64 tests function Concatenate for vectors of length more than 64
func TestConcatenateMore64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
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
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v = v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) > 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenateMore64SmallLast1 tests function Concatenate for vectors of length more than 64
func TestConcatenateMore64SmallLast1(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1,
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
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    v = v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) > 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenateMore64SmallLast2 tests function Concatenate for vectors of length more than 64
func TestConcatenateMore64SmallLast2(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1,
    })
    w := New([]uint8{
        0, 1, 1, 0, 1, 1,
    })
    res := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1,
    })
    v = v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) > 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}
