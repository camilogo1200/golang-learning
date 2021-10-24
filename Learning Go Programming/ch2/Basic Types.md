# Basic Types

## Integers (Signed and Unsigned)

Integers can be signed or unsigned

### Signed 

Signed integers are of 5 types

| **Type**  | **Size**               |                                                                                   |
|:----------|:-----------------------|-----------------------------------------------------------------------------------|
| **int**   | **Platform dependent** |                                                                                   |
| **int8**  | 8 bits/ 1byte          | -2^7 (**-128**) to 2^7-1 (**127**)                                                |
| **int16** | 16 bits / 2 bytes      | -2^15 (**-32.768**) to 2^15 -1 (**32.767**)                                       |
| **int32** | 32 bits /  4 bytes     | -2^31 (**-2.147.483.648**) to 2^31 -1 (**2.147.483.647**)                         |
| **int64** | 64 bits /  8 bytes     | -2^63 (**-9,223,372,036,854,775,808**) to 2^63 -1 (**9,223,372,036,854,775,807**) |

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
	var x int = 2
	var a int8 = 2
	var b int16 = 2
	var c int32 = 2
	var d int64 = 2

	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	//Size of int in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(x))
	fmt.Printf("x's type is %s\n", reflect.TypeOf(x))
	//Size of int8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
	//Size of int16 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("b's type is %s\n", reflect.TypeOf(b))
	//Size of int32 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
	//Size of int64 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))
}

```

### Unsigned 

Signed integers are of 5 types

| **Type**    | **Size**               |                                                                                                                                       |
|:------------|:-----------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| **uint**    | **Platform dependent** |                                                                                                                                       |
| **uint8**   | 8 bits/ 1byte          | 0 to 2^8-1 (**255**)                                                                                                                  |
| **uint16**  | 16 bits / 2 bytes      | 0 to 2^16 -1 (**65.535**)                                                                                                             |
| **uint32**  | 32 bits /  4 bytes     | 0 to 2^32 -1 (**4,294,967,295**)                                                                                                      |
| **uint64**  | 64 bits /  8 bytes     | 0 to 2^64 -1 (**18,446,744,073,709,551,615**)                                                                                         |
| **uintptr** | 64 bits /  8 bytes     | This is an unsigned integer type that is large enough to hold any pointer address. Therefore is size and range are platform dependent |


```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
	var x uint = 2
	var a uint8 = 2
	var b uint16 = 2
	var c uint32 = 2
	var d uint64 = 2

	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfUintInBits := bits.UintSize
    fmt.Println(sizeOfIntInBits)

	//Size of int in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(x))
	fmt.Printf("x's type is %s\n", reflect.TypeOf(x))
	//Size of int8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
	//Size of int16 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("b's type is %s\n", reflect.TypeOf(b))
	//Size of int32 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(c))
	fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
	//Size of int64 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(d))
	fmt.Printf("d's type is %s\n", reflect.TypeOf(d))
}

```

#### uintptr

**Properties**
- A _**uintptr**_ can be converted to unsafe.Pointer and viceversa
- Arithmetic can be performed on the _**uintptr**_
- _**uintptr**_ it holds a pointer address, is just a value and does not references any object.
    - **Therefore**
        - It's value will not be updated if the correspondent object moves. Ex. When goroutine stack changes
        - The corresponding object can be garbage collected.

**When to use**
- Its purpose is to be used along with unsafe.Pointer mainly used for _**unsafe memory access.**_
- When you want to save the pointer address value for printing it or storing it.
    - since the address is just stored and does not reference anything, the corresponding object can be garbage collected.




```go
package main
import (
    "fmt"
    "unsafe"
)
type sample struct {
    a int
    b string
}
func main() {
    s := &sample{a: 1, b: "test"}
    
   //Getting the address of field b in struct s
    p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))
    
    //Typecasting it to a string pointer and printing the value of it
    fmt.Println(*(*string)(p))
}

```

## Floats

| Type        | Size               |
|:------------|:-------------------|
| **float32** | 32 bits or 4 bytes |
| **float64** | 64 bits or 8 bytes |

**float64** is the default float type. 
When you initializea variable with a decimal value and don't specify the float type, the default type inferred wil be **float64**

### **float32** 
**float32** uses single-precision floating point format to store values. The 32 bits are divided into -1 bit sign, 8 bis exponent, and 23 bits mantissa.
- **Size:** 32 bits or 4 bytes
- **Range:** 1.2 E-38 to 3.4E+38
- **Default:** 0.0

### **float64**
**float64** uses a double-precision floating point format to store values. The 64 bits are divided into – 1-bit sign, 11 bits exponent, 52 bits mantissa.
- **Default:** 0.0


# Complex Numbers

| Type           | Size                                    |
|:---------------|:----------------------------------------|
| **complex64**  | real and imaginary part are **float32** |
| **complex128** | real and imaginary part are **float64** |

The default complex type is **complex128**

## initialization

Complex Numbers can be initialized in two ways

- using ```complex``` function. 
    ```go 
    c1:= complex(10,11)
    ```

- using shorthand syntax. if the type is not specified  the type is going to be a **complex128**
     
    ```go 
     c2 := 10 + 11i
     ```

## Parts of a complex number
There are two parts of a complex number. The real and imaginary part.
we use functions to get those

```go
    c := complex(23, 31)
    realPart := real(c)    // gets real part
    imagPart := imag(c)    // gets imaginary part
```

# Byte

**byte** in go is an alias for **uint8** meaning is an integer value. Thia integer value is 8 bits and it represents one byte number between (0-255).

## Define byte 

```go
var rbyte byte := 'a'
```
While declaring byte we have to specify the type. if we don't specify the type, the the result type is meant as a **rune**.

# Rune

rune in go is an alias for **int32** meaning it is an integer value.
This integer value is meant to represent a Unicode Code Point. To undestand rune you have to know what Unicode is.

### what's a GoLang rune
Strings are always defined using characters or bytes. In GoLang, strings are always made of bytes. Go uses UTF-8 encoding, so any valid character can be represented in Unicode code points.

A character is defined using **"code points"** in Unicode. Go language introduced a new term for this code point called **rune**.

- strings are made of bytes and they can contain valid characters that can be represented using runes.
- We can use ```rune()``` function to convert string to an array of runes

```go
// Finding rune of a character in Go
package main
 
import (
    "fmt"
)
 
func main() {
    s := 'a'
 
    s_rune := rune(s)
     
    fmt.Println(s_rune)
}
```

```go
// GoLang String to rune
package main
 
import (
    "fmt"
)
 
func main() {
    s := "GoLang"
 
    s_rune := []rune(s)
    fmt.Println(s_rune) // [71 111 76 97 110 103]
 
}
```
### Understanding difference between byte and rune

The special Unicode character Ö rune value is 214 but it's taking two bytes for encoding
```go
package main
 
import (
    "fmt"
)
 
func main() {
    s := "GÖ"
 
    s_rune := []rune(s)
    s_byte := []byte(s)
     
    fmt.Println(s_rune)  // [71 214]
    fmt.Println(s_byte)  // [71 195 150]
}
```

