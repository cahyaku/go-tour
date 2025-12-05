# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 3. Structs ✮⋆˙
#
### Pointers ✮⋆˙
```azure
Pointer digunakan untuk menyimpan alamat dari sebuah nilai.

Tipe *T menunjuk ke nilai bertipe T.
Nilai kosong dari pointer adalah nil.

var p *int  // p adalah pointer ke int, nil

Operator & mengembalikan alamat dari sebuah nilai.
i: = 42
p = &i      // p sekarang menyimpan alamat i

Operator * mengakses nilai yang ditunjuk oleh pointer.
fmt.Println(*p)         // mencetak 42
*p = 21                 // mengubah nilai i melalui pointer p
fmt.Println(i)          // mencetak 21
```
#
#
### Structs ✮⋆˙
```azure
Struct adalah kumpulan dari berbagai variabel.
Jadi jumlah varibel di dalam struct itu bebas bisa banyak, alias lebih dari dua.
    
type Vertex struct {
    X int
    Y int
}

output:
fmt.Println(Vertex{1,3}) // mencetak {1 3}
```
#
#
### Cara akses struct ✮⋆˙
```azure
Akses field struct menggunakan titik (.).
v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X) // mencetak 4
```
#
#
### Struct Fields ✮⋆˙
```azure
Bagian dari struct dapat diakses lewat pointer ke struct.

Untuk mengakses bagian X dari sebuah struc 
```
#
#
### Pointer to Structs ✮⋆˙
```azure
Bagian dari struct dapat diakases lewat pointer ke struct.

Untuk mengakses bagian X dari sebuah struct bila kita memiliki pointer ke struct p,
kita dapat menulis (*p).X.
Namun, Go secara otomatis mengubahnya menjadi p.X.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // awalnya vertex.X adalah 1, sekarang diubah menjadi 1e9
	fmt.Println(v)
}

Hasil output:
{1000000000 2}

Karena x diubah menjadi 1e9, tapi y tetap 2.
```
#
#
### Struct Literals (Inisialisasi strut)✮⋆˙
```azure
Struct bisa dibuat dengan mengisinya dengan nilai bagian bagiannya.
Selain itu, dapa mengisi hanya sebagian dari kolom dengan menggunakan sintaks
Name: (urutan nama dari bagian-bagian tidak berpengaruh).

type Vertex struct {
    X int
    Y int
}

var (
    v1 = Vertex{1, 2}  // X:1 dan Y:2 memiliki tipe Vertex
    v2 = Vertex{X: 1}  // X:1 dan Y:0 adalah input implisit
    v3 = Vertex{}      // X:0 dan Y:0
    p  = &Vertex{1, 2} // pointer ke Vertex{1, 2} memiliki tipe *Vertex
)

func main() {
    fmt.Println(v1, v2, v3, p)
}

Hasil output:
{1 2} {1 0} {0 0} &{1 2}
```
#
#
### Arrays ✮⋆˙
```azure
Deklarasi tipe dengan [n]T adalah untuk array dengan n dan tipe T.
Contoh membuat array: 
var a [10]int

membuat variabel a yang merupakan array dari 10 integer.

example:
var a [2]string
a[0] = "Hello"
a[1] = "World"
fmt.Println(a[0], a[1]) // mencetak Hello World
fmt.Println(a)          // mencetak [Hello World]

example:
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes) // mencetak [2 3 5 7 11 13]
```