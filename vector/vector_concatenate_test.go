package vector

import (
    "testing"
)

//TestNilConcatenateNil tests function Concatenate of two nil vectors
func TestNilConcatenateNil(t *testing.T) {
    var w Vector
    v, _ := New(nil)
    v.Concatenate(&w)
    if !v.Equal(&w) {
        t.Errorf("vector testing: nil.Concatenate(nil) is incorrect, nil.Concatenate(nil) != nil, but %v",
            v)
    }
}

//TestNilConcatenateEmpty tests function Concatenate of nil vector and empty vector
func TestNilConcatenateEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    w, _ := New(nil)
    v.Concatenate(w)
    if !v.Equal(w) {
        t.Errorf("vector testing: empty.Concatenate(nil) is incorrect, empty.Concatenate(nil) != nil, but %v",
            v)
    }
}

//TestEmptyConcatenateLen1 tests function Concatenate of empty vector and vector of length 1
func TestEmptyConcatenateLen1(t *testing.T) {
    v, _ := New(nil)
    v0, _ := New([]uint8{0})
    v1, _ := New([]uint8{1})
    res, _ := New([]uint8{0, 1})
    v.Concatenate(v0)
    v.Concatenate(v1)
    if !v.Equal(res) {
        t.Errorf("vector testing: empty.Concatenate([0], [1]) is incorrect, is %v, but expected 01",
            v)
    }
}

//TestConcatenateLess64 tests function Concatenate for vectors of length less than 64
func TestConcatenateLess64(t *testing.T) {
    v, _ := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0,
    })
    w, _ := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
    })
    res, _ := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 0,
    })
    v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) < 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenate64 tests function Concatenate for vectors of length 64
func TestConcatenate64(t *testing.T) {
    v, _ := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0,
    })
    w, _ := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
    })
    res, _ := New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0,
    })
    v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) == 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}

//TestConcatenateMore64 tests function Concatenate for vectors of length more than 64
func TestConcatenateMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
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
    v.Concatenate(w)
    if !v.Equal(res) {
        t.Errorf("vector testing: Len(Concatenate) > 64 is incorrect, is %v, but expected %v",
            v, res)
    }
}
