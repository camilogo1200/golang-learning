/*
Variable Declaration

The long form of a variable declaration in Go follows the next format:

var <identifier list > <type>

The var keyword is used to declare one or more variabel identifiers,
followed by the type of the variables


*/

package main

import "fmt"

var name, desc string
var radius int32
var mass float64
var active bool
var satellites []string

// initialized declaration
// var <identifier list> <type> = <value list or initilize expression>
var name2, desc2 string = "Earth", "Planet"

//ommiting variable types
// var <identifier list> = <value list or initilize expression>
var name3, desc3 = "Mars", "Planet"

/*
inferred types
===================
Doble or single quoted (raw) text -> string

Integers: -> int
Decimals (-0.25, 4.0, 7e-12) -> float64

complex Numbers ( -5i, 3i, (0+4i) )-> complex128

Booleans (true, false) -> bool

array values ([2]int{-76, 8080})  -> literal array form [2]int

Map Values :
map[string]int { "sun": 685800, "Earth":6378 } -> literal value map[string]int

Slice values:
[]int{-76, 0, 1244} -> slice type literal -> []int

Struct Values -> literal structure
struct { name string, diameter int}{ "Mars": 3396 }
*/

//Short Variable Declaration
// <identifier list> := <values list or initialized expression>

/*
variable declaration block
===============================
var (
  name string = "Earth"
  desc string = "Planet"
  radius int32 = 6378
  mass float64 = 5.972E+24
  active bool = true
  satellites []string
)

*/

func main() {

	name = "sun"
	desc = "star"
	radius = 685800
	mass = 1.989e+30
	active = true
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	name4 := "Neptune"
	desc4 := "planet"
	satellites4 := []string{
		"Naiad", "Thalassa", "Despina", "Galatea", "Larissa",
		"S/2004 N 1", "Proteus", "Triton", "Nereid", "Halimede",
		"Sao", "Laomedeia", "Neso", "Psamathe",
	}

	fmt.Println(name)

	fmt.Println(desc)
	fmt.Println("Radius (km) ", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Satellites ", satellites)

	fmt.Println(name4)
	fmt.Println(desc4)
	fmt.Println(satellites4)
}
