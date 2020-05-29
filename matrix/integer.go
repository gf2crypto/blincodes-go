package matrix

import "fmt"

func toInt(num interface{}) (int, error) {
    switch n := num.(type) {
    case int:
        return n, nil
    case int8:
        return int(n), nil
    case int16:
        return int(n), nil
    case int32:
        return int(n), nil
    case int64:
        return int(n), nil
    case uint:
        return int(n), nil
    case uint8:
        return int(n), nil
    case uint16:
        return int(n), nil
    case uint32:
        return int(n), nil
    case uint64:
        return int(n), nil
    default:
        return 0, fmt.Errorf("matrix: cannot convert type %T to int", num)
    }
}
