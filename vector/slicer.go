package vector

import "fmt"

//Slicer is abstraction of integer slice
type Slicer interface {
    Len() int
    GetSlice(int, int) Slicer
    GetElement(int) int
}

type intArray []int

func (a intArray) Len() int {
    return len(a)
}

func (a intArray) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a intArray) GetElement(i int) int {
    return int(a[i])
}

type int8Array []int8

func (a int8Array) Len() int {
    return len(a)
}
func (a int8Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}
func (a int8Array) GetElement(i int) int {
    return int(a[i])
}

type int16Array []int16

func (a int16Array) Len() int {
    return len(a)
}

func (a int16Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a int16Array) GetElement(i int) int {
    return int(a[i])
}

type int32Array []int32

func (a int32Array) Len() int {
    return len(a)
}

func (a int32Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a int32Array) GetElement(i int) int {
    return int(a[i])
}

type int64Array []int64

func (a int64Array) Len() int {
    return len(a)
}

func (a int64Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a int64Array) GetElement(i int) int {
    return int(a[i])
}

type uintArray []uint

func (a uintArray) Len() int {
    return len(a)
}

func (a uintArray) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a uintArray) GetElement(i int) int {
    return int(a[i])
}

type uint8Array []uint8

func (a uint8Array) Len() int {
    return len(a)
}

func (a uint8Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a uint8Array) GetElement(i int) int {
    return int(a[i])
}

type uint16Array []uint16

func (a uint16Array) Len() int {
    return len(a)
}

func (a uint16Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a uint16Array) GetElement(i int) int {
    return int(a[i])
}

type uint32Array []uint32

func (a uint32Array) Len() int {
    return len(a)
}

func (a uint32Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a uint32Array) GetElement(i int) int {
    return int(a[i])
}

type uint64Array []uint64

func (a uint64Array) Len() int {
    return len(a)
}

func (a uint64Array) GetSlice(start, end int) Slicer {
    return a[start:end]
}

func (a uint64Array) GetElement(i int) int {
    return int(a[i])
}

//ToSlicer convert integers array to slice
func ToSlicer(array interface{}) (Slicer, error) {
    switch a := array.(type) {
    case []int:
        return intArray(a), nil
    case []int8:
        return int8Array(a), nil
    case []int16:
        return int16Array(a), nil
    case []int32:
        return int32Array(a), nil
    case []int64:
        return int64Array(a), nil
    case []uint:
        return uintArray(a), nil
    case []uint8:
        return uint8Array(a), nil
    case []uint16:
        return uint16Array(a), nil
    case []uint32:
        return uint32Array(a), nil
    case []uint64:
        return uint64Array(a), nil
    case Slicer:
        return a, nil
    default:
        return intArray([]int{}), fmt.Errorf("vector: cannot convert type %T to Slicer", array)
    }
}
