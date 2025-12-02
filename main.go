package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

/**
 * Basic Types in Go
 */
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/**
 * 1. Fungsi untuk menjumlahkan dua bilangan bulat
 */
func add(a int, b int) int {
	return a + b
}

/**
 * 2. Penulisan singkat dari fungsi add di atas
 * Karena tipenya sama, kita bisa menuliskannya sekali saja
 */
func add2(a, b int) int {
	return a + b
}

/**
 * 3. Return multiple values
 * Sebuah fungsi dapat mengembalikan lebih dari satu nilai
 * Disini mengembalikan dua string
 */
func swap(x, y string) (string, string) {
	return y, x
}

/**
 * 4. Named return values
 * Fungsi ini menerima satu parameter bertipe int
 * dan mengembalikan dua nilai bertipe int
 * Jadi (x, y int) di dalam tanda kurung adalah nama dari return values
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/**
 * 5. Variabel global
 * Variabel c, python, java bertipe bool dengan nilai default false
 */
var c, python, java bool

var i, j int = 1, 2

/**
 * FUNGSI UTAMA
 */
func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
	fmt.Println()

	// menggunakan dari package time
	fmt.Println("The time is", time.Now())
	fmt.Println()

	// menggunakan dari package rand
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println()

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println()

	fmt.Println(add(1, 2))
	fmt.Println(add2(3, 2))
	fmt.Println()

	// Menampilkan hasil dari fungsi swap
	a, b := swap("Hello", "cahya")
	fmt.Println(a, b)
	fmt.Println()

	fmt.Println(split(17))
	fmt.Println()

	// Menampilkan nilai variabel global
	fmt.Println(c, python, java)
	fmt.Println()

	// Menampilkan nilai variabel i dan j yang dideklarasikan secara eksplisit
	// serta variabel c, python, java yang dideklarasikan dan diinisialisasi secara singkat
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println()

	// Deklarasi variabel singkat (Short variable declaration)
	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k)
	fmt.Println()

	// Basic Types in Go
	// Menampilkan tipe dan nilai dari variabel yang sudah dideklarasikan di awal
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
