# bitarray

Bitarray implements a slice like data structure for bits. The implementation is space efficient as bits are stored in actual memory bits

## Usage
The following shows some examples:
```go
// Import bitarray into your code and refer to it as `bitarray`
import "github.com/taki-mekhalfa/bitarray"
```
#### Append8 
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append8(0xDE)
    ba.Append8(0xAD)
    ba.Append8(0xBE)
    ba.Append8(0xEF)
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDEADBEEF
    */
```
#### Append16
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append16(0xBADD)
    ba.Append16(0x0C) // This will only append C=0b1100
    ba.Append16(0x0A) // This will only append A=0b1010
    ba.Append16(0xFE)
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XBADDCAFE
    */
```

The same goes for Append32 and Append64

### AppendZero
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append8(0xD0)
    ba.AppendZero();ba.AppendZero()
    ba.AppendZero();ba.AppendZero()
    ba.Append8(0x0D)
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    /* Output
        0XDOOD
    */
```
The same goes for AppendOne

### Len
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append32(0xC0FFEE)
    ba.Append8(0x0E)
    
    // Print the bit sequence
    fmt.Printf("%#X\n", ba.Bytes())
    // Print its length
    fmt.Println(ba.Len())
    // Print the byte slice length
    fmt.Println(len(ba.Bytes()))
    /* Output
        0XCOFFEEE0
        28 // We have 28 bits
        4 // 4 bytes means 32 bits but len=28 indicates that 4 bits are empty
    */
```

### GetBit, SetBit and ClearBit
```go
    // Create a fresh empty bit array
    ba := bitarray.New()
    ba.Append8(0xFA) // 0b11111010
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