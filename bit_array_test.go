package bitarray

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAppendZeroAndOne(t *testing.T) {
	appendZerosAndOnes := func(ba *BitArray, bitSeq string) {
		for _, b := range bitSeq {
			if b == '1' {
				ba.AppendOne()
			} else {
				ba.AppendZero()
			}
		}
	}

	tests := []struct {
		id     int
		bitSeq string
		want   []byte
		len    int
	}{
		{0, "110111101010110110111110111011110000", []byte("\xDE\xAD\xBE\xEF\x00"), 36},
		{1, "0001000000101011001101001001010110100000011100100111110101010111", []byte("\x10+4\x95\xa0r}W"), 64},
		{2, "10000101001110101110001100000011111100011011001111110101", []byte("\x85:\xe3\x03\xf1\xb3\xf5"), 56},
		{3, "1010101001111110100101001010000111101011", []byte("\xaa~\x94\xa1\xeb"), 40},
		{4, "010001011000001111101111011010111100001010101111011000011010001000110001001010010111", []byte("E\x83\xefk\xc2\xafa\xa21)\x70"), 84},
	}

	for _, test := range tests {
		ba := New()
		appendZerosAndOnes(ba, test.bitSeq)
		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: AppendOne or AppendZero returned bad data %s, want %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: AppendOne or AppendZero returned bad size %d, want %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestAppendBit(t *testing.T) {
	appendBits := func(ba *BitArray, bitSeq string) {
		for _, b := range bitSeq {
			if b == '1' {
				ba.AppendBit(1)
			} else {
				ba.AppendBit(0)
			}
		}
	}

	tests := []struct {
		id     int
		bitSeq string
		want   []byte
		len    int
	}{
		{0, "110111101010110110111110111011110000", []byte("\xDE\xAD\xBE\xEF\x00"), 36},
		{1, "0001000000101011001101001001010110100000011100100111110101010111", []byte("\x10+4\x95\xa0r}W"), 64},
		{2, "10000101001110101110001100000011111100011011001111110101", []byte("\x85:\xe3\x03\xf1\xb3\xf5"), 56},
		{3, "1010101001111110100101001010000111101011", []byte("\xaa~\x94\xa1\xeb"), 40},
		{4, "010001011000001111101111011010111100001010101111011000011010001000110001001010010111", []byte("E\x83\xefk\xc2\xafa\xa21)\x70"), 84},
	}

	for _, test := range tests {
		ba := New()
		appendBits(ba, test.bitSeq)
		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: AppendBit returned bad data %s, want %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: AppendBit returned bad size %d, want %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestGetBit(t *testing.T) {
	tests := []struct {
		id       int
		data     []byte
		indicies []int
		want     []byte
	}{
		{
			0,
			[]byte("\x91\xb1\xa9\x8d\x95Ni"),
			[]int{52, 5, 25, 33, 1},
			[]byte{1, 0, 0, 0, 0},
		},
		{
			1,
			[]byte("\xa9 \xb5C\xe2\x0b\x7ff|D\x0f"),
			[]int{41, 16, 60, 5, 83, 58, 36, 54, 53, 81, 38, 72, 73, 47, 76},
			[]byte{0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0},
		},
		{
			2,
			[]byte("\xcaR\x16\xf7\x7f\x1e\xedf\xf5"),
			[]int{3, 7, 25, 9, 8, 48, 54, 18, 17, 62},
			[]byte{0, 0, 1, 1, 0, 1, 0, 0, 0, 1},
		},
		{
			3,
			[]byte("\xc8\x97\xe6r\xb7\xe1\xb3\x05\x03W"),
			[]int{26, 66, 52, 79, 22, 47, 4, 42, 0, 76},
			[]byte{1, 0, 0, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			4,
			[]byte("OZnRy\xd8\x15\x07\x1c]\x0flF7\xb0x\xef\x17\x80IX\xcd\x97\xb7"),
			[]int{55, 13, 43, 135, 187, 69, 137, 173, 151, 1, 109, 77, 15, 34, 97, 185, 81, 125, 87, 129, 6, 161, 22, 60, 120, 36, 82, 61, 21, 10, 95, 189},
			[]byte{1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1},
		},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendBytes(test.data, 0)
		result := make([]byte, len(test.want))

		for i, e := range test.indicies {
			result[i] = ba.GetBit(e)
		}

		if !bytes.Equal(result, test.want) {
			t.Errorf("%d: GetBit returned bad data %v; want %v", test.id, result, test.want)
		}
	}
}

func TestClearBit(t *testing.T) {
	tests := []struct {
		id       int
		data     []byte
		indicies []int
		want     []byte
	}{
		{
			0,
			[]byte("\x1e[>\x83\x9e q"),
			[]int{4, 40, 16, 10, 17, 52, 49},
			[]byte("\x16\x5B\x3E\x83\x9E\x20\x31"),
		},
		{
			1,
			[]byte("\xb5\xac\xe4\xd3k0\xbb\xdbY\x0e\xf4"),
			[]int{38, 71, 26, 87, 59, 40, 19, 17, 15, 23, 86, 5, 7},
			[]byte("\xB0\xAC\xA4\xD3\x69\x30\xBB\xCB\x58\x0E\xF4"),
		},
		{
			2,
			[]byte("\xb7\xff\xfc\x97\x07\x8f\xad\xfa/D\x976\xd3\xcci"),
			[]int{118, 106, 25, 112, 2, 109, 103, 7, 58, 107, 90, 86, 60, 23, 54, 12, 3, 87, 117, 73},
			[]byte("\x86\xF7\xFC\x97\x07\x8F\xAD\xD2\x2F\x04\x94\x16\xD2\xC8\x69"),
		},
		{
			3,
			[]byte("$4\x85\xc4\xd0\xd7\xe6\x9a\x18"),
			[]int{17, 67, 14},
			[]byte("\x24\x34\x85\xC4\xD0\xD7\xE6\x9A\x08"),
		},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendBytes(test.data, 0)
		for _, i := range test.indicies {
			ba.ClearBit(i)
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: ClearBit did not work correctly, found %s, want %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}
	}
}

func TestSetBit(t *testing.T) {
	tests := []struct {
		id       int
		data     []byte
		indicies []int
		want     []byte
	}{
		{
			0,
			[]byte("M\xce'Nw\x1ek\x12\x90"),
			[]int{50, 64, 70, 41, 18, 16, 38},
			[]byte("\x4D\xCE\xA7\x4E\x77\x5E\x6B\x12\x92"),
		},
		{
			1,
			[]byte("\xc5\xbe\x9e\x18}\x10\x0f\xd7\x8f\xc5\xc8\xf2\xcf\xf8\x10"),
			[]int{23, 50, 116, 106, 72, 11, 80, 61, 73, 107, 51, 21, 28, 25, 118, 4, 79, 38, 105, 62, 44, 67, 74},
			[]byte("\xCD\xBE\x9F\x58\x7F\x18\x3F\xD7\x9F\xE5\xC8\xF2\xCF\xF8\x1A"),
		},
		{
			3,
			[]byte("\x15\xc9r\x7fu\x12\xf0\\!!\xe9Lyk\x81E\xcd*2\x99"),
			[]int{95, 51, 128, 66, 154, 122, 28, 25, 151, 134, 153, 94, 6, 39, 125, 65, 73},
			[]byte("\x17\xC9\x72\x7F\x75\x12\xF0\x5C\x61\x61\xE9\x4F\x79\x6B\x81\x65\xCF\x2A\x33\xF9"),
		},
		{
			4,
			[]byte("\xda\xdb\xe5\x80\x7f\xc6\xb0z\xa4\xbe"),
			[]int{66, 4, 16, 73, 13},
			[]byte("\xDA\xDF\xE5\x80\x7F\xC6\xB0\x7A\xA4\xFE"),
		},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendBytes(test.data, 0)
		for _, i := range test.indicies {
			ba.SetBit(i)
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: SetBit did not work correctly, found %s, want %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}
	}
}

func TestAppend8(t *testing.T) {
	tests := []struct {
		id    int
		data  []byte
		sizes []int
		want  []byte
		len   int
	}{
		{
			0,
			[]byte{0xD, 0xE, 0xAD, 0xAB, 0xAB, 0xEE, 0x7, 0xF, 0xCD},
			[]int{4, 4, 8, 0, 4, 8, 3, 1, 0},
			[]byte("\xDE\xAD\xBE\xEF"),
			32,
		},
		{
			1,
			[]byte("\xc9\xef\x14\x14\xdd<gC\x8a\x10"),
			[]int{5, 7, 5, 3, 7, 4, 6, 6, 3, 7},
			[]byte("N\xfaK\xb98h\x80"),
			53,
		},
		{
			2,
			[]byte("M\r\xda\xf1\xd8\x95\x12\xa6\xfb\xee\xc0\x15\xc6\x85w%\xcf\x9d9?"),
			[]int{4, 7, 5, 7, 0, 0, 3, 0, 3, 3, 1, 7, 3, 3, 5, 7, 1, 7, 1, 7},
			[]byte("\xd1\xba\xe2\x9e\x15\xd6\xe9go\xc0"),
			74,
		},
		{
			3,
			[]byte("\x16\x91\xac\xf1\xedy\x14$O\x82\xa1\x9a\xfe#\x1d\xd1\xd8,\xe8\xc7\xc28*w\xa6`\xcf\xf0t\x1c\xab\xe8\xb8\xb7\xb2\x8d"),
			[]int{6, 0, 1, 3, 1, 3, 2, 2, 5, 5, 5, 6, 7, 4, 5, 5, 0, 6, 0, 6, 1, 3, 3, 6, 7, 5, 6, 1, 0, 6, 6, 4, 6, 3, 1, 2},
			[]byte("Xd\x1e \xb5\xf8\xfb\x1b\x07\x05\xba`\x1er\xb8\xe3\x90"),
			132,
		},
		{
			4,
			[]byte("o\xfaQ]J\xc1=\x87\xf6K\xc65\x9c\x92\x96\x10\xa9"),
			[]int{7, 6, 4, 7, 0, 7, 3, 2, 1, 7, 3, 3, 2, 6, 1, 1, 5},
			[]byte("\xdf\xd0\xdd\x83t\xbdD\x84\x80"),
			65,
		},
		{
			5,
			[]byte("\x0D\x00\x0D\xBA\x0D\xC0\x0F\xFE\xEE"),
			[]int{4, 8, 4, 8, 4, 8, 4, 8, 4},
			[]byte("\xD0\x0D\xBA\xDC\x0F\xFE\xE0"),
			52,
		},
	}

	for _, test := range tests {
		ba := New()
		for i, b := range test.data {
			ba.Append8(b, test.sizes[i])
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: Append8 returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestAppend16(t *testing.T) {
	tests := []struct {
		id    int
		data  []uint16
		sizes []int
		want  []byte
		len   int
	}{
		{
			0,
			[]uint16{0xDE, 0xADBE, 0xAE, 0xAB, 0x0E, 0xEE, 0x7, 0xF, 0xCD0F, 0xABAA, 0xDDDD},
			[]int{8, 16, 4, 0, 4, 8, 3, 1, 4, 12, 4},
			[]byte("\xDE\xAD\xBE\xEE\xEE\xFF\xBA\xAD"),
			64,
		},
		{
			1,
			[]uint16{0xE1F1, 0xA806, 0xD6CF, 0x4EB7, 0xC968, 0x5DB0, 0xA2BE},
			[]int{5, 9, 7, 7, 12, 4, 15},
			[]byte("\x88\x1a{yh\x04W\xc0"),
			59,
		},
		{
			2,
			[]uint16{0x230, 0x1059, 0x3410, 0xD291, 0xD6E2, 0xAA16, 0xA254, 0xE7DF, 0x1E87, 0x29B9, 0x32DF, 0xFAFC, 0xB0EC, 0xFAC8, 0x3644, 0xA25A, 0x2FC3},
			[]int{6, 1, 2, 7, 6, 11, 11, 5, 8, 9, 2, 8, 13, 2, 6, 10, 11},
			[]byte("\xc2\x11\x89\x0b%O\xc3\xee\x7f\xc8v\x02K_\x0c"),
			118,
		},
		{
			3,
			[]uint16{0x4FBB, 0xA25D, 0x4DE5, 0x544D, 0xF771, 0x3AB6, 0x6646, 0xC59F, 0x3E23, 0x1CDF, 0xA54F, 0xDB6B, 0x77BF, 0x44BD, 0x35DC, 0xA56B},
			[]int{1, 11, 13, 15, 2, 3, 12, 5, 7, 0, 1, 8, 7, 2, 7, 8},
			[]byte("\xa5\xd6\xf2\xd4Ms#}\x1d\xad\xfbq\xac"),
			102,
		},
		{
			4,
			[]uint16{0x2D3E, 0x6E59, 0xE8BD, 0x847B, 0xABCF, 0xF0D1},
			[]int{4, 9, 2, 16, 8, 4},
			[]byte("\xe2\xcb\x08\xf7\x9e "),
			43,
		},
		{
			5,
			[]uint16{0x724, 0x366D, 0x633C, 0x92CA, 0x45C3, 0x3C8B, 0xF085, 0xF07D, 0x2F6F, 0x4E0E, 0x8E20, 0xF22C, 0xFC35, 0x5958, 0x89BE, 0xB37C, 0xE15F, 0x215D, 0x787A, 0xA444, 0x2726, 0x6988, 0xAB97, 0x5EAC, 0xEAD4, 0xFCB, 0x593F, 0x81C6, 0x7311, 0xA437, 0xD621, 0xC891},
			[]int{3, 0, 10, 14, 0, 7, 16, 5, 9, 5, 7, 13, 6, 2, 1, 9, 8, 16, 6, 12, 12, 9, 1, 6, 1, 16, 0, 12, 6, 12, 14, 14},
			[]byte("\x99\xe2YB\xfc!{or\t\x16j/\x8b\xe4+\xbd\"#\x93b6\x03\xf2\xc7\x19\x147X\x84\x89\x10"),
			252,
		},
	}

	for _, test := range tests {
		ba := New()
		for i, b := range test.data {
			ba.Append16(b, test.sizes[i])
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: Append16 returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba.Len(), test.len)
		}
	}

}

func TestAppend32(t *testing.T) {
	tests := []struct {
		id    int
		data  []uint32
		sizes []int
		want  []byte
		len   int
	}{
		{
			0,
			[]uint32{0xDE, 0xADBEEF, 0xCC, 0xAB, 0x10, 0x00, 0x7, 0xF, 0xA0CDFFEE, 0xCBAAAA, 0xDDDD},
			[]int{8, 24, 4, 0, 4, 8, 3, 1, 16, 20, 4},
			[]byte("\xDE\xAD\xBE\xEF\xC0\x00\xFF\xFE\xEB\xAA\xAA\xD0"),
			92,
		},
		{
			1,
			[]uint32{0x5BE28C56, 0x719991E, 0xD4960426, 0x9F91B7E6},
			[]int{32, 0, 10, 18},
			[]byte("[\xe2\x8cV\t\x9b~`"),
			60,
		},
		{
			2,
			[]uint32{0xEE1439E9, 0xD516EAEA, 0x4A43BA9F, 0x4151A8EA, 0x3D4447CE, 0xBF6953C3, 0x38C8664E, 0x8171CA, 0x2CBB7C3B, 0x837B547A, 0xAE50C237, 0x165E89F9, 0x3F867936, 0x40F6B546, 0x53B8B7FB, 0x9FDC04BA},
			[]int{22, 23, 2, 3, 4, 12, 7, 9, 29, 12, 32, 17, 5, 15, 0, 20},
			[]byte("P\xe7\xa4\xb7WV\xb8\xf0\xe7r\x99v\xf8v\x8fU\xca\x18F\xe8\x9f\x9b5F\xc0K\xa0"),
			212,
		},
		{
			3,
			[]uint32{0xC4AB74AC, 0x6E029B6E, 0x71388525, 0x8E16E2B2, 0xDE703AD0},
			[]int{29, 24, 9, 15, 20},
			[]byte("%[\xa5`\x14\xdbt\x97\x15\x90\x1dh\x00"),
			97,
		},
		{
			4,
			[]uint32{0xB525410F, 0xBC1C527F, 0xBA621264, 0x375D76C6, 0x52C698AC, 0xD841A98A, 0xD039471E, 0xF7905947, 0x209320CB, 0x586118A5, 0xB6A3AB93},
			[]int{32, 19, 16, 6, 27, 32, 16, 1, 4, 26, 4},
			[]byte("\xb5%A\x0f\x8aO\xe2L\x83,i\x8a\xcd\x84\x1a\x98\xa4q\xed\x8c#\x14\xa6"),
			183,
		},
		{
			5,
			[]uint32{0x66977500, 0x615C90F0, 0x4C7455F4, 0xBF2EFE40, 0x513B4605, 0x5FF8A923, 0xB4C6890E, 0x782BB3C, 0xF42CB292},
			[]int{32, 17, 29, 0, 9, 6, 26, 0, 21},
			[]byte("f\x97u\x00Hx1\xd1W\xd0\x0b\x19\x8d\x12\x1c\xcb) "),
			140,
		},
	}

	for _, test := range tests {
		ba := New()
		for i, b := range test.data {
			ba.Append32(b, test.sizes[i])
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: Append32 returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestAppend64(t *testing.T) {
	tests := []struct {
		id    int
		data  []uint64
		sizes []int
		want  []byte
		len   int
	}{
		{
			0,
			[]uint64{0xAABAAAAADF00000D, 0xDE, 0xADBEEF, 0xCC, 0xAB, 0x10, 0x00, 0x7, 0xF, 0xA0CDFFEE, 0xCBAAAA, 0xDDDD},
			[]int{56, 8, 24, 4, 0, 4, 8, 3, 1, 16, 20, 4},
			[]byte("\xBA\xAA\xAA\xDF\x00\x00\x0D\xDE\xAD\xBE\xEF\xC0\x00\xFF\xFE\xEB\xAA\xAA\xD0"),
			148,
		},
		{
			1,
			[]uint64{0x4C8D31768D7DB96C, 0x86699492B0358984, 0x872C66EF42A6CE8D, 0x21441339071CCB37},
			[]int{31, 17, 4, 58},
			[]byte("\x1a\xfbr\xd9\x89\x84\xd5\x10L\xe4\x1cs,\xdc"),
			110,
		},
		{
			2,
			[]uint64{0xCE74C2F3ACE1DE18, 0x19C85B6548D353DE, 0xBF2C05B21073D023, 0x9A3E8E7A3C54EDAA, 0x3D3EF83EC6C3E659},
			[]int{55, 36, 34, 51, 40},
			[]byte("\xe9\x85\xe7Y\xc3\xbc0\xa9\x1aj{\xd0\x83\x9e\x81\x1e\x8ez<T\xed\xaa>\xc6\xc3\xe6Y"),
			216,
		},
		{
			3,
			[]uint64{0xDEDDEFEBA552D14D, 0xE33E15FED14AB3D6, 0x3DEE98E9BD47ECF1, 0xD8B74BE1DBE30CEE, 0x51B28786915C9F44, 0x7980776047F9D705, 0xCB162247C2FDD3B7, 0x1DF1B3EA9F674A21},
			[]int{33, 37, 34, 52, 24, 56, 10, 19},
			[]byte("\xd2\xa9h\xa6\xfbE*\xcfY\xbdG\xec\xf1t\xbe\x1d\xbe0\xce\xe5\xc9\xf4H\x07v\x04\x7f\x9dp^\xdf\xa5\x10\x80"),
			265,
		},
		{
			4,
			[]uint64{0xF7ECD88FC2ED1523, 0xDEB10B6BFF0B5F0F, 0xA9B2343E07B6B632, 0xC84AEEF94193FC1C, 0x4790CAEE9278F65D, 0xCE6229292F5B77C4},
			[]int{36, 38, 16, 61, 13, 43},
			[]byte("\xfc.\xd1R:\xff\xc2\xd7\xc3\xed\x8c\x90\x95\xdd\xf2\x83'\xf89e\xd2R^\xb6\xef\x88"),
			207,
		},
		{
			5,
			[]uint64{0x4537E0AC293CC4E1, 0x74C2812105763166, 0xF499084C2CDDD072, 0xBCBF439F27CE52B1, 0xCF7D73EFE6C9E451, 0xC8D324A9D6122DED},
			[]int{15, 64, 4, 29, 64, 43},
			[]byte("\x89\xc2\xe9\x85\x02B\n\xecb\xccG\xceR\xb1\xcf}s\xef\xe6\xc9\xe4Q\x95:\xc2E\xbd\xa0"),
			219,
		},
	}
	for _, test := range tests {
		ba := New()
		for i, b := range test.data {
			ba.Append64(b, test.sizes[i])
		}

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: Append64 returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestAppendBytes(t *testing.T) {
	tests := []struct {
		id      int
		input   []byte
		padding int
		want    []byte
		len     int
	}{
		{
			0,
			[]byte("\xcam\x1d\x88%\x83DQP\x11\xa7\xbb\x9c"),
			0,
			[]byte("\xcam\x1d\x88%\x83DQP\x11\xa7\xbb\x9c"),
			104,
		},
		{
			1,
			[]byte("\xdf\x00y\xbd)\xf5\xaf\x98\x07r\x03\xb0*\xaea"),
			1,
			[]byte("\xdf\x00y\xbd)\xf5\xaf\x98\x07r\x03\xb0*\xae`"),
			119,
		},
		{
			2,
			[]byte("\x94=\xfb[\xcf\xce\xdc"),
			2,
			[]byte("\x94=\xfb[\xcf\xce\xdc"),
			54,
		},
		{
			3,
			[]byte("\xc5\x89x|\x85xKj\xea)R\xbe\x9f\x9d}"),
			3,
			[]byte("\xc5\x89x|\x85xKj\xea)R\xbe\x9f\x9dx"),
			117,
		},
		{
			4,
			[]byte("J$H%\xac\n\xb9\"_\xa7O(h"),
			4,
			[]byte("J$H%\xac\n\xb9\"_\xa7O(`"),
			100,
		},
		{
			5,
			[]byte("\xfd\x01\x81g89&\xdcw\xd0\xbe\x0cx\xaet\x0f\xf3\xe2\x10\\\xcf"),
			5,
			[]byte("\xfd\x01\x81g89&\xdcw\xd0\xbe\x0cx\xaet\x0f\xf3\xe2\x10\\\xc0"),
			163,
		},
		{
			6,
			[]byte("\x07\xe3\xc9'\xa6\xd8U\x1b\xe9\xfaLl&\x95\x80\x8e"),
			6,
			[]byte("\x07\xe3\xc9'\xa6\xd8U\x1b\xe9\xfaLl&\x95\x80\x80"),
			122,
		},
		{
			7,
			[]byte("\xab\xf4^\xba.\xd0\xd9\xd2\xc0\xd80e8\xcc\x8f\xe0\x05\xdd\xec\x12\xeb\xcen\xc7\xe5\x02\x1e\xf3\xd8\xfb*\x8e"),
			7,
			[]byte("\xab\xf4^\xba.\xd0\xd9\xd2\xc0\xd80e8\xcc\x8f\xe0\x05\xdd\xec\x12\xeb\xcen\xc7\xe5\x02\x1e\xf3\xd8\xfb*\x80"),
			249,
		},
		{
			8,
			[]byte{},
			0,
			[]byte{},
			0,
		},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendBytes(test.input, test.padding)

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: AppendBytes returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", ba.Len(), test.id, test.len)
		}
	}
}

func TestAppendBitArray(t *testing.T) {
	tests := []struct {
		id   int
		a    string
		b    string
		want []byte
		len  int
	}{
		{
			0,
			"1101111010101101",
			"1011111011101111",
			[]byte("\xDE\xAD\xBE\xEF"),
			32,
		},
		{
			1,
			"1101101011111011110000101111101001110010101000101101100010011001100010100000011101110110000010001101111110111011001010010001110",
			"100110011110010000011011000111000010101001011101000110010101101111010000110111100011001011011110111011101011010011110110110010",
			[]byte("\xda\xfb\xc2\xfar\xa2\xd8\x99\x8a\x07v\x08\xdf\xbb)\x1d3\xc868T\xba2\xb7\xa1\xbce\xbd\xddi\xed\x90"),
			253,
		},
		{
			2,
			"1010111110111011100000001111100011110010110111101111011000101000",
			"110000001010101010001111000110001000000100100000111110100001",
			[]byte("\xaf\xbb\x80\xf8\xf2\xde\xf6(\xc0\xaa\x8f\x18\x81 \xfa\x10"),
			124,
		},
		{
			3,
			"1010010001000100011001101001111001000101001010011111000110000100",
			"110100101011010000101100101011001100110001011100110011001111001",
			[]byte("\xa4Df\x9eE)\xf1\x84\xd2\xb4,\xac\xcc\\\xcc\xf2"),
			127,
		},
		{
			4,
			"1101100100101110100110100111011000111100000001101010101111101101",
			"101010000011110101011001111110010100001111011011111000010001011",
			[]byte("\xd9.\x9av<\x06\xab\xed\xa8=Y\xf9C\xdb\xe1\x16"),
			127,
		},
		{
			5,
			"1101001010011101011011011100001110111010100010010001001010110000",
			"110000010110",
			[]byte("\xd2\x9dm\xc3\xba\x89\x12\xb0\xc1`"),
			76,
		},
	}
	for _, test := range tests {
		ba1 := New()
		ba1.AppendString(test.a)

		ba2 := New()
		ba2.AppendString(test.b)

		ba1.AppendBitArray(ba2)
		data := ba1.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: AppendBitArray returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba1.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba1.Len(), test.len)
		}
	}
}

func TestAppendString(t *testing.T) {
	tests := []struct {
		id     int
		bitSeq string
		want   []byte
		len    int
	}{
		{
			0,
			"1101111010101101",
			[]byte("\xDE\xAD"),
			16,
		},
		{
			1,
			"111111000000001111111000001101111000111010011010000101010111000110011010111100000001000111111110101010100001010011101011111100",
			[]byte("\xfc\x03\xf87\x8e\x9a\x15q\x9a\xf0\x11\xfe\xaa\x14\xeb\xf0"),
			126,
		},
		{
			2,
			"11101010110110100010101",
			[]byte("\xea\xda*"),
			23,
		},
		{
			3,
			"11111110010101100101001101000111100011010100101001010101",
			[]byte("\xfeVSG\x8dJU"),
			56,
		},
		{
			4,
			"110110011110001011000101000100010101100111110110000100011111000110010111000000010011100110010111100110101010111101011000010110101101010010000001010110110010001100101101100111001011100100000100010011001100001100111010110100001000110110111011010000100101100",
			[]byte("\xd9\xe2\xc5\x11Y\xf6\x11\xf1\x97\x019\x97\x9a\xafXZ\xd4\x81[#-\x9c\xb9\x04L\xc3:\xd0\x8d\xbbBX"),
			255,
		},
		{
			5,
			"11010001110011000001010011101101011101101101111001101011111001100100110101111111011010001011000101011",
			[]byte("\xd1\xcc\x14\xedv\xdek\xe6M\x7fh\xb1X"),
			101,
		},
	}
	for _, test := range tests {
		ba := New()
		ba.AppendString(test.bitSeq)

		data := ba.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: AppendString returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if ba.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, ba.Len(), test.len)
		}
	}
}

func TestExtract(t *testing.T) {
	tests := []struct {
		id   int
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
		ba.AppendString(test.data)

		v := ba.Extract(test.i, test.j)
		if v != test.want {
			t.Errorf("Range, idx=%d; bad value %d, want %d", test.id, v, test.want)
		}
	}
}

func TestExtractBitArray(t *testing.T) {
	tests := []struct {
		id   int
		data []byte
		i    int
		j    int
		want []byte
		len  int
	}{
		{
			0,
			[]byte{0b00101010},
			2,
			5,
			[]byte{0b10100000},
			3,
		},
		{
			1,
			[]byte{0b00101010},
			6,
			7,
			[]byte{0b10000000},
			1,
		},
		{
			2,
			[]byte{0b00101010},
			7,
			8,
			[]byte{0b00000000},
			1,
		},
		{
			3,
			[]byte{0b00101010},
			0,
			8,
			[]byte{0b00101010},
			8,
		},
		{
			4,
			[]byte{0b00101010},
			2,
			7,
			[]byte{0b10101000},
			5,
		},
		{
			5,
			[]byte{0b00101010},
			4,
			4,
			[]byte{},
			0,
		},
		{
			6,
			[]byte{0b00101010},
			0,
			0,
			[]byte{},
			0,
		},
		{
			7,
			[]byte("\xb7<\xeeK\xce"),
			0,
			40,
			[]byte("\xb7<\xeeK\xce"),
			40,
		},
		{
			8,
			[]byte("/Y\x13\xd2\x13\xf5\x92"),
			13,
			47,
			[]byte("\"zB~\x80"),
			34,
		},
		{
			9,
			[]byte("W\\\n\x12\x0b\xb1\x0e\xf6\x1b\xab7\n\xf5\xb6P"),
			13,
			86,
			[]byte("\x81BAv!\xde\xc3uf\x80"),
			73,
		},
		{
			10,
			[]byte("wu@M]Ue\xe9\x0f|\x14\x1b<\xec\xb5\xaf\n\xc1\x0eU-\x0cJ"),
			34,
			167,
			[]byte("uU\x97\xa4=\xf0Pl\xf3\xb2\xd6\xbc+\x049T\xb0"),
			133,
		},
		{
			11,
			[]byte("wu@M]Ue\xe9\x0f|\x14\x1b<\xec\xb5\xaf\n\xc1\x0eU-\x0cJ"),
			167,
			167,
			[]byte{},
			0,
		},
		{
			12,
			[]byte("E\xfc\t\xb2e\x80q\xfd/\x92d\xdax5\xd2\xbf\x0f\x89\x84(y\xea\xdd\xbcz\xf5\xf2\xb8h\x80,\xef\xc2\x06\xc1"),
			103,
			199,
			[]byte("\x1a\xe9_\x87\xc4\xc2\x14<\xf5n\xde="),
			96,
		},
		{
			13,
			[]byte("\xfc\t\xb2e\x80q\xfd/\x92d\xdax5\xd2\xbf\x0f\x89\x84(y\xea\xdd\xbcz\xf5\xf2\xb8h\x80,\xef\xc2\x06\xc1"),
			0,
			272,
			[]byte("\xfc\t\xb2e\x80q\xfd/\x92d\xdax5\xd2\xbf\x0f\x89\x84(y\xea\xdd\xbcz\xf5\xf2\xb8h\x80,\xef\xc2\x06\xc1"),
			272,
		},
		{
			14,
			[]byte("\xaeq4;!\xc2\x87r\xceP\n\x0bC\x85FA"),
			56,
			57,
			[]byte("\x00"),
			1,
		},
		{
			15,
			[]byte("\xaeq4;!\xc2\x87r\xceP\n\x0bC\x85FA"),
			57,
			58,
			[]byte("\x80"),
			1,
		},
	}

	for _, test := range tests {
		ba := New()
		ba.AppendBytes(test.data, 0)

		res := ba.ExtractBitArray(test.i, test.j)
		data := res.Bytes()
		if !bytes.Equal(data, test.want) {
			t.Errorf("%d: ExtractBitArray returned bad data %s, want: %s", test.id, fmt.Sprintf("%#X", data), fmt.Sprintf("%#X", test.want))
		}

		if res.Len() != test.len {
			t.Errorf("%d: returned bad length %d, want: %d", test.id, res.Len(), test.len)
		}
	}
}
