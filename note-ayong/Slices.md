# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 3. Slices ✮⋆˙
#
### Slices ✮⋆˙
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
```