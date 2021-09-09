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
