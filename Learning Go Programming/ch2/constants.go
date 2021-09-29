/*
The value for a constant is static and cannot be changed after initial assigment
*/
package main

import (
	"fmt"
	"math"
	"time"
)

/*
Typed Contants
=================================================
Go uses the const keyword to indicate the declration of a constant
Unlike variables, however, the declaration must include the literal value to be bound to the identifier

const <identifier list> <type> = <value list or initializer expression>

Constants cannot have any dependency that requires runtime resolution
The compiler must be able to  resolve the value of a constant at compile time

*/

const a1, a2 string = "Mastering", "Go"
const b rune = 'G'
const c bool = false
const d int32 = 111009
const e float32 = 2.71828
const f float64 = math.Pi * 2.0e+3
const g complex64 = 5.0i
const h time.Duration = 4 * time.Second

/*
Untyped Constants
==================================================
Constants are even more interesting when they are untyped.

const <identifier list> = <value list or initializer expression>

The entire value of the resulting constant is stored in memory without any loss in precision
*/
const i = "G is" + " for Go "
const j = 'V'
const k1, k2 = true, !k1
const l = 111*100000 + 9
const m1 = math.Pi / 3.141592
const m2 = 1.41421356237309504880168872420969807856967187537688724209698078569671875376
const m3 = m2 * m2
const m4 = m3 * 1.0e+400
const n = -5.0i * 3
const o = time.Millisecond * 5

/*
Assigning untyped constants
===================================================
*/

func main() {
	var u1 float32 = m2
	var u2 float64 = m2
	u3 := m2

	fmt.Println(m2)
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)

}
