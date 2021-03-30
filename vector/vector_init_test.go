package vector

import "testing"

//TestNewNil tests function New(nil)
func TestNewNil(t *testing.T) {
    v := New(nil)
    if len(v.body) != 0 {
        t.Errorf("vector testing: default init is incorrect, len(body) != 0 (%d != 0)", len(v.body))
    }
    if v.lenLast != 0 {
        t.Errorf("vector testing: default init is incorrect, lenLast != 0 (%d != 0)", v.lenLast)
    }
}

//TestNewEmpty tests function New([]uint8{})
func TestNewEmpty(t *testing.T) {
    v := New([]uint8{})
    if len(v.body) != 0 {
        t.Errorf("vector testing: default init is incorrect, len(body) != 0 (%d != 0)", len(v.body))
    }
    if v.lenLast != 0 {
        t.Errorf("vector testing: default init is incorrect, lenLast != 0 (%d != 0)", v.lenLast)
    }
}

//TestNewLen1 tests function New([]uint8{0})
func TestNewLen1(t *testing.T) {
    v := New([]uint8{0})
    if len(v.body) != 1 {
        t.Errorf("vector testing: init by 0 is incorrect, len(body) != 1 (%d != 1)", len(v.body))
    }
    if v.lenLast != 1 {
        t.Errorf("vector testing: init by 0 is incorrect, lenLast != 1 (%d != 1)", v.lenLast)
    }
}

//TestNewLenLess64 tests function New(array of length less than 64)
func TestNewLenLess64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if len(v.body) != 1 {
        t.Errorf("vector testing: init is incorrect, len(body) != 1 (%d != 1)", len(v.body))
    }
    if v.body[0] != 0xffff00ff00000000 {
        t.Errorf("vector testing: init is incorrect, v.body[0] != 0xffff00ff00000000 (%x != 0xffff00ff00000000)", v.body[0])
    }
    if v.lenLast != 32 {
        t.Errorf("vector testing: init is incorrect, lenLast != 1 (%d != 1)", v.lenLast)
    }
}

//TestNewLenEqual64 tests function New(array of length 64)
func TestNewLenEqual64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
    })
    if len(v.body) != 1 {
        t.Errorf("vector testing: init is incorrect, len(body) != 1 (%d != 1)", len(v.body))
    }
    if v.body[0] != 0xffff00ff0f0f3333 {
        t.Errorf("vector testing: init is incorrect, v.body[0] != 0xffff00ff0f0f3333 (%x != ffff00ff0f0f3333)", v.body[0])
    }
    if v.lenLast != 0 {
        t.Errorf("vector testing: init is incorrect, lenLast != 0 (%d != 0)", v.lenLast)
    }
}

//TestNewLenMore64 tests function New(array of length more than 64)
func TestNewLenMore64(t *testing.T) {
    v := New([]uint8{
        1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1,
        0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1,
        0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
        0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    })
    if len(v.body) != 2 {
        t.Errorf("vector testing: init is incorrect, len(body) != 2 (%d != 2)", len(v.body))
    }
    if v.body[0] != 0xffff00ff0f0f3333 {
        t.Errorf("vector testing: init is incorrect, v.body[0] != 0xffff00ff0f0f3333 (%x != ffff00ff0f0f3333)", v.body[0])
    }
    if v.body[1] != 0x55550ff800000000 {
        t.Errorf("vector testing: init is incorrect, v.body[1] != 0x55550ff800000000 (%x != 55550ff800000000)", v.body[1])
    }
    if v.lenLast != 29 {
        t.Errorf("vector testing: init is incorrect, lenLast != 29 (%d != 29)", v.lenLast)
    }
}

func TestPackBytes(t *testing.T) {
    b := []byte{
        0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
        0xEF, 0x98,
    }
    v := PackBytes(b, len(b)*8)
    exps := "00000001001000110100010101100111100010011010101111001101111011111110111110011000"
    if v.Len() != len(b) * 8 {
        t.Errorf("expected len(v)=%d, but got %d", len(b) * 8, v.Len())
    }
    if v.String() != exps {
        t.Errorf("expected v=\n%s, but got\n%s", exps, v.String())
    }

    v = PackBytes(b, 61)
    exps = "0000000100100011010001010110011110001001101010111100110111101"
    if v.Len() != 61 {
        t.Errorf("expected len(v)=%d, but got %d", 61, v.Len())
    }
    if v.String() != exps {
        t.Errorf("expected v=\n%s, but got\n%s", exps, v.String())
    }

    v = PackBytes(b, 65)
    exps = "00000001001000110100010101100111100010011010101111001101111011111"
    if v.Len() != 65 {
        t.Errorf("expected len(v)=%d, but got %d", 65, v.Len())
    }
    if v.String() != exps {
        t.Errorf("expected v=\n%s, but got\n%s", exps, v.String())
    }

    v = PackBytes(b[5:6], 3)
    exps = "101"
    if v.Len() != 3 {
        t.Errorf("expected len(v)=%d, but got %d", 3, v.Len())
    }
    if v.String() != exps {
        t.Errorf("expected v=\n%s, but got\n%s", exps, v.String())
    }

    v = PackBytes(make([]byte,0), 5)
    exps = "00000"
    if v.Len() != 5 {
        t.Errorf("expected len(v)=%d, but got %d", 5, v.Len())
    }
    if v.String() != exps {
        t.Errorf("expected v=\n%s, but got\n%s", exps, v.String())
    }
}


