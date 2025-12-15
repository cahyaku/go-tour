package main

import (
	"fmt"
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
