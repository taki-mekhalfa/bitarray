# bitarray

Bitarray implements a slice like data structure for bits. The implementation is memory efficient as bits are stored in actual memory bits and
eight bits are represented by one byte in a contiguous block of memory. 

Key features:
* Querying:
	+ `Len()`  returns the length of the bit array (the number of stored bits)
	+ `Bytes()`  returns a zero padded slice of bytes representing the bits contiguously with at most `7` zero padding bits.
	+ `Padding()`  returns the number of padding bits in the slice of bytes returned by `Bytes()`
	+ `GetBit(i)`  returns the bit at the `i`th position as a byte value which is either equal to `00000000` or `00000001`. Indexing start from `0`
	+ `Extract(i,j)`  returns the bits in the range `[i,j]` (`i`th included, `j`th bit excluded ) as a `uint64`. Bits in the range are stored to the left of the returned `uint64` (bit at last position (`j-1`) is stored at the LSB). This is the recommended method if the number of queried bits fits in a `uint64`
	+ `ExtractBitArray(i,j)` method returns another bit array representing the bits in the range `[i,j]` (`i`th included, `j`th bit excluded )
* Changing:
	+ `AppendOne()` or `AppendZero()` appends a `0` or `1` bit to the end of the bit array
	+ `AppendBit(bit)` appends bits `0` or `1` depending on the value of `bit` which is a byte equal to `00000000` or `00000001`
	+ `SetBit(i)` or `ClearBit(i)` sets or clears the bit at the `i`th position
	+ `Append8(v, nbBits)`, `Append16(v, nbBits)`, `Append32(v, nbBits)`, `Append64(v, nbBits)`, `Append(v, nbBits)` appends a `uint8`, `uint16`, `uint32`, `uint64` or a generic `uint` to the bit array and specify the number of bits to append. Bits are appended starting from the least significant bit (LSB) and going to the left for specified number of bits
	+ `AppendBytes(bytes, padding)` appends contiguous bits stores in a slice of bytes. User should specify the number of padding pits: `0` (no padding bits) to `7` (only one bit is used) in the last byte in argument data
	+ `AppendBitArray(ba)` appends the argument bit array to the receiving one
	+ `AppendString(bits)` appends a string sequence of `"0"`s and `"1"`s to the bit array

## Usage
The following shows some examples:
```go
// Import bitarray into your code and refer to it as `bitarray`
import "github.com/taki-mekhalfa/bitarray"
```
#### Appending `uints`
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
	ba.Append8(0xDE, 8) // Appends all the bits in the 0xDE byte
	ba.Append8(0x0A, 4) // Only appends 4 bits so only 0xA
	ba.Append32(0x0D, 4)
	ba.Append(0x00BEEF, 16) // Appends lowest 16 bits so only 0xBEEF
	// Print the bit sequence
	fmt.Printf("%#X\n", ba.Bytes())
	/* Output
	   0XDEADBEEF
	*/
```
### Appending individual bits
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append(0xD0, 8)
    ba.AppendZero()
    ba.AppendZero()
    ba.AppendBit(0)
    ba.AppendBit(0)
    ba.Append(0x0D, 4)
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDOOD
    */
```
The same goes for AppendOne

### Length and padding
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append(0xC0FFEE, 24)
    ba.Append(0x0E, 4)
    
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    // Print its length
    fmt.Println(ba.Len())
    // Print the number of padding bits
    fmt.Println(ba.Padding())
    // Print the byte slice length
    fmt.Println(len(ba.Bytes()))
    /* Output
        0XCOFFEEE0
        28 // We have 28 bits
        4 // indicating that 4 bits are used as a padding
        4 // 4 bytes means 32 bits but len=28 indicates that 4 bits are empty
    */
```

### GetBit, SetBit and ClearBit
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append8(0xFA, 8) // 0b11111010
    fmt.Printf("bit_%d=%d and bit_%d=%d\n", 0, ba.GetBit(0), 7, ba.GetBit(7))
    ba.ClearBit(0)
    fmt.Printf("bit_%d=%d and bit_%d=%d\n", 0, ba.GetBit(0), 7, ba.GetBit(7))
    ba.SetBit(0)
    fmt.Printf("bit_%d=%d and bit_%d=%d\n", 0, ba.GetBit(0), 7, ba.GetBit(7))
    /* Output
        bit_0=1 and bit_7=0
        bit_0=0 and bit_7=0
        bit_0=1 and bit_7=0
    */
```
### AppendBytes
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append(0xDE, 8)
    ba.Append(0xAD, 8)
    ba.Append(0xB, 4)
    ba.AppendBytes([]byte("\xEE\xFA"), 4) // 4 is used to indicate that there are 4 padding bits at the end
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDEADBEEF
    */
```
### AppendString
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.AppendString("11011110101011011011111011101111")
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDEADBEEF
    */
```

### Extract
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.AppendString("00010101100100101010110000001111111111101110101010101001001010")
    res := ba.Extract(20, 44)
    fmt.Printf("%#X\n", res)
    /* Output
        0XC0FFEE
    */
```

### ExtractBitArray
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
	ba.AppendBytes([]byte("\xAB\xDE\xAD\xBE\xEF\xCD"), 0)
	extracted := ba.ExtractBitArray(8, 8+32)
	fmt.Printf("%#X\n", extracted.Bytes())
	fmt.Println(extracted.Len())
    /* Output
        0XDEADBEEF
        32
    */
```

### AppendBitArray
```go
    ba := bitarray.New()
	ba.Append(0xDEAD, 16)
	ba1 := bitarray.New()
	ba1.Append(0xBEEF, 16)

	ba.AppendBitArray(ba1)
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDEADBEEF
    */
```