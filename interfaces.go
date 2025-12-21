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
