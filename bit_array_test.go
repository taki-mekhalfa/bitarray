package bitarray

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAppendZeroAndOne(t *testing.T) {
	ba := New()
	for _, b := range "110111101010110110111110111011110000" {
		if b == '1' {
			ba.AppendOne()
		} else {
			ba.AppendZero()
		}
	}

	data := ba.Bytes()
	if !bytes.Equal(data, []byte("\xDE\xAD\xBE\xEF\x00")) {
		t.Errorf("AppendOne or AppendZero returned bad data %s, want %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", []byte("\xDE\xAD\xBE\xEF\x00")))
	}

	if ba.Len() != 36 {
		t.Errorf("AppendOne or AppendZero returned bad size %d, want %d", ba.Len(), 32)
	}

}

func TestAppendBit(t *testing.T) {
	ba := New()
	for _, b := range "110111101010110110111110111011110000" {
		if b == '1' {
			ba.AppendBit(1)
		} else {
			ba.AppendBit(0)
		}
	}

	data := ba.Bytes()
	if !bytes.Equal(data, []byte("\xDE\xAD\xBE\xEF\x00")) {
		t.Errorf("AppendOne or AppendZero returned bad data %s, want %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", []byte("\xDE\xAD\xBE\xEF\x00")))
	}

	if ba.Len() != 36 {
		t.Errorf("AppendOne or AppendZero returned bad size %d, want %d", ba.Len(), 32)
	}

}

func TestGetBit(t *testing.T) {
	ba := New()
	ba.data = []byte("\xDE\xAD\xBE\xEF")
	ba.padding = 0
	for i, bit := range "11011110101011011011111011101111" {
		var want byte
		if bit == '1' {
			want = 1
		}
		if ba.GetBit(i) != want {
			t.Errorf("GetBit(%d)=%d, want %d", i, ba.GetBit(i), want)
		}
	}
}

func TestSetBit(t *testing.T) {
	tests := []struct {
		position int
		want     []byte
	}{
		{0, []byte("\xDE\xAD\xBE\xEF")},
		{2, []byte("\xFE\xAD\xBE\xEF")},
		{5, []byte("\xDE\xAD\xBE\xEF")},
		{7, []byte("\xDF\xAD\xBE\xEF")},
		{11, []byte("\xDE\xBD\xBE\xEF")},
		{15, []byte("\xDE\xAD\xBE\xEF")},
		{17, []byte("\xDE\xAD\xFE\xEF")},
		{23, []byte("\xDE\xAD\xBF\xEF")},
		{24, []byte("\xDE\xAD\xBE\xEF")},
		{27, []byte("\xDE\xAD\xBE\xFF")},
		{31, []byte("\xDE\xAD\xBE\xEF")},
	}

	ba := New()
	ba.padding = 0

	for _, test := range tests {
		ba.data = []byte("\xDE\xAD\xBE\xEF")
		ba.SetBit(test.position)
		if !bytes.Equal(ba.data, test.want) {
			t.Errorf("SetBit(%d) did not work correctly, found %s, want %s", test.position, fmt.Sprintf("%#X", ba.data), fmt.Sprintf("%#X", test.want))
		}
	}
}

func TestClearBit(t *testing.T) {
	tests := []struct {
		position int
		want     []byte
	}{
		{0, []byte("\x5E\xAD\xBE\xEF")},
		{3, []byte("\xCE\xAD\xBE\xEF")},
		{6, []byte("\xDC\xAD\xBE\xEF")},
		{7, []byte("\xDE\xAD\xBE\xEF")},
		{8, []byte("\xDE\x2D\xBE\xEF")},
		{9, []byte("\xDE\xAD\xBE\xEF")},
		{16, []byte("\xDE\xAD\x3E\xEF")},
		{22, []byte("\xDE\xAD\xBC\xEF")},
		{24, []byte("\xDE\xAD\xBE\x6F")},
		{27, []byte("\xDE\xAD\xBE\xEF")},
		{31, []byte("\xDE\xAD\xBE\xEE")},
	}

	ba := New()
	ba.padding = 0

	for _, test := range tests {
		ba.data = []byte("\xDE\xAD\xBE\xEF")
		ba.ClearBit(test.position)
		if !bytes.Equal(ba.data, test.want) {
			t.Errorf("ClearBit(%d) did not work correctly, found %s, want %s", test.position, fmt.Sprintf("%#X", ba.data), fmt.Sprintf("%#X", test.want))
		}
	}
}

func TestAppend8(t *testing.T) {
	ba := New()
	ba.Append8(0xD, 4)
	ba.Append8(0xE, 4)
	ba.Append8(0xAD, 8)
	ba.Append8(0xAB, 0)
	ba.Append8(0xAB, 4)
	ba.Append8(0xEE, 8)
	ba.Append8(0x7, 3)
	ba.Append8(0xF, 1)
	ba.Append8(0xCD, 0)

	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xBE\xEF")
	expectedLen := 32

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("Append8 returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}
}

func TestAppend16(t *testing.T) {
	ba := New()
	ba.Append16(0xDE, 8)
	ba.Append16(0xADBE, 16)
	ba.Append16(0xAE, 4)
	ba.Append16(0xAB, 0)
	ba.Append16(0x0E, 4)
	ba.Append16(0xEE, 8)
	ba.Append16(0x7, 3)
	ba.Append16(0xF, 1)
	ba.Append16(0xCD0F, 4)
	ba.Append16(0xABAA, 12)
	ba.Append16(0xDDDD, 4)

	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xBE\xEE\xEE\xFF\xBA\xAD")
	expectedLen := 64

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("Append16 returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}
}

func TestAppend32(t *testing.T) {
	ba := New()
	ba.Append32(0xDE, 8)
	ba.Append32(0xADBEEF, 24)
	ba.Append32(0xCC, 4)
	ba.Append32(0xAB, 0)
	ba.Append32(0x10, 4)
	ba.Append32(0x00, 8)
	ba.Append32(0x7, 3)
	ba.Append32(0xF, 1)
	ba.Append32(0xA0CDFFEE, 16)
	ba.Append32(0xCBAAAA, 20)
	ba.Append32(0xDDDD, 4)

	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xBE\xEF\xC0\x00\xFF\xFE\xEB\xAA\xAA\xD0")
	expectedLen := 92

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("Append32 returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}
}

func TestAppend64(t *testing.T) {
	ba := New()
	ba.Append64(0xAABAAAAADF00000D, 56)
	ba.Append64(0xDE, 8)
	ba.Append64(0xADBEEF, 24)
	ba.Append64(0xCC, 4)
	ba.Append64(0xAB, 0)
	ba.Append64(0x10, 4)
	ba.Append64(0x00, 8)
	ba.Append64(0x7, 3)
	ba.Append64(0xF, 1)
	ba.Append64(0xA0CDFFEE, 16)
	ba.Append64(0xCBAAAA, 20)
	ba.Append64(0xDDDD, 4)

	data := ba.Bytes()
	expected := []byte("\xBA\xAA\xAA\xDF\x00\x00\x0D\xDE\xAD\xBE\xEF\xC0\x00\xFF\xFE\xEB\xAA\xAA\xD0")
	expectedLen := 148

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("Append64 returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}
}

func TestAppendBytes(t *testing.T) {
	ba := New()
	ba.AppendBytes([]byte("\xDE\xAD"), 0)
	ba.AppendBytes([]byte("\xC0\xFF\xEE\xEF"), 4)

	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xC0\xFF\xEE\xE0")
	expectedLen := 44

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("AppendBytes returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}

}

func TestAppendBitArray(t *testing.T) {
	ba := New()
	ba.Append(0xDEAD, 16)
	ba1 := New()
	ba1.Append(0xBEEF, 16)

	ba.AppendBitArray(ba1)

	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xBE\xEF")
	expectedLen := 32

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("AppendBitArray returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}

}

func TestAppendFromString(t *testing.T) {
	ba := New()
	bitSeq := "110111101010110110111110111011110000000000001011101010101010110111000000000000001111111111101110"
	ba.AppendFromString(bitSeq)
	data := ba.Bytes()
	expected := []byte("\xDE\xAD\xBE\xEF\x00\x0B\xAA\xAD\xC0\x00\xFF\xEE")
	expectedLen := len(bitSeq)

	if !bytes.Equal(ba.data, expected) {
		t.Errorf("AppendFromString returned bad data %s, want: %s", fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", expected))
	}

	if ba.Len() != expectedLen {
		t.Errorf("returned bad length %d, want: %d", ba.Len(), expectedLen)
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		idx  int
		data string
		i    int
		j    int
		want uint64
	}{
		{0, "00101010", 2, 5, 5},
		{1, "00101010", 6, 7, 1},
		{2, "00101010", 7, 8, 0},
		{3, "00101010", 0, 8, 42},
		{4, "00101010", 2, 7, 21},
		{5, "00010101100100101010110000001111111111101110101010101001001010", 20, 44, 0xC0FFEE},
		{6, "00010101100000000000110000001111111111101110101010101001001010", 9, 44, 0xC0FFEE},
		{7, "10010111001001110001110110110101100111010100101011111111000100", 15, 62, 0x476D6752BFC4},
		{8, "1000000101011110110010110011110000001110000011111100001011011011", 0, 64, 0x815ECB3C0E0FC2DB},
		{9, "00010010100011110001010101001100001101010010000000101101010", 13, 55, 0x38AA61A9016},
		{10, "00111111101101010000001101111001110", 0, 21, 0x7F6A0},
		{11, "1011000001011111100101011010111110111101011", 5, 41, 0xBF2B5F7A},
		{12, "111110010000111010001110110100010000110011001000101110001011110100001000000000111", 11, 75, 0x7476886645C5E840},
		{13, "100100100011110101101110100001001010001010110010101000000101101111111101111001111", 0, 64, 0x923D6E84A2B2A05B},
		{14, "111010110000011001110111100101110110001110111011111101000011000110101001011110010", 15, 30, 0x1DE5},
		{15, "110100110010111100100011011001101111000101111011101110010110000110100101110110010", 18, 20, 0x2},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendFromString(test.data)

		v := ba.Extract(test.i, test.j)
		if v != test.want {
			t.Errorf("Range, idx=%d; bad value %d, want %d", test.idx, v, test.want)
		}

	}
}
