package main

import "fmt"

// Person Membuat entitas Person
type Person struct {
	name string
}

// SayHello
/**
 * Method
 * Membuat method greet pada entitas Person
 */
func (p Person) SayHello() {
	fmt.Println("Hello, I'm", p.name)
}

// SayHello
/**
 * Method adalah function
 * Functionnya, perhatikan perbedaan penulisan dengan method
 * Jadi nama function di tulis terlebih dahulu baru argumen yang diperlukan
 */
func SayHello(p Person) string {
	return "Hello, I'm " + p.name
}

// Tipe tanpa struct
type MyFloat float64

// Method lanjutan
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
