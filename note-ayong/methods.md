# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★ ּ ֶָ֢.
#
#
## 5. °:•.🐠*.•🪸.•:° Methods °:•.🐠*.•🪸.•:° 𖹭 
#
### Method ✮⋆˙
```azure
Go tidak memiliki class.
Namun ada method yang bisa mendefinisikan method pada tipe.
Sebuah method adalah sebuah fungsi dengan argumen khusus receiver.

Receiver adalah cara Go menempelkan function ke sebuah struct (atau tipe lain)
sehingga function itu menjadi method.

Bentuk umum reciver:
func (r ReciverType) MethodName(){
    //... body
}

atau dengan pointer:
func (r *ReciverType) MethodName(){
    //... body
}

Keterangan:
r => adalah receiver variable.
receiverType => adalah tipe yang menerima method.
```
#
#
### Method berdasarkan pemahaman cahya ⋆.˚🦋༘⋆
```azure
Method adalah fungsi yang menerima sebuah receiver.

1. Method akan memudahkan pemanggilan function.
2. Method harus ditulis setelah nama struct.

ex. func (p Person) SayHello() {
    fmt.Println("Hello,", p.Name)
}

Keterangan:
p => adalah receiver variable.
Person => adalah tipe yang menerima method.
Person => adalah nama struct.
SayHello => adalah nama method.
(p Person) => adalah recesiver.

Syarat membuat method:
1. Method hanya bisa merujuk pada satu struct.
2. Method harus memiliki 1 receiver.
3. Receiver tidak boleh berupa pointer.
   ex. func (p *Person) SayHello() { } // ini tidak valid.
   ex. func (p Person) SayHello() { }  // ini valid.
4. Cara mengakses method:

Intinya struct atau entity-nya wajib di definisikan dulu atau diinisialisasi.
Beru bisa memanggil methodnya.
    
// Entitas Person 
type Person struct {
    Name string
}

// Functon SayHello
func (p Person) SayHello() {
    fmt.Println("Hello,", p.Name)
}
    
// Memanggil method SayHello
p := Person{""} // inisialisasi struct
p.SayHello() // memanggil method
```
#
#
### Methods are functions (method adalah func) ✮⋆˙
```azure
Sebuah method hanyalah fungsi dengan argumen sebuah receiver.
Misalkan functin SayHello() menerima sebuah receiver type Person.

Jadi versi function biasanya, SayHello di tulis duluan sebagai berikut:
func SayHello(p Person) string {
    return "Hello," + p.Name
}

Keterangan:
p => adalah receiver variable.
Person => adalah tipe yang menerima method.
string => adalah tipe return dari function.

Jadi perbedaan method dengan function adalah:
1. Method harus ditulis setelah nama struct.
2. Function tidak perlu ditulis setelah nama struct.
```
#
#
### Method lanjutan ✮⋆˙
```azure
1. Method dapat dideklarasikan pada tipe lain (selain struct).
2. Dapat mendeklarasikan method dengan receiver yang tipenya didefinisikan di paket yang sama.

ex.
// membuat tipe Person BUKAN struct
type MyFloat float64

// membuat method Abs pada tipe MyFloat
func (f MyFloat) Abs() float64  {
    if f < 0 {
       return float64(-f)
    }
    return float64(f)
}

// menginisialisasi MyFloat
    f := MyFloat(-math.Sqrt2) 
        
// memanggil method Abs pada tipe MyFloat
    fmt.Println(f.Abs())
```
#
#
### Pointer-receiver ✮⋆˙
```azure

1. Mendeklarasikan method dengan receiver berupa pointer.
2. Tipe receiver memiliki sintaks *T untuk tipe T. 
   (T tidak bisa berupa pointer ke tipe dasar seperti *int).
3. Method pointer receiver dapat mengubah nilai yang ditunjukkan.
   Karena method kadang perlu mengubah receivernya,
   pointer-receiver lebih umum digunakan daripada receiver dengan value.

Jadi ada dua jesni recevier:
1) Value receiver -> func (v Vertex)...
2) Pointer receiver -> func (v *Vertex)...

Perbedaannya:
1) func (v Vertex) Scale(f float64)
   // Ini mengubah foto kopi (aggap kertas asli tidak berubah).
2) func (v *Vertex) Scale(f float64)
  // Ini mengubah kertas aslinya (nilai aslinya berubah)
  // Karena pointer mengarah langsung ke lokasinya.

Kesimpulan:
- Versi pointer lebih sering dipakai karena bisa menguba data.
- Lebih hemat memori (tidak copy struct besar)
- Konsisten (kalau satu method pakai pointer, biasanya semua pakai pointer)
```
#
#
### Pointers and Functions (Pointer dan fungsi) ✮⋆˙
```azure
1. Tanpa pointer (T) -> fungsi menerima salinan data.
2. Dengan pointer (*T) -> fungsi menerima alamat data asli
Yup intinya masih sama, untuk mengubah data asli harus pakai pointer.
    
Example:
Misalnya ada sebuah struct berikut:
type Vertex struct {
    X, Y float64
}

1. Versi tanpa pointer (tidak bisa mengbuah data asli)
func Scale (v Vertex, f float64){
    v.X = v.X * f
    v.Y = v.Y * f
}

Jika diinisialisasi:
v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(v)    

Hasil output: {3, 4}
Keterangan: data tidak berubah, karena v dalam Scale itu copy (jadi itu hanya mengubah salianan).
    
2. Versi method dengan pointer (idiom GO)
func (V *Vertex) Scale(f float64) {
    v.X *= f
    v.Y *= f
}

Jika dipanggil
v.Scale(10)

Go otomatis mengirimkan &v ke method pointer receiver
Scale (&, 10)

Jika * dihapus, misalkan
func Scale(v Vertex, f float64) {
	v.X *= f
}

Program tetap bisa dicompile tapi datanya tidak berubah.

```
#
#
### Methods and Pointer Indirection (Method dan pointer tidak langsung) ✮⋆˙
```azure
Pada Func:
1. Kalau parameternya *T, Harus dikirim pointer (&v)
2. Go tidak otomatis mengubah value jadi pointer

Pada Method:
1. Bisa dipanggil pakai value atau pointer
2. Go otomatis mengubah value -> pointer jika perlu

Example:
1. Pada Function (fungsi harus eksplisit pakai pointer)
    func ScaleFunc(v *Vertex, f float){
        v.X = v.X * f
        v.Y = v.Y * f
    }

    // pemanggilnya
    var v Vertex

    ScaleFunc(&v, 5)     => BENAR
    ScaleFunc(v, 5)      => SALAH (karena ScaleFunc meminta *Vertex)

2. Pada method dengan pointer-receiver (lebih flesibel)
    func (v *Vertex) Scale(f float64) {
        v.X = v.X * f
        v.Y = v.Y * f
    }
    
    // pemanggilnya
    var v Vertex
    v.Scale(5)          => BENAR

    p:= &v
    p.Scale(10)         => BENAR (dan lebih fleksibel)
    
Mengapa v.Scale(5) boleh? Padahal v itu bukan pointer.
Karena GO otomatis melakukan ini:
v.Scale(5) -> diubah menjadi (&v).Scale(5)
```
#
#
### Methods and Pointer Indirection (Method dan pointer tidak langsung) 2 ✮⋆˙
```azure
Fungsi yang menerima argumen nilai harus mendapatkan nilai dari tipe yang spesifik.
Intinya:
- Function = ketat soal tipe.
- Method = lebih fleksibel dan pintar.

example:
1) Function: tipe harus pas (ketat)
    // Function ini minta value Vertex, bukan pointer.
    func AbsFunc(v Vertex) float64 {
        return v.X + v.Y
    }
    
    // Pemanggilnya
    AbsFunc(v) // OK Vertex itu value
    AbsFunc(&v) // ERROR (bukan vertex, tapi *Vertex)
     
2) Method: lebih fleksibel    
    func (v Vertex) Abs() float64 {
        return v.X + v.Y
    }
            
    // Pemanggilnya
    var v Vertex
    fmt.Println(v.Abs()) // OK
    
    p := &v
    fmt.Println(p.Abs()) // OK
    
    
Example sederhana:
    type User struct {
	Name string
}

// FUNCTION (ketat)
func ShowFunc(u User) {
	fmt.Println(u.Name)
}

// METHOD (fleksibel)
func (u User) ShowMethod() {
	fmt.Println(u.Name)
}

func main() {
	u := User{"Cahya"}
	p := &u

    // kalau function ketat:
	ShowFunc(u)    // ✅
	// ShowFunc(p) // ❌ ERROR

    // kalau method fleksibel:
	u.ShowMethod() // ✅
	p.ShowMethod() // ✅ (auto dereference)
}
```
#
#
### Memilih receiver sebagai nilai atau pointer ✮⋆˙
```azure
Ada dua alasan mengapa memilih receiver sebagai nilai atau pointer:
1. Supaya method dapat mengubah nilai yang ditunjuk oleh receiver.
2. Menghindari menyalin nilai setiap kali method dipanggil.

Hal ini akan lebih efisien jika receiver adalah sebuah struct yang besar.

Intinya:
- Value receiver:
    func (v Vertex) Method()
- Pointer receiver:
    func (v *Vertex) Method()

Tujuan penggunaan pointer receiver adalah susya bisa mengubah data asli.

Kesimpulan:
Jika sebuah method perlu mengubah isi struct, maka wajib menggunakan pointer receiver.

Example:
    type Counter struct {
        Value int
    }
    
    func (c *Counter) Increment() {
        c.Value++
    }
    
    func (c *Counter) Get() int {
    return c.Value
    }

Kenapa dua-duanya pointer?
- Increment() -> ubah data
- Get() -> konsisten dan efisien. 
```
#
#
### Sedikit tambahan cahya ✮⋆˙
```azure
Sebuah struct dapat memiliki banyak method receiver.

Example:
    type User struct {
        Name string
        Age int
    }
    
    // value receiver
    func (u User) Info() string {
        return u.Name
    }
    
    // pointer receiver
    func (u *User) Birthday() {
        u.Age++
    }
    
    // pointer receiver
    func (u *User) Rename(name string) {
        u.Name = name
    }

1. Info() => hanya baca data
2. Birthday => ubah data
3. Rename => ubah data

Jadi yang benar satu struct bisa punya banyak method,
dan method-method itu bisa memakai pointer receiver.
```
