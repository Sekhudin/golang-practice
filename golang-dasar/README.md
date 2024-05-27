# Golang Dasar

## Inisiasi Projek
Untuk memulai projek Go harus diawali dengan inisiasi *module* dengan perintah di bawah ini:
```bash
# nama module biasanya sesuai dengan nama folder
go mod init [nama-module]
```

## Projek `Hello World`
- Buat *file* dengan nama `hello-world.go`
- Sintaks
  ```go
  // nama paket
  package main
  // import library fmt
  import "fmt"

  // fungsi main >> mencetak "hello world"
  func main() {
  	fmt.Println("Hello World!")
  }
  ```

  Fungsi main adalah fungsi dimana projek akan  dijalankan (selayaknya method static `main` pada `Java`)
- Build (*Compile*) program Go dengan menjalankan perintah berikut pada terminal:
  ```bash
  go build
  ```
  Ketika perintah tersebut dijalankan maka akan menghasilkan *file* baru hasil kompilasi, biasanya sesuai dengan nama file-nya contoh: `hello-world.exe`(jika pada windows)
- Jalankan program dengan mengeksekusi file hasil kompilasi tersebut.<br>
Pada termninal jalankan (windows):
  ```bash
  #out >> "hello world!"
  ./hello-world.exe
  ```
- Menjalankan tanpa proses kompilasi. <br>
Untuk menjalankan program tanpa kompilasi (build), maka gunakan perintah ini pada terminal:
  ```bash
  go run [nama-file]
  ```
  Maka program Go akan dijalankan tanpa menghasilkan *file* ekstensi `.exe`. Ini hanya digunakan pada proses **Development**.

## Mutiple Main Function
- Pada Go, *function* bersifat ***unique***, artinya tidak boleh ada nama fungsi yang sama pada satu module/projek.
- Oleh karena itu, misalkan kita memiliki projek yang terdiri dari beberapa *file* maka tidak boleh ada *function* `main` di lebih dari 1 *file*. Jika tidak, **maka akan terjadi duplikat dan program tidak bisa dikompilasi**.
- Akan tetapi tetap bisa dijalankan *file* per-*file* (hanya untuk proses belajar saja).

## Tipe Data Number
### *Integer* (Bilangan Bulat)
Baca selengkapnya, [`integer`](https://www.w3schools.com/go/go_integer_data_type.php) pada Go.

### *Floating Point* (Bilangan Desimal)
Baca selengkapnya, [`float`](https://www.w3schools.com/go/go_float_data_type.php) pada Go.

### Alias
Adalah tipe data *number* alias (nama lain), misalnya:
- type `byte` adalah alias untuk type `uint8`
- type `rune` adalah alias untuk type `int32`
- type `int` adalah alias untuk type `Minimal int32`
- type `uint` adalah alias untuk type `Minimal uint32`

## Tipe Data Boolean
- Memiliki 2 nilai yaitu benar dan salah (`true` dan `false`).
- Direpresentasikan dengan kata kunci `bool`.

## Tipe data String
- Adalah kumpulan karakter.
- Jumlah karakter `string` pada Go bisa nol hingga tak hingga.
- Direpresentasikan dengan kata kunci `string`.
- Nilai `string` selalu diawali dan diakhiri dengan karakter petik dua. contoh:
  ```go
  var name string = "john doe"
  ```
- Fungsi pada `string`<br>
  `len("string")` untuk mengetahui panjang karakter.
  `"string"[n]` untuk mendapatkan karakter ke-n.

## Variabel pada Go
- Ketika menggunakan kata kunci `var` maka wajib langsung menginisiasikan nilainya. Contoh:
  ```go
  // var varname type = [value]
  var name string = "doe";
  // untuk mengubahnya,
  name = "john"
  ```
- Agar tidak menggunakan `var` maka gunakan tanda `:=`.<br>
Contoh:
  ```go
  //inisiasi awal
  name := "john doe"
  // untuk mengubahnya
  name = "doe john"
  ```
Baca selengkapnya, [variabel](https://www.w3schools.com/go/go_variables.php).

## Constant
- Constant adalah variabel yang tidak bisa diubah lagi setelah pertama kali diberi nilai.
- Mirip dengan variabel, yang membedakan hanyalah pada penggunaan kata kunci-nya. Pada constant menggunkan kata kunci `const` bukan `var`.
- Saat pembuatan constant kita wajib langsung menginisialisasikan data. Contoh:
  ```go
  // const CONSTNAME type = value
  const name string = "john doe"

  //atau dengan multiple constant
  const (
    firstName string = "john"
    lastName = "doe"
  )
  ```

Baca selengkapnya, [constant](https://www.w3schools.com/go/go_constants.php).

## Konversi Tipe Data
Pada beberapa kasus kita perlu mengkonversikan tipe data satu ke tipe data lain, misalnya konvesi dari `int32` ke `int64`.<br>
COntoh:
```go
// nilai awal
var nilai int32 = 32768
// setelah konversi
var nilai2 int64 = int64(nilai)
fmt.Println(nilai, nilai2)
// out >> 32768 32768
```
Perlu diperhatikan, ketika konversi nilai **harus memperhatikan cakupan data** dari tipe-nya.
```go
// nilai awal
var nilai int32 = 32768
// setelah konversi
var nilai3 int16 = int16(nilai)
// Terjadi number overflow, nilai melebihi cakupan dari int16
fmt.Println(nilai, nilai3)
// out >> 32768 -32768
```
**Konversi data String**
Misalkan ada sebuah kode seperti ini,
```go
	var name string = "john"
	var char = name[1];

	fmt.Println(char)
  // out >> 111
```
Kenapa keluaran-nya `111`? hal tersebut terjadi karena variabel `char` bertipe `byte` (alias dari `uint8`). Oleh karena itu, kita perlu meng-konversikan variabel `char` ke tipe data `string`, caranya:
```go
  // type string
	var name string = "john"
  // type byte
	var char = name[1];
  // dikonversikan
  var str string = string(char)
	fmt.Println(str)
  // out >> "o"
```

## *Type Declarations*
- *Type declarations* adalah kemampuan untuk membuat ulang tipe data baru dari tipe data yang sudah ada.
- Digunakan untuk membuat alias dari suatu tipe agar lebih mudah dimenegerti.<br>
Contoh:
```go
// type declarations
type PhoneNumber string
type NIP string
type Age int

// penggunaan
var phoneNumber PhoneNumber = "082222222222"
var nip NIP = "00000000 000000 0 000"
var age Age = 22
```

## Operasi Matematika
Baca selengkanya, [*arithmetic*](https://www.w3schools.com/go/go_arithmetic_operators.php), [*assignment*](https://www.w3schools.com/go/go_assignment_operators.php), [*comparison*](https://www.w3schools.com/go/go_comparison_operators.php), [logika](https://www.w3schools.com/go/go_logical_operators.php), [*bitwise*](https://www.w3schools.com/go/go_bitwise_operators.php).

## Tipe Data Array (Larik)
- `Array` adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama.
- Perlu menentukan jumlah data yang dapat ditampung.
- Daya tampung tidak dapat bertambah setelah `array` dibuat.
- Index `array` dimulai dari `0`.<br>
Contoh:
  ```go
  // var ARRAYNAME [n] type
  var names [2] string
  names[0] = "john"
  names[1] = "doe"

  fmt.Println(names)
  // out >> [john doe]

  // mengambil data ke-n
  fmt.Println(names[0])
  // out >> "john"
  ```
  Pembuatan `array` juga bisa dilakukan dengan langsung mengisi nilainya. <br>
  Contoh:
  ```go
  // var ARRAYNAME = [n] type{...values}
  var nilai = [2] int {1, 2}

  // atau

  // var ARRAYNAME = [...] type{...values}
  var nilai = [...] int {1, 3, 4, 5}
  // panjang array akan otomatis di set-4, sesuai banyak nilai yang di-inisiasikan diawal.
  ```
  **Catatan**: pada Go, tidak ada operasi hapus atau tambah item pada array karena panjangnya akan tetap, tidak bertambah dan tidak berkurang. sebagai ganti dari operasi hapus, kita bisa mngubah nilai index tertentu dengan nilai kosong (misal `""` atau `0`).

Baca selengkapnya, [array](https://www.w3schools.com/go/go_arrays.php).

## Tipe Data Slice
- Slice adalah **potongan dari data array**.
- Tipe data ini mirip array, yang membedakan hanya **slice berukuran dinamis**.
- Slice dan array **selalu terkoneksi**.
- Slice mengakses seluruh atau sebagian data di array yang menjadi koneksinya.

### Detail tipe data slice
- Slice memiliki 3 data yaitu **poninter, length** dan **capacity**.
- Pointer adalah penunjuk data pertama di array pada slice.
- Length adalah panjang slice.
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.<br>
  ```go
  // misal ada sebuah array bulan terdiri dari
  // SLICES[LOW:HIGH]
    "januari", // 0
    "februari", // 1
    "maret", // 2
    "april", // 3
    "mei", // 4 ----> pointer (penunjuk pertama) --> LOW
    "juni", // 5
    "juli", // 6
    "agustus", // 7 ---> HIGH
    "september", // 8
    "oktober", // 9
    "november", // 10
    "desember", // 11 ---> penunjuk terakhir
  
  // kemudian dari array tersebut kita buat slice
  // bulan[4:7] --> slice bulan
  // maka diketahui
  // pointer = 4 --> index ke-4 dari array (pointer pertama dari slices)
  // lenght = 3 ---> panjang slices (dari 4 ke 6 --> enam berasal dari (7-1))
  // capacity = 8 --> jarak penunjuk pertama ke penunjuk akhir (dari 4 ke 11)
  // angota slice = "mei" "juni" "juli"
  ```

  Contoh:
  ```go
  	names:= [...]string{"john", "miura" , "doe"}
    slicesName := names[1:2] // out >> ["miura"]
    slicesName1 := names[1:] // out >> ["miura" , "doe"]
    slicesName2 := names[:1] // out >> ["john"]
    slicesName3 := names[:] // out >> ["john", "miura" , "doe"]
  ```
  Deklarasi slices:
  ```go
  // var NAMASLICES = []type{...values}
  var s = []int{1, 2, 3}

  //atau
  
  // NAMASLICES := []type{..values}
  myslice := []int{1,2,3}
  ```
Baca selengkapnya, [slices](https://www.w3schools.com/go/go_slices.php).

## Tipe Data Map