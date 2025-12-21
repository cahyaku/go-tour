# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★ ּ ֶָ֢.
#
#
## 6. 𒅒𒈔𒅒 Interfaces °:•.𖹭
#
### Interfaces ✮⋆˙
```azure
Sebuah type interface didefinisikan sebagai sebuah kumpulan method penanda.

Nilai dari tipe interface dapat menyimpan nilai apapun yang mengimplementasikan method tesebut.

Sebagai contoh yang sederhana:
    // Buat interface
    type Speaker interface {
        Speak()
    }
    Artinya: siapa pun yang punya method Speak() dianggap sebagai Speaker.
    
    // Buat struct
    type Human struct {}
    func (h Human) Speak() {
        fmt.Println("Hello!")
    }
    
    // Buat struct lain
    type Dog struct {}
    func (d Dog) Speak() {
        fmt.Println("Goukk Goukk!")
    }
    
    // Pakai interface
    func MakeSound(s Speaker) {
    	s.Speak()
    }
        
    // Inisiasi
    h := Human{}
    d := Dog{}
    MakeSound(h)
    // Output: Hello!

    MakeSound(d)
    // Output: Goukk Goukk!
```
#
#
### Interfaces are implemented implicitly (Interface dipenuhi secara implisit) ✮⋆˙
```azure
Sebuah tipe mengimplementasikan sebuah interface dengan meningimplemetasikan method-methodnya.
Tidak ada deklarasi eksplisit; tidak ada perintah "implements".

Interface implisit memisahkan definisi dari sebuah interface dari implementasinya,
yang bisa saja muncul dalam paket manapun tanpa adanya penataan sebelumnya.

Umumnya, yang paling sering terjadi adalah interface punya satu method. 

Contoh paling sederhana:
    type Greeter interface {
        Greet() string
    }
    type Human struct {
        Name string
    }
    func (h Human) Greet() string {
        return "Hello, " + h.Name + "!"
    }
    
Kesimpulan:
- Interface adalah kumpulan method penanda.
- Nilai dari tipe interface dapat menyimpan nilai apapun yang mengimplementasikan method tesebut. 

Interface dapat menggunakan method dari package lain tanpa perlu menyediakan implementasi sendiri.
```
#
#
### Interface values (Nilai interface) ✮⋆˙
```azure
Isi interface dapat dibayangkan sebagai sebuah pasangan nilai sebuah tipe:

	(nilai, tipe)

Sebuah intreface mengandung sebuah nilai dari tipe tertentu.
Memanggil suatu method terhadap suatu interface akan mengeksekusi method
dengan nama yang sama pada tipe yang dipegangnya.

1. Inti:
- Interface menyimpan dua hal sekaligus (nilai, tipe)
- nilai = data aslisnya
- tipe = tipe konkret dari data itu
- walaupun variablenya bertipe interface, Go tetap ingat tipe aslinya.

2. Contoh sederhana Interface
    // interface
    type Greeter interface {
        Greet()
    }
    
    // struct
    type Person struct {
        Name string
    }

    // method
    func (p Person) Greet() {
        fmt.Println("Hallo,", p.NAme)
    }

3. Menyimpan struct ke interface
        var g Greeter
        p := Person {Name: "Cahya"}
        g = p
        
        // Sekarang isi ge adalah:
        (nilali = Person{Name: "Cahya"}, tipe = Person)
        // Walaupun g bertipe Greeter, isis aslinya adalah Person.
        
        // Memangil method lewat interface
        g.Greet()
        
        // Yang terjadi:
        - Go melihat g (interface)
        - Go cek isinya (Person{Name:"Cahya", Person})
        - Go menjalankan Person.Greet()
        
        // Hasil outpunya:
        Halo, Cahya

Method yang dijalankan tergantung tipe yang disimpan di dalam interface.


Contoh jika ada tipe lain, misalnya Dog:
    
type Dog struct

func (d Dog) Greet() {
    fmt.Println("Guk guk!") 
}

var g Greeter
g = Person{Name: "Cahya"}
g.Greet() // Person.Greet()

g. = Dog{}
g.Greet() // Dog.Greet()
```
#
#
### Interface values with nil underlying values (Interface dengan nilai nil) ✮⋆˙
```azure
- Jika nilai sebenarnya dari interface adalah nil, method akan dipanggil dengan receiver bernilai nil.
- Hal tersebut mengakibatkan null pointger exception atau kesahalahan karena pointer bernilai kosong.

Dalam go ada dua jenis nilai nil:
1. Interface = nil
    (nilai = nil, tipe = nil)
2. Interface berisi nilai nil
    (nilai nil, tipe = *T)

Keterangan:
1. Interface berisi nilai nil = methodnya tetap bisa dipanggil
2. Interface nil = method TIDAK bisa dipanggil (panic)

Example:

type Printer interface {
    Print()
}

type Message struct {
	Text string
}

func (m *Message) Print() {
if m == nil {
		fmt.Println("Receiver nil, tapi method tetap jalan")
    return
        }
	fmt.Println(m.Text)
}

func main() {
	// 1️⃣ Interface nil
	var p1 Printer
	fmt.Println("p1 == nil:", p1 == nil)
	// p1.Print() // ❌ panic jika dibuka

	// 2️⃣ Interface berisi nilai nil
	var msg *Message = nil
	var p2 Printer = msg

	fmt.Println("p2 == nil:", p2 == nil)
	p2.Print()
}

Hasil output:
p1 == nil: true
p2 == nil: false
Receiver nil, tapi method tetap jalan
```
#
#
### Nil interface values (isi interface dengan nil) ✮⋆˙
```azure
- Sebuah isi interface yang nil tidak mengandung nilai dan tipe konkrit.
- Memanggil sebuah method pada sebuah interface nil akan mengakibatkan
  kesalahan run-time karena tidak ada tipe di dalam interface yang mengindikasikan
  method konkrit yang akan dipanggil.

Intinya:Interface yang benar-benar kosong tidak ada nilai dan tipe,
Go tidak tahu method mana yang harus dipanggil.


example:
type Greeter interface {
	Greet()
}

func main() {
	var greeter Greeter

	fmt.Println(greeter == nil) // true

	greeter.Greet() // ❌ runtime panic
}
```
#
#
### The empty interface (Interface kosong) ✮⋆˙
```azure
- Interface kosong adalah interface tanpa method sama sekali.
- Tipe interface yang tidak memiliki method dikenal juga termasuk interface kosong.

example:
    interface{}

Semua tipe bisa disimpan dalam interface

Penjelasan:
1. Interface biasa -> kontrak
   Jadi: interface biasa -> harus bisa Drive()
       
2. Interface kosong -> tanpa kontrak sama sekali
   Jadi interface kosong -> "Terserah kamu saja"
   
Example:
// Karena interface kosong, jadi bisa diisi berbagao tipe berbeda
var anything interface{}

func main() {
	printValue("Go")
	printValue(100)
	printValue(true)
}

func printValue(v interface{}) {
	fmt.Printf("Nilai: %v | Tipe: %T\n", v, v)
}

Hasil output:
Nilai: Go | Tipe: string
Nilai: 100 | Tipe: int
Nilai: true | Tipe: bool
```
#
#
### Kapan iinterface kosong sebaiknya dipakai? ✮⋆˙
```azure
- Saat tipe belum diketahui
- Saat membuat fungsi umum
- Saat menangani data dinamis

Contoh penggunaan nyata:
- fmt.Print
- json.Unmarshal
- map[string]interface{}

Ringkasan singkat interface kosong:
✔ interface{} = interface tanpa method
✔ Bisa menyimpan nilai dari tipe apa pun
✔ Digunakan saat tipe tidak diketahui
✔ Perlu type assertion untuk mengolah nilai
✔ Jangan dipakai berlebihan
```

