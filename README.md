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

## Credits

[Arrico Handyanto (Vercel)](https://arricohandyanto.vercel.app) -
[Arrico Handyanto (LinkedIn)](https://www.linkedin.com/in/arricohandyanto/)
