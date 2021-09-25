// Package bitarray implements a simple bit-array data structure.
// The implementation is memory efficient as bits are actually stored in
// on memory bit (excluding a constant overhead).

package bitarray

import (
	"fmt"
	"strconv"
)

const uintSize = 32 << (^uint(0) >> 63) // 32 or 64
// UintSize is the size of a uint in bits (32 or 64 depending on the platform).
const UintSize = uintSize

// BitArray is the actual data structure. Users are supposed to use `New` method to instantiate a new bit-array
type BitArray struct {
	data    []byte
	padding int
}

// New returns a new, ready to be used, empty bit-array.
func New() *BitArray {
	return &BitArray{
		data:    []byte{0},
		padding: 8,
	}
}

// Len returns the length (number of bits) of the bit-array.
func (ba *BitArray) Len() int {
	return (len(ba.data) << 3) - int(ba.padding)
}

// Bytes returns the underlying bit data as a slice of bytes.
// Note that this method returns a copy, hence modifications on the returned slice are not reflected on the content of the bit-array.
// As the number of bits might not be a multiple of 8,
// the slice is zero padded and the user should rely on Len to infer the number of padding bits.
func (ba *BitArray) Bytes() []byte {
	if ba.Len() == 0 {
		return []byte{}
	}

	data := make([]byte, len(ba.data))
	copy(data, ba.data)
	return data
}

// Padding returns the number of padding bits at the end of the byte slice returned by the Bytes method/
// It is guaranteed that padding is between 0 and 7 (included) and 0 if the bit-array is empty
func (ba *BitArray) Padding() int {
	if ba.Len() == 0 {
		return 0
	}
	return ba.padding
}

// AppendOne appends a `1` to the bit array.
func (ba *BitArray) AppendOne() {
	if ba.padding != 0 {
		ba.padding -= 1
		ba.data[len(ba.data)-1] |= (1 << ba.padding)
		return
	}

	ba.data = append(ba.data, 1<<7)
	ba.padding = 7
}

// AppendZero appends a `0` to the bit array.
func (ba *BitArray) AppendZero() {
	if ba.padding != 0 {
		ba.padding -= 1
		return
	}

	ba.data = append(ba.data, 0)
	ba.padding = 7
}

// AppendBit appends a `0` or `1` depending on the value of the bit argument.
// If bit is neither 0 nor 1, AppendBit will panic.
func (ba *BitArray) AppendBit(bit byte) {
	if bit == 0 {
		ba.AppendZero()
	} else if bit == 1 {
		ba.AppendOne()
	} else {
		panic(fmt.Sprintf("bit should be 0 or 1, given %d", bit))
	}
}

// GetBit returns the bit at position `index` and will panic if index is out of range.
func (ba *BitArray) GetBit(index int) byte {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := (index &^ 0x7) >> 3
	r := index - (b << 3)

	return (ba.data[b] & (0b10000000 >> r)) >> (^r & 0x7)
}

// SetBit sets the bit at position `index` to `1` and will panic if index is out of range.
func (ba *BitArray) SetBit(index int) {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := (index &^ 0x7) >> 3
	r := index - (b << 3)

	ba.data[b] |= 0b10000000 >> r
}

// ClearBit clears the bit at position `index` (sets it to `0`) and panics if index is out of range.
func (ba *BitArray) ClearBit(index int) {
	if index >= ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", index, ba.Len()))
	}
	b := (index &^ 0x7) >> 3
	r := index - (b << 3)

	ba.data[b] &^= 0b10000000 >> r
}

// Append appends the `nbBits` lowest bits stored in v (of type uint).
// It will panic if nbBits is larger than 32 or 64 depending on the size of uint on the running machine.
func (ba *BitArray) Append(v uint, nbBits int) {
	if UintSize == 32 {
		ba.Append32(uint32(v), nbBits)
		return
	}

	ba.Append64(uint64(v), nbBits)

}

// Append8 appends the `nbBits` lowest bits stored in v (of type uint8).
// It will panic if nBbits is larger than 8.
func (ba *BitArray) Append8(v uint8, nbBits int) {
	if nbBits < 0 || nbBits > 8 {
		panic(fmt.Sprintf("nbBits should not be between 0 and 8, given %d", nbBits))
	}

	v = v & (0b11111111 >> (8 - nbBits))

	if nbBits <= ba.padding {
		ba.data[len(ba.data)-1] |= v << (ba.padding - nbBits)
		ba.padding -= nbBits
		return
	}

	ba.data[len(ba.data)-1] |= v >> (nbBits - ba.padding)
	ba.padding = (8 - nbBits + ba.padding)
	ba.data = append(ba.data, v<<ba.padding)
}

// Append16 appends the `nbBits` lowest bits stored in v (of type uint16).
// It will panic if nBbits is larger than 16.
func (ba *BitArray) Append16(v uint16, nbBits int) {
	if nbBits < 0 || nbBits > 16 {
		panic(fmt.Sprintf("nbBits should not be between 0 and 16, given %d", nbBits))
	}

	if nbBits > 8 {
		ba.Append8(uint8(v>>8), nbBits-8)
		nbBits = 8
	}
	ba.Append8(uint8(v), nbBits)
}

// Append32 appends the `nbBits` lowest bits stored in v (of type uint32).
// It will panic if nBbits is larger than 32.
func (ba *BitArray) Append32(v uint32, nbBits int) {
	if nbBits < 0 || nbBits > 32 {
		panic(fmt.Sprintf("nbBits should not be between 0 and 32, given %d", nbBits))
	}

	if nbBits > 16 {
		ba.Append16(uint16(v>>16), nbBits-16)
		nbBits = 16
	}

	ba.Append16(uint16(v), nbBits)

}

// Append64 appends the `nbBits` lowest bits stored in v (of type uint64).
// It will panic if nBbits is larger than 64.
func (ba *BitArray) Append64(v uint64, nbBits int) {
	if nbBits < 0 || nbBits > 64 {
		panic(fmt.Sprintf("nbBits should not be between 0 and 64, given %d", nbBits))
	}

	if nbBits > 32 {
		ba.Append32(uint32(v>>32), nbBits-32)
		nbBits = 32
	}

	ba.Append32(uint32(v), nbBits)

}

// AppendBytes appends a slice of bytes to the bit-array where
// padding represents the number of padding bits in the last byte of the input slice.
// Padding must be between 0 and 7 (included) and 0 if the slice of bytes is empty, otherwise AppendBytes will panic
func (ba *BitArray) AppendBytes(bytes []byte, padding int) {
	if padding < 0 || padding > 8 {
		panic(fmt.Sprintf("padding should be between 0 and 7; given %d", padding))
	}

	if len(bytes) == 0 {
		if padding != 0 {
			panic(fmt.Sprintf("input byte slice is empty but padding is not 0: %d", padding))
		}
		return
	}

	for i := 0; i < len(bytes)-1; i++ {
		ba.Append8(bytes[i], 8)
	}

	lastByte := bytes[len(bytes)-1]
	ba.Append8(lastByte>>byte(padding), 8-padding)

}

// AppendBitArray appends another bit-array to the receiving one
func (ba *BitArray) AppendBitArray(ba1 *BitArray) {
	ba.AppendBytes(ba1.Bytes(), ba1.Padding())
}

// AppendString appends a stringified bit sequence to the bit-array.
// It will panic if the bit sequence is not valid (consisting only of 0's and 1's).
func (ba *BitArray) AppendString(bitSeq string) {
	pieces64 := (len(bitSeq) &^ 0x111111) >> 6
	r := len(bitSeq) - (pieces64 << 6)
	for i := 0; i < pieces64; i++ {
		v, err := strconv.ParseUint(bitSeq[i<<6:(i+1)<<6], 2, 64)
		if err != nil {
			panic(fmt.Sprintf("the bit sequence appears to be invalid: %v", err))
		}

		ba.Append64(v, 64)
	}

	if r != 0 {
		v, err := strconv.ParseUint(bitSeq[pieces64<<6:], 2, 64)
		if err != nil {
			panic(fmt.Sprintf("the bit sequence appears to be invalid: %v", err))
		}

		ba.Append64(v, r)
	}
}

// Extract extracts a range defined by [i, j] from the bit-array into a uint64.
// Semantics of range are pretty similar to slice indexing in golang,
// the bit at position i is included, the bit at position j is excluded.
// Indexes must not be negative, j must be strictly greater than i and
// the number of bits in the range should not exceed 64 (in order to fit in a uint64), otherwise Extract will panic.
// The returned uint64 is filled from left to right. Left most empty bits are filled with 0.
func (ba *BitArray) Extract(i, j int) uint64 {
	if i < 0 || j < 0 {
		panic(fmt.Sprintf("negative indexes are invalid; given (i=%d, j=%d)", i, j))
	}
	if i >= j {
		panic(fmt.Sprintf("invalid indexes %d >= %d", i, j))
	}
	if j > ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", j, ba.Len()))
	}
	if j-i > 64 {
		panic(fmt.Sprintf("the number of queried bits should not be greater than 64 bits; j - i = %d", j-i))
	}

	var result uint64
	startingByte := (i &^ 0x7) >> 3
	endingByte := ((j - 1) &^ 0x7) >> 3
	i, j = i&0x7, (j-1)&0x7

	b := ba.data[startingByte]
	b &= 0xff >> i
	result |= uint64(b)

	if startingByte == endingByte {
		result = result >> (^j & 0x7)
		return result
	}

	for k := startingByte + 1; k < endingByte; k++ {
		b = ba.data[k]
		result = (result << 8) | uint64(b)
	}

	result = (result << (j + 1)) | uint64(ba.data[endingByte]>>(7-j))

	return result
}

// ExtractBitArray extracts a range defined by [i, j] from the bit-array into a new bit-array.
// Semantics of range are pretty similar to slice indexing in golang,
// the bit at position i is included, the bit at position j is excluded.
// Indexes must not be negative, j must be greater or equal to i and j must not be out of range, otherwise ExtractBitArray will panic.
func (ba *BitArray) ExtractBitArray(i, j int) *BitArray {
	if i < 0 || j < 0 {
		panic(fmt.Sprintf("negative indexes are invalid; given (i=%d, j=%d)", i, j))
	}
	if i > j {
		panic(fmt.Sprintf("invalid indexes %d > %d", i, j))
	}
	if j > ba.Len() {
		panic(fmt.Sprintf("bit index out of range [%d] with length %d", j, ba.Len()))
	}

	if j == i {
		return New()
	}

	startingByte := (i &^ 0x7) >> 3
	endingByte := ((j - 1) &^ 0x7) >> 3
	i, j = i&0x7, (j-1)&0x7

	data := make([]byte, 1, endingByte-startingByte+1)
	res := &BitArray{data: data, padding: 8}

	b := ba.data[startingByte]
	b &= 0xff >> i
	res.Append8(b, 8-i)

	if startingByte == endingByte {
		res.data[0] &^= 0b11111111 >> (j - i + 1)
		res.padding += 7 - j
		return res
	}

	for k := startingByte + 1; k < endingByte; k++ {
		res.Append8(ba.data[k], 8)
	}

	res.Append8(ba.data[endingByte]>>(7-j), j+1)
	return res
}
