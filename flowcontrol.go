package main

import (
	"fmt"
	"math"
	"time"
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

/**
 * Switch
 */
func printOperatingSystem() {

	fmt.Print("Go runs on ")
	switch os := "linux"; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

/**
 * Switch evaluation order
 * Kondisi switch dievaluasi dari atas ke bawah,
 * berhenti saat sebuah kondisi terpenuhi.
 */
func whenIsSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Sekarang.")
	case today + 1:
		fmt.Println("Besok.")
	case today + 2:
		fmt.Println("Dua hari lagi.")
	default:
		fmt.Println("Masih jauh.")
	}
}

/**
 * Switch without a condition
 */
func printGreeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Goof evening!")
	}
}

/**
 * Stacking Defers
 */
func exampleDeferLoop() {
	fmt.Println("Counting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
