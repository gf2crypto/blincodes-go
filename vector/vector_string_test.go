package vector

import (
    "fmt"
    "testing"
)

//TestStringNil tests function String for nil vector
func TestStringNil(t *testing.T) {
    v, _ := New(nil)
    if fmt.Sprint(v) != "" {
        t.Errorf("vector testing: String() is incorrect, str(nil) != \"\" (%s != \"\")",
            fmt.Sprint(v))
    }
}

//TestStringEmpty tests function String for empty vector
func TestStringEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if fmt.Sprint(v) != "" {
        t.Errorf("vector testing: String() is incorrect, str([]uint8{}) != \"\" (%s != \"\")",
            fmt.Sprint(v))
    }
}

//TestStringLen1 tests function String for vector of length 1
func TestStringLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    if fmt.Sprint(v) != "0" {
        t.Errorf("vector testing: String() is incorrect, str([]uint8{0}) != \"\" (%s != \"\")",
            fmt.Sprint(v))
    }
}

//TestStringLenLess64 tests function String for vector of length less than 64
func TestStringLenLess64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if fmt.Sprint(v) != "11111111111111110000000011111111" {
        t.Errorf("vector testing: String() is incorrect, str(v) != \"11111111111111110000000011111111\" (%s != \"11111111111111110000000011111111\")",
            fmt.Sprint(v))
    }
}

//TestStringLen64 tests function String for vector of length 64
func TestStringLen64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    if fmt.Sprint(v) != "1111111111111111000000001111111100001111000011110011001100110011" {
        t.Errorf("vector testing: String() is incorrect, str(v) != \"1111111111111111000000001111111100001111000011110011001100110011\" (%s != \"1111111111111111000000001111111100001111000011110011001100110011\")",
            fmt.Sprint(v))
    }
}

//TestStringLenMore64 tests function String for vector of length more than 64
func TestStringLenMore64(t *testing.T) {
    v, _ := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if fmt.Sprint(v) != "111111111111111100000000111111110000111100001111001100110011001101010101010101010000111111111" {
        t.Errorf("vector testing: String() is incorrect, str(v) != \"111111111111111100000000111111110000111100001111001100110011001101010101010101010000111111111\" (%s != \"111111111111111100000000111111110000111100001111001100110011001101010101010101010000111111111\")",
            fmt.Sprint(v))
    }
}
