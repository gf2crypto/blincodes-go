package vector

//Rotate cyclically shifts vector for r position
//Return pointer to the vector
// If r < 0 than Rotate makes left shift,
// If r > 0 than Rotate makes right shift
// Mnemonic rule:
//             << 5   >> 5
//      ------------|------------
//            -5    0    5
func (v *Vector) Rotate(r int) *Vector {
    if r == 0 || v.Len() == 0 {
        return v
    }
    if r > 0 {
        r = r % v.Len()
    } else {
        r = (v.Len() - ((-r) % v.Len())) % v.Len()
    }
    if r == 0 {
        return v
    }
    v.rotateRightForMultipleWordSize(uint(r / wordSize))
    rest := r % wordSize
    if rest != 0 {
        v.rotateRightForLessWordSize(r % wordSize)
    }
    return v
}

// RotateLeft cyclically shifts vector for r position left
func (v *Vector) RotateLeft(r uint) *Vector {
    return v.Rotate(-r)
}

//RotateRight cyclically shifts vector for r position right
func (v *Vector) RotateRight(r uint) *Vector {
    return v.Rotate(r)
}

//The function makes right multiple circular shifts for q*wordsize < length of vector
func (v *Vector) rotateRightForMultipleWordSize(q uint) {
    // It is supposed that v has len more than 0
    newbody := make([]uint64, len(v.body))
    for i := 0; i < len(v.body); i++ {
        newbody[(q+i)%len(v.body)] = v.body[i]
    }
    if v.lenLast != 0 {
        mask := ((1 << (wordSize - v.lenLast)) - 1)
        rest := newbody[len(v.body)-1] & mask
        newbody[len(v.body)-1] &= (((1 << lenLast) - 1) << (wordSize - lenLast))
        var tmp int
        for i := 0; i < q; i++ {
            tmp = newbody[i] & mask
            newbody >>= wordSize - v.lenLast
            newbody ^= rest
            rest = tmp
        }
    }
    v.body = newbody
}

//The function makes right multiple circular shifts for r in case of 0 < r < wordSize
func (v *Vector) rotateRightForLessWordSize(r uint) {
    // It is supposed that v has len more than 0
    mask := ((1 << r) - 1)
    var tmp int
    rest := 0
    end := len(v.body) - 1
    for i := 0; i < end; i++ {
        tmp = v.body[i] & mask
        v.body[i] >>= r
        v.body[i] ^= (rest << (wordSize - r))
        rest = tmp
    }
    if v.lenLast >= r {
        tmp = v.body[end] & (mask << (wordSize - lenLast))
        v.body[end] >>= r
        v.body[end] ^= (rest << (wordSize - r))
        // set unused bits to zero
        v.body[end] &= (((1 << v.lenLast) - 1) << (wordSize - v.lenLast))
        v.body[0] ^= (tmp << (v.lenLast - r))
    } else {
        // v.lenLast < r
        v.body[0] ^= (v.body[end] >> (r - v.lenLast))
        v.body[0] ^= ((rest & ((1 << (r - v.lenLast)) - 1)) << v.lenLast)
        v.body[end] = rest & (((1 << v.lenLast) - 1) << (r - v.lenLast))
    }
}
