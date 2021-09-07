package bitarray

import "fmt"

const len8tab = "" +
	"\x00\x01\x02\x02\x03\x03\x03\x03\x04\x04\x04\x04\x04\x04\x04\x04" +
	"\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05\x05" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06\x06" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07\x07" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08" +
	"\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08\x08"

type BitArray struct {
	data    []byte
	padding byte
}

func New() *BitArray {
	return &BitArray{
		data:    []byte{0},
		padding: 8,
	}
}

func (ba *BitArray) Len() int {
	return 8*len(ba.data) - int(ba.padding)
}

func (ba *BitArray) Bytes() []byte {
	data := make([]byte, len(ba.data))
	copy(data, ba.data)
	return data
}

func (ba *BitArray) Append8(v uint8) {
	ba.append8(v, false)
}

func (ba *BitArray) Append16(v uint16) {
	b2, b1 := byte(v>>8), byte(v&0xff)

	if b2 != 0 {
		ba.append8(b2, false)
		ba.append8(b1, true)
		return
	}

	ba.append8(b1, false)

}

func (ba *BitArray) Append32(v uint32) {
	withTrailing := false

	if b4 := byte(v >> 24); b4 != 0 {
		ba.append8(b4, false)
		withTrailing = true
	}

	if b3 := byte((v >> 16) & 0xff); b3 != 0 || withTrailing {
		ba.append8(b3, withTrailing)
		withTrailing = true
	}

	if b2 := byte((v >> 8) & 0xff); b2 != 0 || withTrailing {
		ba.append8(b2, withTrailing)
		withTrailing = true
	}

	if b1 := byte(v & 0xff); b1 != 0 || withTrailing {
		ba.append8(b1, withTrailing)
	}
}

func (ba *BitArray) Append64(v uint64) {
	withTrailing := false

	if b8 := byte(v >> 56); b8 != 0 {
		ba.append8(b8, false)
		withTrailing = true
	}

	if b7 := byte((v >> 48) & 0xff); b7 != 0 || withTrailing {
		ba.append8(b7, withTrailing)
		withTrailing = true
	}

	if b6 := byte((v >> 40) & 0xff); b6 != 0 || withTrailing {
		ba.append8(b6, withTrailing)
		withTrailing = true
	}

	if b5 := byte((v >> 32) & 0xff); b5 != 0 || withTrailing {
		ba.append8(b5, withTrailing)
	}

	if b4 := byte(v >> 24); b4 != 0 {
		ba.append8(b4, false)
		withTrailing = true
	}

	if b3 := byte((v >> 16) & 0xff); b3 != 0 || withTrailing {
		ba.append8(b3, withTrailing)
		withTrailing = true
	}

	if b2 := byte((v >> 8) & 0xff); b2 != 0 || withTrailing {
		ba.append8(b2, withTrailing)
		withTrailing = true
	}

	if b1 := byte(v & 0xff); b1 != 0 || withTrailing {
		ba.append8(b1, withTrailing)
	}
}

func (ba *BitArray) AppendOne() {
	if ba.padding != 0 {
		ba.padding -= 1
		ba.data[len(ba.data)-1] |= (1 << ba.padding)
		return
	}

	ba.data = append(ba.data, 1<<7)
	ba.padding = 7
}

func (ba *BitArray) AppendZero() {
	if ba.padding != 0 {
		ba.padding -= 1
		return
	}

	ba.data = append(ba.data, 0)
	ba.padding = 7
}

func (ba *BitArray) GetBit(index int) byte {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := index / 8
	r := index - 8*b

	return (ba.data[b] & (0b10000000 >> r)) >> (7 - r)
}

func (ba *BitArray) SetBit(index int) {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := index / 8
	r := index - 8*b

	ba.data[b] |= 0b10000000 >> r
}

func (ba *BitArray) ClearBit(index int) {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := index / 8
	r := index - 8*b

	ba.data[b] &^= 0b10000000 >> r
}

func (ba *BitArray) append8(v uint8, withTrailing bool) {
	data, padding := ba.data, ba.padding
	var rb1, rb2, rp byte
	var oneByte bool

	if withTrailing {
		rb1, rb2, oneByte, rp = mergeWithTrailing(data[len(data)-1], padding, v)
	} else {
		rb1, rb2, oneByte, rp = mergeWithoutTrailing(data[len(data)-1], padding, v)
	}

	data[len(data)-1] = rb1
	if !oneByte {
		data = append(data, rb2)
	}

	ba.data = data
	ba.padding = rp
}

func mergeWithoutTrailing(b1, padding, b2 byte) (rb1 byte, rb2 byte, oneByte bool, rp byte) {
	if b2 == 0 {
		return b1, 0, true, padding - 1
	}

	b2Len := len8tab[b2]
	if b2Len <= padding {
		rb1 = b1 | b2<<(padding-b2Len)
		return rb1, 0, true, padding - b2Len
	}

	rb1 = b1 | b2>>(b2Len-padding)

	rp = (8 - b2Len + padding)
	rb2 = b2 << rp
	return rb1, rb2, false, rp
}

func mergeWithTrailing(b1, padding, b2 byte) (rb1 byte, rb2 byte, oneByte bool, rp byte) {
	if padding == 8 {
		return b2, 0, true, 0
	}

	rb1 = b1 | b2>>(8-padding)
	rb2 = b2 << padding

	return rb1, rb2, false, padding
}
