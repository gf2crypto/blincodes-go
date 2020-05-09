package vector

import (
    "testing"
)

//TestNilXorNil tests function XOR for two nil vectors
func TestNilXorNil(t *testing.T) {
    var w Vector
    v := New(nil)
    v.Xor(&w)
    if v.Len() != 0 {
        t.Errorf("vector testing: nil XOR nil is incorrect, nil XOR nil != nil, but %v",
            v)
    }
}

//TestNilXorEmpty tests function XOR for nil vector and empty vector
func TestNilXorEmpty(t *testing.T) {
    w := New([]uint8{})
    v := New(nil)
    v.Xor(w)
    if v.Len() != 0 {
        t.Errorf("vector testing: nil XOR empty is incorrect, nil XOR empty != nil, but %v",
            v)
    }
}

//TestEmptyXorNil tests function XOR for empty vector and nil vector
func TestEmptyXorNil(t *testing.T) {
    v := New([]uint8{})
    w := New(nil)
    v.Xor(w)
    if v.Len() != 0 {
        t.Errorf("vector testing: empty XOR nil is incorrect, empty XOR nil != nil, but %v",
            v)
    }
}

//TestXorLess64 tests function Xor for vectors of length less than 64
func TestXorLess64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    u := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    w := New([]uint8{
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    v.Xor(w)
    if v.Len() != 32 {
        t.Errorf("vector testing: XOR is incorrect, length of result != 32 (%d != 32)",
            v.Len())
    }
    if v.body[0] != res.body[0] {
        t.Errorf("vector testing: XOR is incorrect, %v XOR %v = %v != %v",
            u, w, v, res)
    }
}

//TestXor64 tests function Xor for vectors of length 64
func TestXor64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
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
        1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
        1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0,
    })
    v.Xor(w)
    if v.Len() != 64 {
        t.Errorf("vector testing: XOR is incorrect, length of result != 32 (%d != 32)",
            v.Len())
    }
    if v.body[0] != res.body[0] {
        t.Errorf("vector testing: XOR is incorrect, %v XOR %v = %v != %v",
            u, w, v, res)
    }
}

//TestXorMore64 tests function Xor for vectors of length more than 64
func TestXorMore64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
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
        1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0,
        1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0,
        1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
        0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    })
    v.Xor(w)
    if v.Len() != 93 {
        t.Errorf("vector testing: XOR is incorrect, length of result != 32 (%d != 32)",
            v.Len())
    }
    if v.body[0] != res.body[0] && res.body[0] != res.body[1] {
        t.Errorf("vector testing: XOR is incorrect, %v XOR %v = %v != %v",
            u, w, v, res)
    }
}
