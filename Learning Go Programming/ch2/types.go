/*

Numeric types:

byte, int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64, complex64
rune

String type:
string

Boolean type:
bool

Error type:
error
++++++++++++++++++++++++++++++++++++++++++++++++
preassigned values

Boolean = true-false

constant counter = iota

Uninitialized value = nil
++++++++++++++++++++++++++++++++++++++++++++++++
Functions

Initialization = make(), new()

Collections = append(), cap(), copy(), delete()

Complex Numbers = complex(), imag(), real()

Error Handling = panic(), recover()

*/

package main

import (
	"fmt"
	"math/bits"
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
