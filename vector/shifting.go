package vector

//ShiftRight shifts vector for r position right
//Return pointer to the shifted vector
func (v *Vector) ShiftRight(r uint) *Vector {
    res := newEmpty(v.Len())
    start := int(r / wordSize)
    l := r % wordSize
    mask := uint64((1 << l) - 1)
    for i := start; i < len(res.body); i++ {
        res.body[i] = (v.body[i-start] >> l)
        if i > start && mask != 0 {
            res.body[i] ^= ((v.body[i-start-1] & mask) << (wordSize - l))
        }
        if i == len(res.body)-1 && v.lenLast != 0 {
            res.body[i] &= (((1 << v.lenLast) - 1) << (wordSize - v.lenLast))
        }
    }
    return res
}

//ShiftLeft shifts vector for r position left
//Return pointer to the shifted vector
func (v *Vector) ShiftLeft(r uint) *Vector {
    res := newEmpty(v.Len())
    start := int(r / wordSize)
    l := r % wordSize
    mask := uint64((1<<l)-1) << (wordSize - l)
    for i := start; i < len(res.body); i++ {
        res.body[i-start] = (v.body[i] << l)
        if i < len(res.body)-1 {
            res.body[i-start] ^= ((v.body[i+1] & mask) >> (wordSize - l))
        }
    }
    return res
}

//RotateLeft cyclical shifts vector for r position left
//Return pointer to the shifted vector
func (v *Vector) RotateLeft(r uint) *Vector {
    if v.Len() == 0 {
        r = 0
    } else {
        r = uint(r % uint(v.Len()))
    }
    return v.ShiftLeft(r).Xor(v.ShiftRight(uint(v.Len()) - r))
}

//RotateRight cyclical shifts vector for r position right
//Return pointer to the shifted vector
func (v *Vector) RotateRight(r uint) *Vector {
    if v.Len() == 0 {
        r = 0
    } else {
        r = uint(r % uint(v.Len()))
    }
    return v.ShiftLeft(uint(v.Len()) - r).Xor(v.ShiftRight(r))
}
