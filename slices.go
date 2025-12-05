package main

import "fmt"

/**
 * Demonstrasi perilaku slices terhadap array asalnya
 * Ketika slice diubah, array asalnya juga ikut berubah
 */
func demonstrateSlicesBehaviour() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
