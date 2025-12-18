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
2. Dapat mendeklarasikan method dengan receiver yang tipenya didefinisikan di paket yang sama/

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
