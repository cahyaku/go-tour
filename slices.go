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

	a := names[0:2] // array slice dari index 0 dan 1, karena 2 tidak termasuk
	b := names[1:3] // array slice dari index 1 dan 2, karena 3 tidak termasuk
	fmt.Println(a, b)

	b[0] = "XXX" // mengubah index 1 dari array names melalui slice b[0], karena b[0] menunjuk ke names[1]
	fmt.Println(a, b)
	fmt.Println(names)
}

/**
 * Inisialisasi slice
 */
func combineIntAndBool() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	// Jika multi line, maka wajib diberikan tanda koma di akhir.
	// Ingat strct dapat memiliki banyak field dengan tipe berbeda.
	fmt.Println(s)
}

/**
 * Nilai default slice
 */
func sliceExample() {
	s := []int{2, 3, 5, 7, 11, 13}

	// Mengambil slice dari index 1 sampai 4 (4 tidak termasuk)
	s = s[1:4]
	fmt.Println(s)

	// Mengambil slice dari awal sampai index 2 (2 tidak termasuk)
	// indexnya kini adalah 0,1,2 dengan nilai 3,5,7
	// Jadi s[0]=3, s[1]=5, s[2]=7
	// Maka hasilnya adalah 3,5 (karena 2 tidak termasuk)
	s = s[:2]
	fmt.Println(s)

	// Mengambil slice dari index 1 sampai akhir
	// indexnya kini adalah 0,1 dengan nilai 3,5
	// Jadi s[0]=3, s[1]=5
	// Maka hasilnya adalah 5 (karena dimulai dari index 1)
	s = s[1:]
	fmt.Println(s)
}

/**
 * Demonstrasi panjang dan kapasitas slice
 */
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/**
 * Nil Slice
 */
func nilsSliceDemo() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSliceDemo(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sliceMakeLenCapDemo() {
	// Membuat slice int, len=5 dan cap=5 dengan nilai default 0
	a := make([]int, 5)
	printSliceDemo("a", a)

	// Membuat slice int, len=0 dan cap=5
	b := make([]int, 0, 5)
	printSliceDemo("b", b)

	// Menambah panjang slice b menjadi 2
	c := b[:2]
	printSliceDemo("c", c)

	// Menambah panjang slice c menjadi 5 mulai dari index 2
	d := c[2:5]
	printSliceDemo("d", d)
}

/**
 * Slice of slice
 */
func displayTicTacToeBoard() {

	// Membuat papan tic-tac-toe
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Giliran untuk para pemain
	board[0][0] = "x"
	board[2][2] = "o"
	board[1][2] = "x"
	board[1][0] = "o"
	board[0][2] = "x"

	// Menampilkan papan
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}
}

/**
 * Menambahkan elemen ke slice
 */
func sliceAppendDemo() {
	var s []int
	printSlice(s)

	// append pada slice nil, menyimpan elemen 0
	s = append(s, 0)
	printSlice(s)

	// menambah satu elemen lagi
	s = append(s, 1)
	printSlice(s)

	// menambah banyak elemen sekaligus
	s = append(s, 2, 3, 4)
	printSlice(s)
}

/**
 * Perintah Range
 */
func powerOfTwoRangeDemo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println(pow)
}
