# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
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

```