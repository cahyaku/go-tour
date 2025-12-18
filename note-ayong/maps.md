# ⋆.˚🦋༘⋆ GO TOUR CAHYA CUTY ᯓ★ ּ ֶָ֢.
#
#
## 4. Maps ✮⋆˙
#
### Create Map ✮⋆˙
```azure
Sebuah map memetakan sebuah kunci (key) dengan nilainya.

Nilai kosong dari sebuah map adalah nil. Sebuah map yang nil tidak memiliki kunci,
tidak juga dapat ditambahkan kunci baru.

Fungsi make mengembalikan sebuah map dengan tipenya, diinisialisasi dan siap untuk digunakan.

example: pada function createLocationMap()
```
#
#
### Inisialisasi Map ✮⋆˙
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
#
#
### Inisialisasi map lanjutan ✮⋆˙
```azure
Jika level teratas dari tipe hanya nama tipe,
maka bisa menghilangkannya dari inisialisasi elemen map.

Jadi ini lebih singkat, jika tipe key sama dengan value.
ex.
    var m = map[string]Vertex{
        "Bell Labs":{40.68433, -74.39967},
        "Google":{37.42202, -122.08408}},
    }
```
#
#
### Operasi Map (Mutating Maps) ✮⋆˙
```azure
1. Mengisi atau mengubah elemen dalam map.
    m[key] = elem
2. Mengambil elemen dari map.
    elem = m[key]
3. Menghapus elemen dari map. 
    delete(m,key)
4. Menguji apakah sebuah key ada dalam map.
    elem, ok = m[key]

example: pada function mapOperation()
```
#
#
### Nilai Fungsi ✮⋆˙
```azure
Fungsi adalah suatu nilai, fungsi dapat dikirimkan kemanapun seperti nilai lainnya.

Nilai fungsi bisa digunakan sebagai argumen pada fungsi lainnya dan sebagai nilai kembalian.

example: pada func compute()
```
#
#
### Fungsi closure ✮⋆˙
```azure
Fungsi pada Go bisa closure. 
Closure adalah sebuah nilai fungsi yang merujuk variabel dari blok fungsinya.
Fungsi closure bisa mengakses dan mengisi variabel yang dirujuk, 
dalam artian fungsi tersebut "terikat" ke variabel.

example: pada function runAccumulatorDemo()


Simplenya kita panggil fungsi blok luarnya, tapi fungsi blok luarnya juga bisa dipanggil sendiri.

example:
func taskCreator() func(string) {
    total := 0

    return func(title string) {
        total++
        fmt.Println("Task ke-", total, ":", title)
    }
}

// Panggil fungsi blok luarnya
addTask := taskCreator()

addTask("Belajar Go")
addTask("Ngopi")
addTask("Tidur")

// Output:
Task ke- 1 : Belajar Go
Task ke- 2 : Ngopi
Task ke- 3 : Tidur
```

