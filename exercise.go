package main

import (
	"fmt"
	"strings"
)

// Sqrt
/**
 * Exercise 1 : Loops and Functions
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

// Pic
/**
 * Excercise 2: Fungsi pic
 * Fungsi ini:
 * - Menerima lebar (dx) dan tinggi (dy) gambar
 * - Mengembalikan slice 2 dimensi [][]uint8
 * - Setiap nilai uint8 dianggap sebagai warna abu-abu (0-255)
 */
func Pic(dx, dy int) [][]uint8 {
	// membuat slice utama (baris)
	// slice dengan dy baris dan setiap barus nantinya berisi slice []uint8
	img := make([][]uint8, dy)

	// loop baris y, loop vertikal (atas ke bawah)
	for y := 0; y < dy; y++ {
		// membuat kolom untuk setiap baris
		img[y] = make([]uint8, dx)

		// loop kolom x, horizontal (kiri ke kanan)
		for x := 0; x < dx; x++ {
			// isi warna abu-abu
			img[y][x] = uint8(x * y)
		}
	}
	return img
}

// WordCount
/**
 * Exercise 3: Maps
 * Fungsi ini mengembalikan sebuah map dari perhitungan setiap kata di dalam string s.
 */
func WordCount(s string) map[string]int {
	// buat map kosong
	result := make(map[string]int)

	// pecah string menjadi kata-kata
	words := strings.Fields(s)

	// hitung setiap kata
	for _, word := range words {
		result[word]++
	}
	return result
}

// fibonacci mengembalikan sebuah closure
// Exercise 4: Fungsi dengan closure
// yang setiap kali dipanggil akan mengembalikan
// angka fibonacci berikutnya
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// Exercise 5 Stringers

type IPAddr [4]byte

// method String()
func (ip IPAddr) String() string {
	// ambil tipa elemen, gabungkan pakai titik, kemudian kembalikan sebagai String.
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Exercise 6 Error
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot calculate square root of negative number: %v", float64(e))
}

func CalculateSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}
