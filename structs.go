package main

import "fmt"

/**
 * Pointer digunakan untuk menyimpan alamat memori dari sebuah variabel.
 */
func (v Vertex) showPointerStruct() {
	i, j := 42, 2701

	p := &i // menunjuk ke i jadi isinya 42
	fmt.Println(*p)
	*p = 21        // set i melalui pointer
	fmt.Println(i) // lihat i terbaru dari 42 jadi 21

	p = &j         // menunjuk ke j jadi isinya 2701
	*p = *p / 37   // bagi nilai j melalui pointer 2071 / 37 = 73
	fmt.Println(j) // lihat j terbaru dari 2701 jadi 73
}

// Vertex
/**
 * STRUCTS
 * Structs adalah tipe data komposit yang mengelompokkan
 * kumpulan properti/properti yang memiliki tipe data berbeda
 * menjadi satu kesatuan.
 *
 * Sesuaikan parameternya masing-masing.
 * karena ada yang perlu dan tidak perlu, jika tidak tinggal tambahkan nilai default nol aja.
 */
type Vertex struct {
	X int
	Y int
	Lat,
	Long float64
}
