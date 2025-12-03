# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 1. Basic ✮⋆˙
#
### Syntax dasar GO ✮⋆˙
```azure
1. :=   digunakan untuk deklarasi dan inisialisasi variabel secara bersamaan.
2. var  digunakan untuk deklarasi variabel tanpa inisialisasi.
```
##### ★ contoh deklarasi variabel
```azure
1. x := 10              Deklarasi dan inisialisasi variabel x dengan nilai 10
2. z, y := 15, 25       Deklarasi dan inisialisasi variabel z dan y dengan nilai 15 dan 25
3. var y int            Deklarasi variabel y bertipe int tanpa inisialisasi
4. var z = 20           Deklarasi variabel z dengan inisialisasi nilai 20
```
#
#
### Public vs Private ✮⋆˙
```azure
1. Public
   - Nama variabel, fungsi, atau tipe data yang diawali dengan huruf kapital (A-Z) bersifat public.
   - Dapat diakses dari paket lain.
1. Private
   - Nama variabel, fungsi, atau tipe data yang diawali dengan huruf kecil (a-z) bersifat private.
   - Hanya dapat diakses dalam paket yang sama.
```
#
#
### var vs :=
```azure
1. var
   - Digunakan untuk mendeklarasikan variabel dengan tipe data yang eksplisit.
   - Contoh: var x int = 10
2. :=
   - Digunakan untuk mendeklarasikan dan menginisialisasi variabel secara otomatis dengan tipe data yang sesuai.
   - Contoh: x := 10
   
var dapat digunakan berulang kali untuk mendeklarasikan variabel baru,
sedangkan := hanya dapat digunakan sekali untuk mendeklarasikan variabel baru dalam satu blok kode.
```
##### ★ contoh var dapat digunakan berulang dalam 1 file tapi beda scope
```go
package main
import "fmt"

var x = 10 // scope: package
func main() {
	fmt.Println(x) // 10
	var x = 20     // scope: fungsi main (menimpa x package)
	fmt.Println(x) // 20
	{
		var x = 30     // scope: blok ini saja
		fmt.Println(x) // 30
	}
	fmt.Println(x) // 20 (block sebelumnya tidak mempengaruhi)
}
```
#
#
### Format verbs ✮⋆˙
```azure
Lebih lengkap cari saja format verbs di dokumentasi resmi Go: https://pkg.go.dev/fmt
1. %v       nilai default
2. %#v      nilai lengkap (dengan tipe data)
3. %T       tipe data
4. %%       karakter persen
```
#
#
### Import Package ✮⋆˙
```azure
1. import "fmt"         
   Mengimpor package fmt untuk input/output
   Untuk format dan menampilkan output (println, printf, input, format string, dll).

2. import "math"     
   Mengimpor package math untuk operasi matematika
   Untuk operasi matematika seperti akar kuadrat, trigonometri, dll.
```
#
#
### Exported names ✮⋆˙
```azure
1. Nama yang diawali dengan huruf kapital (misalnya, Println, Sqrt) dapat diakses dari package lain.
2. Nama yang diawali dengan huruf kecil (misalnya, println, sqrt) hanya dapat diakses dalam package yang sama.
Example:
   - fmt.Println()                      dapat diakses dari package lain
   - math.Sqrt()                        dapat diakses dari package lain
   
    - myPackage.myFunction()            hanya dapat diakses dalam package yang sama
   - myPackage.myVariable               hanya dapat diakses dalam package yang sama
    
```
#
#
### Function ✮⋆˙
```azure
1. func namaFunction(parameter tipe) tipeKembalian {
       // kode fungsi
   }
2. func add(a int, b int) int {
       return a + b
   }
```
#
#
### Function Continued ✮⋆˙
```azure
1. func namaFunction(parameter tipe) (tipeKembalian1, tipeKembalian2) {
       // kode fungsi
   }
2. func add(a, b int) (int) {
       return a + b
   }
```
#
#
### Multiple results ✮⋆˙
```azure
func swap(a int, b int) (int, int) {
   return b, a
}

Cara pemanggilan:
    x, y := swap(3, 4)
    Keterangan: x akan berisi 4 dan y akan berisi 3
                Jadi benar memang di tukar, karena swap artinya menukar
                itu ditukar pada func swap di atas.
```
#
#
### Named return values ✮⋆˙
```azure
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

Keterangan: 
    1 parameter input yaitu sum bertipe int
    2 nilai kembalian yaitu x dan y, keduanya bertipe int
    Jadi dengan named return values, kita bisa memberikan nama pada nilai kembalian langsung di deklarasi fungsi.
    Jumlahnya juga bisa lebih dari satu.
    
Cara pemanggilan:
    a, b := split(17)
    Keterangan: a akan berisi 7 dan b akan berisi 10
```
#
#
### Variables ✮⋆˙
```azure
1. var namaVariabel tipe = nilai
2. var x int = 10

Tapi bisa juga ditulis pada tingkat package atau function level seperti ini:

    var c, python, java bool
    func main() {
        var i int
        fmt.Println(i, c, python, java)
    }
```
#
#
### Variables with initializers ✮⋆˙
```azure
Deklarasi variabel bisa mengikutkan inisasialisasi, satu inisialisasi per variabel.

var i, j int = 1, 2
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```
#
#
### Short variable declarations ✮⋆˙
```azure
var i, j int = 1, 2
k := 3

func main() {
    fmt.Println(i, j, k)
}
```
#
#
### Basic types ✮⋆˙
```azure
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias untuk uint8

rune // alias untuk int32
     // merepresentasikan sebuah kode Unicode

float32 float64

complex64 complex128
-----------------------------

Ketika membutuhkan nilai integer anda sebaiknya menggunakan tipe int 
kecuali memiliki alasan tertentu menggunakan tipe berukuran khusus atau unsigned integer.
```
#
#
### Zero values (Nilai kosong) ✮⋆˙
```azure
Nilai kosong adalah:
1. 0 untuk tipe numerik
2. false untuk tipe boolean
3. "" (string kosong) untuk tipe string
```
#
#
### Type conversions ✮⋆˙
```azure

var x int = 42
var y float64 = float64(x)
var z uint = uint(y)

Keterangan:
    1. Konversi tipe harus dilakukan secara eksplisit di Go.
    2. Tidak ada konversi otomatis antar tipe.
```
#
#
### Type inference (Inferensi tipe) ✮⋆˙
```azure
Saat mendeklarasikan variabel tanpa menyebutkan tipe, Go akan menginferensi tipe berdasarkan nilai inisialisasi.
Example:
    var i int 
    j := i              // j diinferensi sebagai int karena i bertipe int
    
    var x = 42          // x diinferensi sebagai int
    var y = 3.14        // y diinferensi sebagai float64
    var s = "hello"     // s diinferensi sebagai string
```
#
#
### Constants-konstanta tidak bertipe ✮⋆˙
```azure
Constantas dideklarasikan dengan kata kunci const dan nilainya tidak dapat diubah setelah dideklarasikan.
Konstanta dapat berupa tipe numerik, string, atau boolean.
Example:
    const Pi = 3.14
    const Truth = true
    const World = "世界"
    const Cahya = "Cahya CutY"
```
#
#
### Numeric constants-konstanta bertipe ✮⋆˙
```azure
Konstanta yang tidak bertipe menggunakan tipe yang dibutuhkan oleh konteksnya.
Example:
    const Big = 1 << 100
    const Small = Big >> 99
```
#
#
# Constants with type vs Constants without type ✮⋆˙
```azure
Konstanta bertipe memiliki tipe yang tetap, sedangkan konstanta tidak bertipe dapat digunakan dalam berbagai konteks tipe.
example tidak bertipe:
    const a = 10     // untyped integer
    const b = 3.14   // untyped float
    const c = "hi"   // untyped string

    var x int = a    // boleh
    var y float64 = b // boleh
    var z string = c  // boleh

    
example bertipe:
    const a int = 10
    const b float64 = 3.14
    const c string = "hi"
    
    var x int = a       // boleh
    var y float64 = b   // boleh
    var z string = c    // boleh

// var f float64 = a   // ❌ error, a bertipe int
```

