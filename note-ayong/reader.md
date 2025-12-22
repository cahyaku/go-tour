# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 7. Reader ✮⋆˙
#
### Reader ✮⋆˙
```azure
io.Reader adalah interface untuk "membaca data sedikit demi sedikit" dari suatu sumber.
Sumbernya bisa apa saja: (file, string, jaringan, zip, enkripsi, dan lain lain)

Analogi sehari-hari:
Bayangkan jika kamu minum air pakai sedotan
Air = data
Sedotan = io.Reader
Mulut = program kamu
Setiap sedotan = baca sebagian data
Kamu tidak minum semuanya sekaligus, tapi sedikit demi sedikit.

// IO.Reader
io.Reader adalah interface:
    type Reader interface {
        Read(p []byte) (n inte, err error)
    }
Tipe apapun yang punya method Read -> dianggap io.Reader.
    
    
// Method Read
Read(p []byte) (n inte, err error)
Maknanya:
• b []byte -> wadah untuk menaruh data
• n -> berapa byte yang berhasil dibaca
• err -> error (biasanya io. EOF kalau habis)
Data tidak dikembalikan langsung, tapi diisi ke slice.

Alasan menggunakan slice:
- slice bisa diisi ulang
- efisien (tidak bikin data baru terus)
- cocok untuk data besar

io.EOF = End Of File 
Artinya data sudah habis, tidak ada lagi yang bisa dibaca.
```
#
#
### Image (Gambar) ✮⋆˙
```azure
Dalam pemrograman Go, gambar tidak selalu PNG, JPG, atau GIF.
Bagi GO:
Sebuah gambar adalah sesuatu yang bisa menjawab pertanyaan:
- ukurannya berapa?
- warna pixel di posisi tertentu apa?
Iutlah sebabnya image.Ingae dibuat sebagai interface, bukan struct.

Interfaceimage.Image
    type Image interface {
        ColorModel() color.Model
        Bounds() Rectangle
        At(x, y int) color.Color
    }
Artinya: Semua tipe yang punya tiga method ini -> dianggap sebagai Image oleh Go.
```