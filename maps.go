package main

import "fmt"

/**
 * Create Map
 */
func createLocationMap() {
	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		0, 0, 40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/**
 * Inisialisasi Map
 */
func printLocationMap1() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{0, 0, 40.68433, -74.39967},
		"Google":    Vertex{0, 0, 37.42202, -122.08408},
	}
	fmt.Println(m)
}

/**
 * Inisialisasi map lanjutan
 * Abaikan saja nol pada vertex, karena itu digunakan pada bagian lain, jadi sementara saja
 */
func printLocationMap2() {
	var m = map[string]Vertex{
		"Bell Labs": {0, 0, 40.68433, -74.39967},
		"Google":    {0, 0, 37.42202, -122.08408},
	}
	fmt.Println(m)
}

/**
 * Map Operations
 * Managing Map Values
 */
func manageMapValues() {
	fmt.Println("--- Map Operations ---")

	m := make(map[string]int) // membuat map

	m["Answer"] = 42 // mengisi atau mengubah elemen dalam map
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // menghapus elemen dalam map
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // mengecek apakah elemen dalam map ada
	fmt.Println("The value:", v, "Present?", ok)
}

/**
 * Nilai Fungsi
 */
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/**
 * Function closure
 * dan di bawah terdapat function runAccumulatorDemo()
 */
func createAccumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Function menjalakan demo accumulator
func runAccumulatorDemo() {
	pos, neg := createAccumulator(), createAccumulator()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
