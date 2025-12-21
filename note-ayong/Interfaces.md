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
### Interface values (Nilai interface)
```azure
Isi interface dapat dibayangkan sebagai sebuah pasangan nilai sebuah tipe:
	(nilai, tipe)
Sebuah intreface mengandung sebuah nilai dari tipe tertentu.
Memanggil suatu method terhadap suatu interface akan mengeksekusi method
dengan nama yang sama pada tipe yang dipegangnya.
```