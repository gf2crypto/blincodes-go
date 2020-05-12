package vector

import "testing"

//TestNewNil tests function New(nil)
func TestNewNil(t *testing.T) {
    v, _ := New(nil)
    if len(v.body) != 0 {
        t.Errorf("vector testing: default init is incorrect, len(body) != 0 (%d != 0)", len(v.body))
    }
    if v.lenLast != 0 {
        t.Errorf("vector testing: default init is incorrect, lenLast != 0 (%d != 0)", v.lenLast)
    }
}

//TestNewEmpty tests function New([]uint8{})
func TestNewEmpty(t *testing.T) {
    v, _ := New([]uint8{})
    if len(v.body) != 0 {
        t.Errorf("vector testing: default init is incorrect, len(body) != 0 (%d != 0)", len(v.body))
    }
    if v.lenLast != 0 {
        t.Errorf("vector testing: default init is incorrect, lenLast != 0 (%d != 0)", v.lenLast)
    }
}

//TestNewLen1 tests function New([]uint8{0})
func TestNewLen1(t *testing.T) {
    v, _ := New([]uint8{0})
    if len(v.body) != 1 {
        t.Errorf("vector testing: init by 0 is incorrect, len(body) != 1 (%d != 1)", len(v.body))
    }
    if v.lenLast != 1 {
        t.Errorf("vector testing: init by 0 is incorrect, lenLast != 1 (%d != 1)", v.lenLast)
    }
}

//TestNewLenLess64 tests function New(array of length less than 64)
func TestNewLenLess64(t *testing.T) {
    v, _ := New([]uint8{
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
    v, _ := New([]uint8{
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
    v, _ := New([]uint8{
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
