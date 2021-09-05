package bitarray

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMergeWithoutTrailing(t *testing.T) {
	tests := []struct {
		b1      byte
		paddin  byte
		b2      byte
		rb1     byte
		rb2     byte
		oneByte bool
		rp      byte
	}{
		{0b10010000, 4, 0b00000010, 0b10011000, 0b00000000, true, 2},
		{0b10010000, 4, 0b00101010, 0b10011010, 0b10000000, false, 6},
		{0b10000000, 7, 0b00000000, 0b10000000, 0b00000000, true, 6},
		{0b11100010, 1, 0b00010001, 0b11100011, 0b00010000, false, 4},
		{0b10000000, 5, 0b00001111, 0b10011110, 0b00000000, true, 1},
		{0b11111111, 0, 0b00100101, 0b11111111, 0b10010100, false, 2},
	}

	for _, test := range tests {
		rb1, rb2, oneByte, rp := mergeWithoutTrailing(test.b1, test.paddin, test.b2)
		if rb1 != test.rb1 || rb2 != test.rb2 || oneByte != test.oneByte || rp != test.rp {
			t.Errorf("mergeWithoutTrailing(%#08b, %d, %#08b)=[%#08b,%#08b, %t, %d]; want=[%#08b,%#08b, %t, %d]",
				test.b1, test.paddin, test.b2,
				rb1, rb2, oneByte, rp,
				test.rb1, test.rb2, test.oneByte, test.rp)
		}

	}
}

func TestMergeWithTrailing(t *testing.T) {
	tests := []struct {
		b1      byte
		paddin  byte
		b2      byte
		rb1     byte
		rb2     byte
		oneByte bool
		rp      byte
	}{
		{0b10010000, 4, 0b00000010, 0b10010000, 0b00100000, false, 4},
		{0b10010000, 4, 0b00101010, 0b10010010, 0b10100000, false, 4},
		{0b10000000, 7, 0b00000000, 0b10000000, 0b00000000, false, 7},
		{0b11100010, 1, 0b00010001, 0b11100010, 0b00100010, false, 1},
		{0b10000000, 5, 0b00001111, 0b10000001, 0b11100000, false, 5},
		{0b11111111, 0, 0b00100101, 0b11111111, 0b00100101, false, 0},
		{0b00000000, 8, 0b01010101, 0b01010101, 0b00000000, true, 0},
	}

	for _, test := range tests {
		rb1, rb2, oneByte, rp := mergeWithTrailing(test.b1, test.paddin, test.b2)
		if rb1 != test.rb1 || rb2 != test.rb2 || oneByte != test.oneByte || rp != test.rp {
			t.Errorf("mergeWithTrailing(%#08b, %d, %#08b)=[%#08b,%#08b, %t, %d]; want=[%#08b,%#08b, %t, %d]",
				test.b1, test.paddin, test.b2,
				rb1, rb2, oneByte, rp,
				test.rb1, test.rb2, test.oneByte, test.rp)
		}

	}
}

func TestAppend8(t *testing.T) {
	tests := []struct {
		idx     int
		ints    []uint8
		data    []byte
		padding byte
		len     int
	}{
		{0, []uint8{0xDE, 0xAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{1, []uint8{0xDE, 0xAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{2, []uint8{0xDE, 0xAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{3, []uint8{0xDE, 0xAD, 0x0B, 0x0E, 0x3B, 0x03}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{4, []uint8{0x07}, []byte("\xE0"), 5, 3},
		{5, []uint8{0xC0, 0x3F, 0x03, 0xEE}, []byte("\xC0\xFF\xEE"), 0, 24},
		{6, []uint8{0xC0, 0x3F, 0x03, 0xEE, 0xE}, []byte("\xC0\xFF\xEE\xE0"), 4, 28},
		{7, []uint8{0xDE, 0xAD, 0x07, 0x06, 0x05, 0x13, 0x02}, []byte("\xDE\xAD\xFA\xCE"), 0, 32},
		{8, []uint8{0xDE, 0xAD, 0x07, 0x06, 0x05, 0x13, 0x02, 0x1F}, []byte("\xDE\xAD\xFA\xCE\xF8"), 3, 37},
		{9, []uint8{0x1B, 0xF}, []byte("\xDF\x80"), 7, 9},
	}

	for _, test := range tests {
		ba := New()
		for _, e := range test.ints {
			ba.Append8(e)
		}

		data := ba.Bytes()

		if ba.padding != test.padding {
			t.Errorf("Append8 with test_idx: %d; returned bad padding %d, want: %d", test.idx, ba.padding, test.padding)
		}

		if ba.Len() != test.len {
			t.Errorf("Append8 with test_idx: %d; returned bad length %d, want: %d", test.idx, ba.Len(), test.len)
		}

		if !bytes.Equal(data, test.data) {
			t.Errorf("Append8 with test_idx: %d; returned bad data %s, want: %s", test.idx, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.data))
		}
	}

}

func TestAppend16(t *testing.T) {
	tests := []struct {
		idx     int
		ints    []uint16
		data    []byte
		padding byte
		len     int
	}{
		{0, []uint16{0xDEAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{1, []uint16{0xDE, 0xAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{2, []uint16{0xDE, 0xAD, 0xBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{3, []uint16{0xDEAD, 0x0B, 0x0E, 0x3B, 0x03}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{4, []uint16{0xDEAD, 0x0B, 0x0E, 0x3B03}, []byte("\xDE\xAD\xBE\xEC\x0C"), 2, 38},
		{5, []uint16{0x07}, []byte("\xE0"), 5, 3},
		{6, []uint16{0xC0, 0x3F, 0x03, 0xEE}, []byte("\xC0\xFF\xEE"), 0, 24},
		{7, []uint16{0xC03F, 0x03, 0xEE}, []byte("\xC0\x3F\xFB\x80"), 6, 26},
		{8, []uint16{0xC0, 0x3F, 0x03, 0xEE, 0xE}, []byte("\xC0\xFF\xEE\xE0"), 4, 28},
		{9, []uint16{0xDE, 0xAD, 0x07, 0x06, 0x05, 0x13, 0x02}, []byte("\xDE\xAD\xFA\xCE"), 0, 32},
		{10, []uint16{0xDEAD, 0x0706, 0x0513, 0x02}, []byte("\xDE\xAD\xE0\xD4\x4E"), 0, 40},
	}

	for _, test := range tests {
		ba := New()
		for _, e := range test.ints {
			ba.Append16(e)
		}

		data := ba.Bytes()

		if ba.padding != test.padding {
			t.Errorf("Append16 with test_idx: %d; returned bad padding %d, want: %d", test.idx, ba.padding, test.padding)
		}

		if ba.Len() != test.len {
			t.Errorf("Append16 with test_idx: %d; returned bad length %d, want: %d", test.idx, ba.Len(), test.len)
		}

		if !bytes.Equal(data, test.data) {
			t.Errorf("Append16 with test_idx: %d; returned bad data %s, want: %s", test.idx, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.data))
		}
	}

}

func TestAppend32(t *testing.T) {
	tests := []struct {
		idx     int
		ints    []uint32
		data    []byte
		padding byte
		len     int
	}{
		{0, []uint32{0xDEADBEEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{1, []uint32{0xDEAD, 0xBEEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{2, []uint32{0xDE, 0xADBE, 0xEF}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{3, []uint32{0xDEAD, 0x0B, 0x0E, 0x3B, 0x03}, []byte("\xDE\xAD\xBE\xEF"), 0, 32},
		{4, []uint32{0xDEAD0B0E, 0x3B, 0x03}, []byte("\xDE\xAD\x0B\x0E\xEF"), 0, 40},
		{5, []uint32{0x07}, []byte("\xE0"), 5, 3},
		{6, []uint32{0xC0, 0x3F, 0x03, 0xEE}, []byte("\xC0\xFF\xEE"), 0, 24},
		{7, []uint32{0xC03F03EE}, []byte("\xC0\x3F\x03\xEE"), 0, 32},
		{8, []uint32{0xC0, 0x3203, 0xEE, 0x0E}, []byte("\xC0\xC8\x0F\xBB\x80"), 6, 34},
		{9, []uint32{0xDEAD, 0x07, 0x06, 0x05, 0x13, 0x02}, []byte("\xDE\xAD\xFA\xCE"), 0, 32},
		{10, []uint32{0xDEAD, 0x07, 0x06, 0x05, 0x13, 0x02, 0x1F}, []byte("\xDE\xAD\xFA\xCE\xF8"), 3, 37},
		{11, []uint32{0x1B, 0xF}, []byte("\xDF\x80"), 7, 9},
	}

	for _, test := range tests {
		ba := New()
		for _, e := range test.ints {
			ba.Append32(e)
		}

		data := ba.Bytes()

		if ba.padding != test.padding {
			t.Errorf("Append32 with test_idx: %d; returned bad padding %d, want: %d", test.idx, ba.padding, test.padding)
		}

		if ba.Len() != test.len {
			t.Errorf("Append32 with test_idx: %d; returned bad length %d, want: %d", test.idx, ba.Len(), test.len)
		}

		if !bytes.Equal(data, test.data) {
			t.Errorf("Append32 with test_idx: %d; returned bad data %s, want: %s", test.idx, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.data))
		}
	}

}