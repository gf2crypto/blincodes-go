package vector

import (
    "testing"
)

//TestCopyNil tests function Copy for nil vector
func TestCopyNil(t *testing.T) {
    v := New(nil)
    w := v.Copy()
    if len(w.body) != 0 {
        t.Errorf("vector testing: Copy() is incorrect, len(Copy().body) != 0 (%d != 0)",
            len(w.body))
    }
    if w.lenLast != 0 {
        t.Errorf("vector testing: Copy() is incorrect, Copy().lenLast != 0 (%d != 0)",
            w.lenLast)
    }
}

//TestCopyEmpty tests function Copy for empty vector
func TestCopyEmpty(t *testing.T) {
    v := New([]uint8{})
    w := v.Copy()
    if len(w.body) != 0 {
        t.Errorf("vector testing: Copy() is incorrect, len(Copy().body) != 0 (%d != 0)",
            len(w.body))
    }
    if w.lenLast != 0 {
        t.Errorf("vector testing: Copy() is incorrect, Copy().lenLast != 0 (%d != 0)",
            w.lenLast)
    }
}

//TestCopyLen1 tests function Copy for vector of length 1
func TestCopyLen1(t *testing.T) {
    v := New([]uint8{0})
    w := v.Copy()
    if len(w.body) != len(v.body) {
        t.Errorf("vector testing: Copy() is incorrect, len(v.Copy()) != v (%d != %d)",
            len(w.body), len(v.body))
    }
    for i, b := range w.body {
        if b != v.body[i] {
            t.Errorf("vector testing: Copy() is incorrect, v.Copy().body != v.body (%v != %v)",
                w.body, v.body)
        }
    }
    if w.lenLast != v.lenLast {
        t.Errorf("vector testing: Copy() is incorrect, v.Copy().lenLast != v.lenLast (%v != %v)",
            w.lenLast, v.lenLast)
    }
}

//TestCopyLenLess64 tests function Copy for vector of length less than 64
func TestCopyLenLess64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := v.Copy()
    if len(w.body) != len(v.body) {
        t.Errorf("vector testing: Copy() is incorrect, len(v1.Copy()) != v1 (%d != %d)",
            len(w.body), len(v.body))
    }
    for i, b := range w.body {
        if b != v.body[i] {
            t.Errorf("vector testing: Copy() is incorrect, v1.Copy().body != v1.body (%v != %v)",
                w.body, v.body)
        }
    }
    if w.lenLast != v.lenLast {
        t.Errorf("vector testing: Copy() is incorrect, v1.Copy().lenLast != v1.lenLast (%v != %v)",
            w.lenLast, v.lenLast)
    }
}

//TestCopyLen64 tests function Copy for vector of length 64
func TestCopyLen64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    w := v.Copy()
    if len(w.body) != len(v.body) {
        t.Errorf("vector testing: Copy() is incorrect, len(v2.Copy()) != v2 (%d != %d)",
            len(w.body), len(v.body))
    }
    for i, b := range w.body {
        if b != v.body[i] {
            t.Errorf("vector testing: Copy() is incorrect, v2.Copy().body != v2.body (%v != %v)",
                w.body, v.body)
        }
    }
    if w.lenLast != v.lenLast {
        t.Errorf("vector testing: Copy() is incorrect, v2.Copy().lenLast != v2.lenLast (%v != %v)",
            w.lenLast, v.lenLast)
    }
}

//TestCopyLenMore64 tests function Copy for vector of length more than 64
func TestCopyLenMore64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := v.Copy()
    if len(w.body) != len(v.body) {
        t.Errorf("vector testing: Copy() is incorrect, len(v3.Copy()) != v3 (%d != %d)",
            len(w.body), len(v.body))
    }
    for i, b := range w.body {
        if b != v.body[i] {
            t.Errorf("vector testing: Copy() is incorrect, v3.Copy().body != v3.body (%v != %v)",
                w.body, v.body)
        }
    }
    if w.lenLast != v.lenLast {
        t.Errorf("vector testing: Copy() is incorrect, v3.Copy().lenLast != v3.lenLast (%v != %v)",
            w.lenLast, v.lenLast)
    }
}
