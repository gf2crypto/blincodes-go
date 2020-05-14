package vector

import (
    "testing"
)

//TestPrettyStringNil tests function PrettyString for nil vector
func TestPrettyStringNil(t *testing.T) {
    v, _ := New(0)
    if v.PrettyString() != "" {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != \"\"",
            v.PrettyString())
    }
}

//TestPrettyStringEmpty tests function PrettyString for empty vector
func TestPrettyStringEmpty(t *testing.T) {
    v, _ := New([]byte{})
    if v.PrettyString() != "" {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != \"\"",
            v.PrettyString())
    }
}

//TestPrettyStringLen1 tests function PrettyString for vector of length 1
func TestPrettyStringLen1(t *testing.T) {
    v, _ := New([]uint32{0})
    if v.PrettyString() != "-" {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != \"\"",
            v.PrettyString())
    }
    v, _ = New([]uint32{1})
    if v.PrettyString() != "1" {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != \"\"",
            v.PrettyString())
    }
}

//TestPrettyStringLenLess64 tests function PrettyString for vector of length less than 64
func TestPrettyStringLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := "1111111111111111--------11111111"
    if v.PrettyString() != res {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != %s",
            v.PrettyString(), res)
    }
}

//TestPrettyStringLen64 tests function PrettyString for vector of length 64
func TestPrettyStringLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    res := "1111111111111111--------11111111----1111----1111--11--11--11--11"
    if v.PrettyString() != res {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != %s",
            v.PrettyString(), res)
    }
}

//TestPrettyStringLenMore64 tests function PrettyString for vector of length more than 64
func TestPrettyStringLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    res := "1111111111111111--------11111111----1111----1111--11--11--11--11-1-1-1-1-1-1-1-1----111111111"
    if v.PrettyString() != res {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != %s",
            v.PrettyString(), res)
    }
    v, _ = New([]uint8{
        0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
    })
    res = "-1111111111111111--------11111111----1111----1111--11--11--11--11-1-1-1-1-1-1-1-1----111111111-"
    if v.PrettyString() != res {
        t.Errorf("vector testing: PrettyString() is incorrect, %s != %s",
            v.PrettyString(), res)
    }
}
