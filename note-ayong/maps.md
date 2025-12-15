# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★
#
#
## 4. Maps ✮⋆˙
#
### MAP ✮⋆˙
```azure
Sebuah map memetakan sebuah kunci (key) dengan nilainya.

Nilai kosong dari sebuah map adalah nil. Sebuah map yang nil tidak memiliki kunci,
tidak juga dapat ditambahkan kunci baru.

Fungsi make mengembalikan sebuah map dengan tipenya, diinisialisasi dan siap untuk digunakan.

example: pada function createLocationMap()
```
#
#
### Inisialisasi Maps
```azure
Map diinisialisasi seperti pada struct, tapi dengan menyertakan keynya.
Jadi harus berpasangan  dengan value nilai.

example:
    var m = map[string]Vertex{
        "Bell Labs": Vertex{
		    40.68433, -74.39967,
	},
        "Google": Vertex{
		    37.42202, -122.08408,
	},
}
```