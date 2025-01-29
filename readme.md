# Gin Backend

Ini adalah proyek latihan untuk membuat aplikasi CRUD menggunakan framework Gin dengan database MySQL.

> **Note**: Pastikan Anda sudah menginstall Go dan juga memiliki Postman atau aplikasi sejenis untuk melakukan testing API.

## Instalasi

1. Clone repositori ini:
    ```bash
    git clone https://github.com/kamachiii/GDG-golang-gin.git
    ```
2. Masuk ke direktori proyek:
    ```bash
    cd GDG-golang-gin
    ```
3. Instal dependensi:
    ```bash
    go mod tidy
    ```

## Menjalankan Aplikasi

Untuk menjalankan aplikasi, gunakan perintah berikut:
```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8080`.

## Struktur Proyek

- `database.sql`: Berisi skrip SQL yang digunakan untuk membuat dan menginisialisasi database yang diperlukan oleh aplikasi.
- `go.mod`: Berisi informasi tentang modul Go dan dependensi yang digunakan dalam proyek.
- `go.sum`: Berisi checksum dari setiap dependensi yang digunakan untuk memastikan integritas.
- `main.go`: File utama untuk menjalankan aplikasi.
- `readme.md`: Berisi dokumentasi proyek, termasuk cara instalasi, penggunaan, dan informasi lainnya.


## Konfigurasi

```json
{
    "port": "8080",
    "database": {
        "host": "localhost",
        "port": "3306",
        "user": "root",
        "password": "",
        "dbname": "gdg_backend_go_gin"
    }
}
```

## Rute API

Berikut adalah beberapa rute API yang tersedia:
- `GET /`: Mengecek status db connection
- `GET /albums`: Mendapatkan semua data.
- `GET /albums/:id`: Mendapatkan data berdasarkan id.
- `POST /album`: Membuat data baru.
- `PUT /album/:id`: Memperbarui data berdasarkan ID.
- `DELETE /album/:id`: Menghapus data berdasarkan ID.

## Format POST dan PUT

Berikut adalah format raw untuk menambah dan mengupdate data:
```json
{
    "id": "5",
    "title": "Si Jabeng",
    "price": 89.99
}
```

