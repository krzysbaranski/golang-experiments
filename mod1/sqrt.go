package main

import (
	"fmt"
	"math"
)


func Sqrt(x float64) float64 {
	var z float64
	const DIFF = 0.000000000000000001
	z = 1.0
	for i := 0; i < 100; i++ {
		//fmt.Println(i)
		//fmt.Print(" ")
		//fmt.Print(" z*z=")
		//fmt.Print(z * z)
		//fmt.Print(" ")
		diff := (z*z - x) / (2 * z)
		//fmt.Print("diff=")
		//fmt.Print(diff)
		//fmt.Print(" z=")
		//fmt.Println(z)
		z = z - diff
		if (diff > 0 && diff < DIFF) || (diff < 0 && diff > DIFF*-1) {
			return z
		}

	}
	return z
}

func main() {
	var s = Sqrt(2)
	fmt.Printf(`sqrt is %v (%v^2=%v)`, s, s,s*s)
	fmt.Println()
	fmt.Printf(`diff from stdlib %v`, s-math.Sqrt(2))

}
