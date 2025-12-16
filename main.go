package main

import (
	"fmt"
	"go-tour/contoh"
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
 * Constanta
 */
const (
	Big   = 1 << 100
	Small = Big >> 99
)

/**
 * Fungsi needInt menerima parameter bertipe int dan mengembalikan nilai bertipe int
 */
func needInt(x int) int {
	return x*10 + 1
}

/**
 * Fungsi needFloat menerima parameter bertipe float64 dan mengembalikan nilai bertipe float64
 */
func needFloat(x float64) float64 {
	return x * 0.1
}

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
	fmt.Println("==========================")
	fmt.Println("====== BASIC GO ======")
	fmt.Println("==========================")

	// menggunakan dari package time
	fmt.Println("The time is", time.Now())
	fmt.Println("----------------------")

	// menggunakan dari package rand
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("----------------------")

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("----------------------")

	fmt.Println(add(1, 2))
	fmt.Println(add2(3, 2))
	fmt.Println("----------------------")

	// Menampilkan hasil dari fungsi swap
	a, b := swap("Hello", "cahya")
	fmt.Println(a, b)
	fmt.Println("----------------------")

	fmt.Println(split(17))
	fmt.Println("----------------------")

	// Menampilkan nilai variabel global
	fmt.Println(c, python, java)
	fmt.Println("----------------------")

	// Menampilkan nilai variabel i dan j yang dideklarasikan secara eksplisit
	// serta variabel c, python, java yang dideklarasikan dan diinisialisasi secara singkat
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println("----------------------")

	// Deklarasi variabel singkat (Short variable declaration)
	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k)
	fmt.Println("----------------------")

	// Basic Types in Go
	// Menampilkan tipe dan nilai dari variabel yang sudah dideklarasikan di awal
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("----------------------")

	// Nilai kosong untuk tiap tipe data
	var i1 int
	var f1 float64
	var b1 bool
	var s1 string
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)
	fmt.Println("----------------------")

	// Konversi tipe data
	var x1, y1 int = 3, 4
	var f2 float64 = math.Sqrt(float64(x1*x1 + y1*y1))
	var z1 uint = uint(f2)
	fmt.Println(x1, y1, z1)
	fmt.Println("----------------------")

	// Inferensi tipe data: menentukan tipe data secara otomatis
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	fmt.Println("----------------------")

	// Membuat constanta
	const Pi = 3.14
	fmt.Println("Value of Pi:", Pi)
	const World = "世界"
	fmt.Println("Hello", World)
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println("----------------------")

	// Menggunakan constanta Big dan Small
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println()

	// FlOW CONTROL
	fmt.Println("==========================")
	fmt.Println("====== FLOW CONTROL ======")
	fmt.Println("==========================")

	fmt.Println("1. Perbedaan akses fungsi dalam package dan antar package")
	// LOOPING
	// 1. Ini diakses dari package lain yaitu package contoh
	contoh.PrintLoop()
	// 2. Ini diakses dari package utama (main), jadi masih satu package
	printLoop()
	fmt.Println()

	printLoop2()
	printLoop3()
	fmt.Println()

	// IF STATEMENT
	fmt.Println(sqrt(2))
	isEven(7)
	isEven(8)
	fmt.Println()

	// IF STATEMENT dengan statement short
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	// IF and Else
	fmt.Println(pow2(3, 2, 10), pow2(3, 3, 20))
	fmt.Println()

	// Exercise 1: Loops and Functions
	fmt.Println(Sqrt(2))
	fmt.Println()

	// Switch Statement
	printOperatingSystem()
	fmt.Println()

	// Switch evaluation order
	whenIsSaturday()
	fmt.Println()

	// Switch with no condition
	printGreeting()
	fmt.Println()

	// Defer statement
	// Menunda eksekusi hingga fungsi yang memanggilnya selesai
	// Jadi "World" akan dicetak setelah "Hello" dan "Cahya"
	defer fmt.Println("World") // ini akan muncul terakhir karena defer
	fmt.Println("Hello")
	fmt.Println("Cahya")

	// Contoh defer dalam loop
	exampleDeferLoop()
	fmt.Println()

	// FlOW CONTROL
	fmt.Println("=====================")
	fmt.Println("====== Structs ======")
	fmt.Println("=====================")

	// Pointer ke struct
	showPointerStruct()
	fmt.Println()

	// Structs
	// Jangan hiraukan Lat dan Long, karena itu di pakai pada maps
	// Jadi sementara bisa dikosongkan dulu
	fmt.Println(Vertex{1, 2, 0, 0})
	fmt.Println()

	// Akses properti struct
	v1 := Vertex{1, 2, 0, 0}
	fmt.Println("Ini structnya X:", v1.X)
	fmt.Println("Ini structnya Y:", v1.Y)
	fmt.Println()

	// Pointer ke struct
	v2 := Vertex{1, 2, 0, 0}
	p := &v2
	p.X = 1e9 // diubah melalui pointer menjadi 1e9
	fmt.Println("Setelah diubah melalui pointer, ini structnya X:", v2.X, "dan Y:", v2.Y)
	fmt.Println()

	// Inisialisasi struct dan pointer ke struct
	showVertexInitialization()
	fmt.Println()

	// Arrays
	showArrayExample()
	fmt.Println()

	fmt.Println("====================")
	fmt.Println("====== Slices ======")
	fmt.Println("====================")

	// Demonstrasi perilaku slices terhadap array asalnya
	demonstrateSlicesBehaviour()
	fmt.Println()

	// Inisialisasi slice
	combineIntAndBool()
	fmt.Println()

	// Nilai default slice
	sliceExample()
	fmt.Println()

	// PANJANG dan KAPASITAS SLICE
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	// Mengubah panjang slice menjadi 0
	s = s[:0]
	printSlice(s)
	// Menambah panjang slice menjadi 4
	s = s[:4]
	printSlice(s)
	// Menghapus dua nilai pertama dari slice
	s = s[2:]
	printSlice(s)
	fmt.Println()

	// Nil Slices
	nilsSliceDemo()
	fmt.Println()

	// Membuat slice dengan fungsi make
	sliceMakeLenCapDemo()
	fmt.Println()

	// Slice of slice
	displayTicTacToeBoard()
	fmt.Println()

	// Menambahkan elemen ke slice dengan fungsi Append
	sliceAppendDemo()
	fmt.Println()

	// Perintah Range
	powerOfTwoRangeDemo()
	fmt.Println()

	// Exercise 2: SLICE
	//pic.Show(Pic)

	// Maps
	fmt.Println("======================")
	fmt.Println("======== Maps ========")
	fmt.Println("======================")

	// Create maps
	createLocationMap()
	fmt.Println()

	// Inisialisasi Map
	printLocationMap1()
	fmt.Println()

	// Inisialisasi Map lanjutan
	printLocationMap2()
	fmt.Println()

	// Operasi Maps
	manageMapValues()
	fmt.Println()

	// Exercise 3
	//wc.Test(WordCount)

	// Nilai Fungsi
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println()

	// Fungsi closure, hmm aku pahamnya mirip seperti rekusif function sih.
	runAccumulatorDemo()
	fmt.Println()

	// Menampilkan hasil dari function Fibonacci dengan closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println()
}
