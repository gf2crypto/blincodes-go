package vector

import (
    "testing"
)

//TestCopyNil tests function Copy for nil vector
func TestCopyNil(t *testing.T) {
    v, _ := New(nil)
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}

//TestCopyEmpty tests function Copy for empty vector
func TestCopyEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}

//TestCopyLen1 tests function Copy for vector of length 1
func TestCopyLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}

//TestCopyLenLess64 tests function Copy for vector of length less than 64
func TestCopyLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}

//TestCopyLen64 tests function Copy for vector of length 64
func TestCopyLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}

//TestCopyLenMore64 tests function Copy for vector of length more than 64
func TestCopyLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := v.Copy()
    if !w.Equal(v) {
        t.Errorf("vector testing: Copy() is incorrect, %v != %v",
            w, v)
    }
}
