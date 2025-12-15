# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 3. Slices ✮⋆˙
#
### Slices ✮⋆˙
```azure
1. Sebuah array memiliki ukuran yang tetap. 
2. Sebuah slice ukurannya bisa dinamis.
3. Slice lebih sering digunakan dibandingkan array.
4. Setiap ada slice index dimulai dari 0 lagi, jadi diulangi dari 0.
    
Tipe []T adalah slice dengan elemen bertipe T.
Slice dibantuk dengan menspesifikasikan dua indeks, batas bawah dan batas atas
    a[bawah:atas]

Notasi di atas memotong rentang dari slice a yang mengikutkan elemen bawah, tapi tidak memasukkan bagian terakhir atas.

Ekspresi berikut membuat sebuah slice yang mengikutkan elemen 1 sampai 3 dari slice a
    a[1:4]
    
Example:
    func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)
}

Hasilnya:[3 5 7] 
Tujuh tidak termasuk karena batas atas tidak termasuk dalam slice.
```
#
#
### Slices adalah referensi ke array ✮⋆˙
```azure
Slice adalah referensi ke array.
1. Slice tidak menyimpan data, ia hanya mendeskripsikan sebuah bagian dari array.
2. Mengubah nilai dari elemen dari sebuah slic juga mengubah elemen di array-nya.
3. Slice lain yang berbagi elemen array yang sama akan mendapatkan perubahan yang sama.

Example:
func main() {
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
    
Hasilnya:
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George] // Slice hanya menunjukkan array yang sama, tapi b mengubah nilai di index 1
[John XXX George Ringo] // Karena index 1 diubah ke "XXX", maka names juga berubah

```
#
#
### Slice Literal (Inisialisasi Slice) ✮⋆˙
```azure
Menginisialisasi slice mirip dengan array tapi tanpa mendefinisikan panjangnya.           
Misalnya: [3]bool{true, true, false}.
Berikut ini membuat array yang sama kemudian membuat slice yang mengacu kepadanya:
[]bool{true, true, false}

example: pada func combineIntAndBool().
```
#
#
### Nilai Default Slice ✮⋆˙
```azure
Saat memotong slice,
kita dapat memperhatikan batas bawah atau atas sehingga GO
akan menggunakan nilai defaultnya.

example: pada method sliceExample().
```
#
#
### Slice Lenght dan Capacity ✮⋆˙
```azure
Slice memiliki panjang dan kapasitas.
1. len() digunakan untuk mendapatkan panjang slice saat ini.
2. cap() digunakan untuk mendapatkan kapasitas slice, 
   yaitu jumlah elemen maksimum yang dapat diakses dari posisi awal slice hingga akhir array.

Example: pada method printSlice()
```
#
#
### Nil Slice ✮⋆˙
```azure
Slice bisa bernilai nil.
1. Slice nil memiliki panjang dan kapasitas 0 dan tidak menunjuk ke array manapun.
2. Slice nil berbeda dengan slice kosong, yang memiliki panjang dan kapasitas 0 tetapi menunjuk ke array.
3. Slice nil dapat diperiksa dengan membandingkannya dengan nil.

Example: pada method nilSliceDemo()
```
#
#
### Membuat Slice dengan `make` ✮⋆˙
```azure
Slice dapat dibuat dengan fungsi make, 
kita juga dapat membuat array dengan ukuran yang dinamis.

Fungsi make mengalokasikan array yang dikosongkan dan 
mengembalikan sebuah slice yang mengacu pada array tersebut.

make([]T, len, cap)
len = jumlah elemen yang langsung aktif
cap = kapasitas maksimum yang tersedia di belakang memori slice
Jika cap tidak ditulis, maka cap == len.

Example: pada func sliceMakeLenCapDemo()
```
#
#
### Slice dari Slice ✮⋆˙
```azure
Slice dapat dibuat dari slice lain.

Example: pada func displayTicTacToeBoard()
```
#
#
### Menambahkan Elemen ke Slice ✮⋆˙
```azure
Menambahkan elemen ke slice dengan fungsi append.

func append(s []T, vs ...T) []T
Paramater s - sebuah slice dari tipe T
Paramater vs - elemen-elemen bertipe T yang akan ditambahkan ke slice s

Hasil dari append adalah sebuah slice yang berisi semua elemen
dari slice awal dengan nilai yang ditambahkan.

Jika array awal s terlalu kecil untuk nilai yang ditambahkan,
maka array yang lebih besar akan dialokasikan. Kembaliannya
yaitu slice yang merujuk ke array baru.

Example: pada func sliceAppendDemo()
```
#
#
### Range ✮⋆˙
```azure
Perintah range pada perulangan for digunakan untuk mengintegrasikan sebuah slice atau map.

Saat melakukan perulangan dengan range pada slice, dua nilai akan dikembalikan pada setiap iterasi.
Nilai pertama yaitu indeks, dan nilai kedua yaitu salinan dari elemen pada indeks tersebut.

Simplenya range: digunakan untuk mengambil index dan nilai satu-per-satu.
Example:

func main() {
    angka := []it {10,20,30}
    for i,v := range angka {
        fmt.Println(i,v)
    }
}

Hasil output:
0 10
1 20
2 30

Jadi i adalah index dan v adalah nilai pada index tersebut.
```
#
#
### Range Continued ✮⋆˙
```azure
Gunakan operator _ untuk melewatkan nilai index.
Example:

func main() {
    angka := []int{10,20,30}
    for _,v := range angka {
        fmt.Println(v)
    }
}

Jadi hanya nilai v yang akan diambil.
Karena operator _ digunakan untuk melewatkan nilai index, begitu pula sebaliknya.

```