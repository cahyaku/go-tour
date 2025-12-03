package main

import (
	"fmt"
)

/**
 * Exercise 1: Loops and Functions
 */
func Sqrt(x float64) float64 {
	z := 1.0 // tebakan awal
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iterasi", i+1, ":", z)
	}
	return z
}
