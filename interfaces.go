package main

import "fmt"

/**
 * 1. Interface dipenuhi secara implisit
 */
type I interface {
	M()
}

type T struct {
	S string
}

// Method berikut berarti type T mengimplementasikan interface I,
// tapi kita tidak perlu secara eksplisit mendeklarasikannya.
func (t T) M() {
	fmt.Println(t.S)
}

/**
 * 2. Nilai interface (Interface Values)
 */

// Interface dengan method tunggal
type Printer interface {
	Print()
}

// Struct pertama
type Message struct {
	Text string
}

// Method milik *Message
func (m *Message) Print() {
	fmt.Println(m.Text)
}

// Tipe lain yang juga mengimplementasikan Printer
type Number float64

func (n Number) Print() {
	fmt.Println(n)
}

// Fungsi untuk melihat isi interface
func showInterfaceValue(p Printer) {
	fmt.Printf("(nilai = %v, tipe = %T)\n", p, p)
}

/**
 * 3. Interface dengan nilai nil
 */
// Interface
type Player interface {
	Play()
}

// Struct
type Song struct {
	Title string
}

// Method dengan pointer receiver
func (song *Song) Play() {
	if song == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("Playing:", song.Title)
}

func describe(p Player) {
	fmt.Printf("(nilai = %v, tipe = %T)\n", p, p)
}

/**
 * 4. Nil interface values (Isi interface dengan nil)
 */
type Action interface {
	Do()
}

func showInfo(a Action) {
	fmt.Printf("(%v, %T)\n", a, a)
}

/**
 * 5. The empty interface (interface kosong)
 */
var value interface{}

func printInfo(data interface{}) {
	fmt.Printf("(nilai = %v, tipe = %T)\n", data, data)
}

/**
 * 6. semua langsung di main
 * 7. Penggunaan switch untuk tipe
 */
func cekTipe(data interface{}) {
	switch nilai := data.(type) {
	case int:
		fmt.Printf("Dua kali %v adalah %v\n", nilai, nilai*2)
	case string:
		fmt.Printf("%q adalah %v bytes panjangnya\n", nilai, len(nilai))
	default:
		fmt.Printf("Saya tidak kenal dengan tipe %T!\n", nilai)
	}
}

/**
 * Stringers
 */
type User struct {
	FullName string
	AgeYears int
}

func (u User) String() string {
	return fmt.Sprintf("%s (%d years)", u.FullName, u.AgeYears)
}
