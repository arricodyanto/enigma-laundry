# challenge-godb

## Getting started

Sebuah aplikasi CLI sederhana yang dapat digunakan untuk mencatat transaksi pada
toko "Enigma Laundry"

## Clonning repository

Untuk menjalankan aplikasi ini, silakan clone repository ini ke local storage
device Anda dengan menggunakan perintah berikut.

```
https://git.enigmacamp.com/enigma-20/arrico-handyanto/challenge-godb.git
```

## Setup Database

Sebelum melakukan run pada project di repository ini, lakukan konfigurasi
database yang akan Anda gunakan di local. Perhatikan baris kode berikut yang ada
pada directory /db/connection.go

```
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "enigma_laundry"
)
```

Ubah value pada variable di atas sesuai dengan database yang Anda miliki. Jika
belum memiliki database, silakan buat database baru dengan nama "enigma_laundry"

## Run SQL

Untuk melakukan inisialisasi schema pada database, Anda dapat menjalankan sql
dengan Querying Tool dengan menggunakan script pada directory berikut

```
- /db/sql/DDL.sql
- /db/sql/DML.sql
```

## Run The App

Project Anda sudah siap digunakan. Jalankan perintah berikut untuk memulai.

```
go run .
```

## Manuals

Berikut adalah beberapa fitur yang bisa digunakan pada aplikasi ini, yang akan
muncul di menu utama setelah aplikasi dijalankan.

### Customer Management

Menu ini dapat digunakan untuk melakukan management terhadap table master
mst_customer. Anda bisa melakukan CREATE, READ, UPDATE, dan DELETE terhadap
record customer yang tersimpan pada database. Fitur DELETE Custmer sudah
ditambahkan CASCADE ON DELETE, sehingga data pada table lain yang terelasi
dengan customer terkait yang sedang dihapus juga akan ikut hilang, yaitu data
pada tabel trx_bill_detail, dan trx_bill.

### Service Management

Menu ini dapat digunakan untuk melakukan managament terhadap table master
mst_service. Tabel ini digunakan untuk menyimpan jenis layanan yang dapat
diberikan oleh toko 'Enigma Laundry.' Sama halnya dengan menu pertama, Anda bisa
melakukan CREATE, READ, UPDATE, dan DELETe terhadap record service yang
tersimpan pada database. Fitur ini juga dilengkapi dengan CASCADE ON DELETE
sehingga akan memberikan dampak pada data relasi terkait pada tabel lain. Ketika
Anda menghapus salah satu service yang telah digunakan pada tabel trx_bill dan
trx_bill_detail, maka hal ini juga akan mengupdate dan mengkalkulasi ulang harga
yang telah tersimpan.

### Show All Transaction

Pada menu ini Anda bisa melihat semua transaksi yang terekam pada database. Anda
juga bisa melihat detail dan layanan apa saja yang digunakan dalam satu buah
transaksi.

### Insert New Transaction

Menu ini memungkinkan Anda untuk mencatat setiap transaksi baru di toko 'Enigma
Laundry.' Pada menu ini Anda bisa melakukan perulangan detail layanan yang
digunakan pada sebuah transaksi. Fitur ini hanya bisa digunakan untuk customer
yang telah terdaftar.

### Exit

Exit Application.

## Credits

[Arrico Handyanto (Vercel)](https://arricohandyanto.vercel.app) -
[Arrico Handyanto (LinkedIn)](https://www.linkedin.com/in/arricohandyanto/)
