package vector

import (
    "testing"
)

//TestLenNil tests function Len for nil vector
func TestLenNil(t *testing.T) {
    v, _ := New(nil)
    if v.Len() != 0 {
        t.Errorf("vector testing: Len() is incorrect, nil.Len() != 0 (%d != 0)",
            v.Len())
    }
}

//TestLenEmpty tests function Len for empty vector
func TestLenEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if v.Len() != 0 {
        t.Errorf("vector testing: Len() is incorrect, []uint8{}.Len() != 0 (%d != 0)",
            v.Len())
    }
}

//TestLenLen1 tests function Len for vector of length 1
func TestLenLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    if v.Len() != 1 {
        t.Errorf("vector testing: Len() is incorrect, []uint8{0}.Len() != 1 (%d != 1)",
            v.Len())
    }
}

//TestLenLenLess64 tests function Len for vector of length less than 64
func TestLenLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if v.Len() != 32 {
        t.Errorf("vector testing: Len() is incorrect, v1.Len() != 32 (%d != 32)",
            v.Len())
    }
}

//TestLenLen64 tests function Len for vector of length 64
func TestLenLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    if v.Len() != 64 {
        t.Errorf("vector testing: Len() is incorrect, v2.Len() != 64 (%d != 64)",
            v.Len())
    }
}

//TestLenLenMore64 tests function Len for vector of length more than 64
func TestLenLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if v.Len() != 93 {
        t.Errorf("vector testing: Len() is incorrect, v3.Len() != 93 (%d != 93)",
            v.Len())
    }
}

//TestSupportNil tests function Support for nil vector
func TestSupportNil(t *testing.T) {
    v, _ := New(nil)
    if len(v.Support()) != 0 {
        t.Errorf("vector testing: Support() is incorrect, nil.Support() != [] (%v != [])",
            v.Support())
    }
}

//TestSupportEmpty tests function Support for empty vector
func TestSupportEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if len(v.Support()) != 0 {
        t.Errorf("vector testing: Support() is incorrect, []uint8{}.Support() != [] (%v != [])",
            v.Support())
    }
}

//TestSupportLen1 tests function Support for vector of length 1
func TestSupportLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    if len(v.Support()) != 0 {
        t.Errorf("vector testing: Support() is incorrect, []uint8{0}.Support() != [] (%v != [])",
            v.Support())
    }
    v, _ = New([]uint8{1})
    sup := v.Support()
    if len(sup) != 1 && sup[0] != 0 {
        t.Errorf("vector testing: Support() is incorrect, []uint8{1}.Support() != [0] (%v != [0])",
            v.Support())
    }
}

//TestSuppLenLess64 tests function Support for vector of length less than 64
func TestSuppLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    sup := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 24, 25, 26, 27, 28, 29, 30, 31}
    vsup := v.Support()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestSuppLen64 tests function Support for vector of length 64
func TestSuppLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    sup := []int{
        0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
        24, 25, 26, 27, 28, 29, 30, 31,
        36, 37, 38, 39, 44, 45, 46, 47,
        50, 51, 54, 55, 58, 59, 62, 63}
    vsup := v.Support()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestSuppLenMore64 tests function Support for vector of length more than 64
func TestSuppLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    sup := []int{
        0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
        24, 25, 26, 27, 28, 29, 30, 31,
        36, 37, 38, 39, 44, 45, 46, 47,
        50, 51, 54, 55, 58, 59, 62, 63,
        65, 67, 69, 71, 73, 75, 77, 79,
        84, 85, 86, 87, 88, 89, 90, 91,
        92,
    }
    vsup := v.Support()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Support() is incorrect, v.Support() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestZeroesNil tests function Zeroes for nil vector
func TestZeroesNil(t *testing.T) {
    v, _ := New(nil)
    if len(v.Zeroes()) != 0 {
        t.Errorf("vector testing: Zeroes() is incorrect, nil.Zeroes() != [] (%v != [])",
            v.Zeroes())
    }
}

//TestZeroesEmpty tests function Zeroes for empty vector
func TestZeroesEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if len(v.Zeroes()) != 0 {
        t.Errorf("vector testing: Zeroes() is incorrect, []uint8{}.Zeroes() != [] (%v != [])",
            v.Zeroes())
    }
}

//TestZeroesLen1 tests function Zeroes for vector of length 1
func TestZeroesLen1(t *testing.T) {
    v, _ := New([]uint8{1})
    if len(v.Zeroes()) != 0 {
        t.Errorf("vector testing: Zeroes() is incorrect, []uint8{1}.Zeroes() != [] (%v != [])",
            v.Zeroes())
    }
    v, _ = New([]uint8{0})
    sup := v.Zeroes()
    if len(sup) != 1 && sup[0] != 0 {
        t.Errorf("vector testing: Zeroes() is incorrect, []uint8{0}.Zeroes() != [0] (%v != [0])",
            v.Zeroes())
    }
}

//TestZeroesLenLess64 tests function Zeroes for vector of length less than 64
func TestZeroesLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    sup := []int{16, 17, 18, 19, 20, 21, 22, 23}
    vsup := v.Zeroes()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestZeroesLen64 tests function Zeroes for vector of length 64
func TestZeroesLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    sup := []int{
        16, 17, 18, 19, 20, 21, 22, 23,
        32, 33, 34, 35, 40, 41, 42, 43,
        48, 49, 52, 53, 56, 57, 60, 61}
    vsup := v.Zeroes()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestZeroesLenMore64 tests function Zeroes for vector of length more than 64
func TestZeroesLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    sup := []int{
        16, 17, 18, 19, 20, 21, 22, 23,
        32, 33, 34, 35, 40, 41, 42, 43,
        48, 49, 52, 53, 56, 57, 60, 61,
        64, 66, 68, 70, 72, 74, 76, 78,
        80, 81, 82, 83,
    }
    vsup := v.Zeroes()
    if len(vsup) != len(sup) {
        t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
            sup, vsup, sup)
    }
    for i, b := range vsup {
        if sup[i] != b {
            t.Errorf("vector testing: Zeroes() is incorrect, v.Zeroes() != %v (%v != %v)",
                sup, vsup, sup)
        }
    }
}

//TestWtNil tests function Wt for nil vector
func TestWtNil(t *testing.T) {
    v, _ := New(nil)
    if v.Wt() != 0 {
        t.Errorf("vector testing: Wt() is incorrect, nil.Wt() != 0 (%d != 0)",
            v.Wt())
    }
}

//TestWtEmpty tests function Wt for empty vector
func TestWtEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if v.Wt() != 0 {
        t.Errorf("vector testing: Wt() is incorrect, []uint8{}.Wt() != 0 (%d != 0)",
            v.Wt())
    }
}

//TestWtLen1 tests function Wt for vector of length 1
func TestWtLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    if v.Wt() != 0 {
        t.Errorf("vector testing: Wt() is incorrect, []uint8{0}.Wt() != 0 (%d != 0)",
            v.Wt())
    }
}

//TestWtLenLess64 tests function Wt for vector of length less than 64
func TestWtLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if v.Wt() != 24 {
        t.Errorf("vector testing: Wt() is incorrect, v.Wt() != 24 (%d != 24)",
            v.Wt())
    }
}

//TestWtLen64 tests function Wt for vector of length 64
func TestWtLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    if v.Wt() != 40 {
        t.Errorf("vector testing: Wt() is incorrect, v2.Wt() != 40 (%d != 40)",
            v.Wt())
    }
}

//TestWtLenMore64 tests function Wt for vector of length more than 64
func TestWtLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if v.Wt() != 57 {
        t.Errorf("vector testing: Wt() is incorrect, v3.Len() != 57 (%d != 57)",
            v.Wt())
    }
}
