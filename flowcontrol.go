package main

import (
	"fmt"
	"math"
)

/**
 * Looping dengan for biasa
 */
func printLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

/**
 * Looping dengan for seperti while
 * Harus ada statement init dan post
 * ex. for ; sum < 1000;  {
 * ini auto hilang ketika di save, mungkin karena langsung diformat
 */
func printLoop2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/**
 * Looping dengan for seperti while tanpa statement init dan post
 */
func printLoop3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/**
 * IF STATEMENT
 */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/**
 * IF STATEMENT simple
 */
func isEven(x int) {
	if x%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}

/**
 * IF with a short statement
 * Kondisi if singkat
 */
func pow(x, n, lim float64) float64 {
	// Hasil pangkat dari math.Pow disimpan di variabel v
	// Jika v kurang dari lim, maka kembalikan nilai v
	// Jika tidak, kembalikan nilai lim
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

/**
 * IF and ELSE
 */
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v tidak dapat digunakan disini
	return lim
}
