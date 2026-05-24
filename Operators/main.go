package main

import (
	"fmt"
	"math"
)

func main() {
	//Math Operators
	/*
			1- require same type e.g ((int)123+(float)233.5)-> invalid operation
		    need to convert the float to in or vise versa to be valid operation
	*/
	var anInt int = 234
	var aFloat float64 = 2234.556
	//result :=  anInt + aFloat [will generaate mismatch types error]
	var convFloat int = int(aFloat)
	result := anInt + convFloat
	fmt.Println(result)

	i1, i2, i3 := 34, 34, 55
	f1, f2, f3 := 455.44, 4433.55, 2212.68

	var intSum int = i1 + i2 + i3
	fmt.Println("Integers sum  value", intSum)
	var floatSum float64 = f1 + f2 + f3
	fmt.Println("Floats sum  value", floatSum)
	finalProductPrice := float64(intSum) * floatSum
	fmt.Printf("final product price: %v\n", finalProductPrice)

	//Pi value
	fmt.Println("PI value is:-", math.Pi)
	//Rounded pi value float to decrease precession points
	fmt.Println("PI value is:-", math.Round(math.Pi*100)/100)
}
