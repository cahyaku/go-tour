# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 2. Flow Control ✮⋆˙
# 
### For Loop ✮⋆˙
```azure
Go hanya punya satu perintah perulangan yaitu "for".
Example:
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
#
#
### For Continued ✮⋆˙
```azure
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```
#
#
### For as While ✮⋆˙
```azure
for condition {
    // kode yang dijalankan selama kondisi benar
}

example:
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

Untuk menghentikan perulangan sama aja pakai "break" atau "return".
```
# 
#
### IF Statement ✮⋆˙
```azure
if kondisi {
    // kode yang dijalankan jika kondisi benar
} else {
    // kode yang dijalankan jika kondisi salah
}

example:
func main() {
    x := 10
    if x%2 == 0 {
        fmt.Println("Genap")
    } else {
        fmt.Println("Ganjil")
    }
}
```
#
#
### IF with a short statement ✮⋆˙
```azure
if statement; kondisi {
    // kode yang dijalankan jika kondisi benar
} else {
    // kode yang dijalankan jika kondisi salah
}

example:
func main() {
    if x := compute(); x < 0 {
        fmt.Println("Negatif")
    } else {
        fmt.Println("Positif atau Nol")
    }
```





