package main

import (
	"fmt"
)

// Sqrt
/**
 * Exercise: Loops and Functions
 * Fungsi Sqrt menghitung akar kuadrat dari x menggunakan metode Newton-Raphson
 */
func Sqrt(x float64) float64 {
	z := 1.0 // tebakan awal
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iterasi", i+1, ":", z)
	}
	return z
}
