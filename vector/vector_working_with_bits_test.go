package vector

import (
    "fmt"
    "testing"
)
import "bytes"

//TestGetLen1 tests function Get for vector of length 1
func TestGetLen1(t *testing.T) {
    var buf bytes.Buffer
    v := New([]uint8{0})
    for i := 0; i < v.Len(); i++ {
        fmt.Fprintf(&buf, "%v", v.Get(i))
    }
    if buf.String() != "0" {
        t.Errorf("vector testing: Get() is incorrect, %s",
            buf.String())
    }
}

//TestGetLenLess64 tests function Get for vector of length less than 64
func TestGetLenLess64(t *testing.T) {
    var buf bytes.Buffer
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    buf.Reset()
    for i := 0; i < v.Len(); i++ {
        fmt.Fprintf(&buf, "%v", v.Get(i))
    }
    if buf.String() != "11111111111111110000000011111111" {
        t.Errorf("vector testing: v1.Get() is incorrect, %s",
            buf.String())
    }
}

//TestGetLen64 tests function Get for vector of length 64
func TestGetLen64(t *testing.T) {
    var buf bytes.Buffer
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    buf.Reset()
    for i := 0; i < v.Len(); i++ {
        fmt.Fprintf(&buf, "%v", v.Get(i))
    }
    if buf.String() != "1111111111111111000000001111111100001111000011110011001100110011" {
        t.Errorf("vector testing: v2.Get() is incorrect, %s",
            buf.String())
    }
}

//TestGetLenMore64 tests function Get for vector of length more than 64
func TestGetLenMore64(t *testing.T) {
    var buf bytes.Buffer
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    buf.Reset()
    for i := 0; i < v.Len(); i++ {
        fmt.Fprintf(&buf, "%v", v.Get(i))
    }
    if buf.String() != "111111111111111100000000111111110000111100001111001100110011001101010101010101010000111111111" {
        t.Errorf("vector testing: v2.Get() is incorrect, %s",
            buf.String())
    }
}

//TestBitsNil tests function Bits for nil vector
func TestBitsNil(t *testing.T) {
    v := New(nil)
    if len(v.Bits()) != 0 {
        t.Errorf("vector testing: nil.Bits() is incorrect, %v",
            v.Bits())
    }
}

//TestBitsEmpty tests function Bits for empty vector
func TestBitsEmpty(t *testing.T) {
    v := New([]uint8{})
    if len(v.Bits()) != 0 {
        t.Errorf("vector testing: []uint8{}.Bits() is incorrect, %v",
            v.Bits())
    }
}

//TestBitsLen1 tests function Bits for vector of length 1
func TestBitsLen1(t *testing.T) {
    vnil := []uint8{0}
    v := New(vnil)
    for i, b := range v.Bits() {
        if vnil[i] != b {
            t.Errorf("vector testing: []uint8{}.Bits() is incorrect, (vnil[%d] = %d) != %d",
                i, vnil[i], b)
        }
    }
}

//TestBitsLenLess64 tests function Bits for vector of length less than 64
func TestBitsLenLess64(t *testing.T) {
    v1 := []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    }
    v := New(v1)
    for i, b := range v.Bits() {
        if v1[i] != b {
            t.Errorf("vector testing: v1.Bits() is incorrect, (v1[%d] = %d) != %d",
                i, v1[i], b)
        }
    }
}

//TestBitsLen64 tests function Bits for vector of length 64
func TestBitsLen64(t *testing.T) {
    v2 := []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    }
    v := New(v2)
    for i, b := range v.Bits() {
        if v2[i] != b {
            t.Errorf("vector testing: v2.Bits() is incorrect, (v2[%d] = %d) != %d",
                i, v2[i], b)
        }
    }
}

//TestBitsLenMore64 tests function Bits for vector of length more than 64
func TestBitsLenMore64(t *testing.T) {
    v3 := []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    }
    v := New(v3)
    for i, b := range v.Bits() {
        if v3[i] != b {
            t.Errorf("vector testing: v3.Bits() is incorrect, (v3[%d] = %d) != %d",
                i, v3[i], b)
        }
    }
}

//TestIter tests function for iterating over vector's bits
func TestIter(t *testing.T) {
    v3 := []uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    }
    v := New(v3)
    i := 0
    for b := range v.Iter() {
        if v3[i] != b {
            t.Errorf("vector testing: v3.Bits() is incorrect, (v3[%d] = %d) != %d",
                i, v3[i], b)
        }
        i++
    }
}
