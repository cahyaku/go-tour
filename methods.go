package main

import (
	"fmt"
	"math"
)

/**
 * 1.  Person Membuat entitas Person
 */
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
 * 2. Method adalah function
 * Functionnya, perhatikan perbedaan penulisan dengan method
 * Jadi nama function di tulis terlebih dahulu baru argumen yang diperlukan
 */
func SayHello(p Person) string {
	return "Hello, I'm " + p.name
}

// Tipe tanpa struct
type MyFloat float64

/**
 * 3. Method lanjutan
 */
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/**
 * 4. Pointer-receiver
 */
type Vertex3 struct {
	X, Y float64
}

// Dengan value receiver (hanya mengubah salinannya)
// method Abs() hanya membaca nilai, tidak mengubah data
func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Pinter-receiver (mebgubah data aslinya - mengubah isi struct)
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
 * 5. Pointers and Functions (Pointer dan fungsi)
 */
type Vertex4 struct {
	X, Y float64
}

// menerima salian Vertex
func Abs(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// fungsi scale mengubah nilai asli vertex 4
func Scale(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
 * 6. Method and pointer indirection
 */
// buat structnya dulu
type Vertex5 struct {
	X, Y float64
}

// Scale Method scale
func (v *Vertex5) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc2 function scale
func ScaleFunc2(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
 * 7. Method and pointer indirection (2)
 */
type Vertex6 struct {
	X, Y float64
}

// methods
func (v Vertex6) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Function
func AbsFunc(v Vertex6) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
 * 8. Memilih nilai receiver sebagai nilai atau pointer
 */
type Vertex7 struct {
	X, Y float64
}

func (v *Vertex7) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
