# Belajar Golang RESTFul

Belajar Golang RESTFul API ini merupakan tutorial sederhana bagi pemula yang ingin membuat API sederhana tanpa menggunakan Framework. Projek ini mencoba menggunakan metode *Clean Architecture* agar teman-teman lebih mudah untuk memahaminya. Projek ini boleh teman-teman pelajari sebagai dasar untuk mempelajari bagaimana kita membuat suatu API dalam bahasa Golang. Pada projek ini pun kita menjelaskan dari mulai persiapan dokumentasi API, library pendukung, persiapan database MySQL yang digunakan sampai uji coba (unit test) sudah kami jelaskan dengan lengkap.

Bagi kalian yang ingin belajar lebih mudah bisa kunjungi website kami [Santekno](https://www.santekno.com/) dan jika ingin belajar golang restful API bisa langsung ke topik [Belajar Golang RestFul](https://www.santekno.com/tutorial/golang/restful-api-dan-clean-architecture/).

## Cara Menjalankan Projek
Teman-teman bisa juga langsung mempelajari pada repository ini jika ingin langsung menjalankan programnya. Tetapi perlu disiapkan beberapa langkahnya seperti:
1. Clone repository ini dengan perintah
   ```bash
   git clone git@github.com:santekno/learn-golang-restful.git
   ```
2. Siapkan database yang dipakai yaitu MySQL, jika teman-teman sudah punya Docker, dari projek ini sudah disediakan juga `docker-compose.yml` untuk membuat MySQL Server menggunakan Docker.
   ```bash
   docker-compose up -d
   ```
3. Jika sudah tersedia MySQL Server, maka kita buat database `article` dengan table `articles`, bisa execute query ini pada DBMS yang kamu pakai
   ```sql
   create table articles(
      id integer primary key auto_increment,
      title varchar(255) not null,
      content varchar(255) not null,
      create_at datetime not null,
      update_at datetime not null
    ) engine = InnoDB;
   ```
4. Jalankan projek pada terminal dengan perintah
   ```bash
   go mod vendor
   make build
   ```
5. Sukses, projek sudah berjalan dan bisa kita coba API yang ada pada dokumentasinya.

## Aturan Penamaan Projek
* Nama folder harus menggunakan `dash` jika lebih dari dua kata. Contoh `middleware-chain`, `util-category` (Kebab case)
* Nama file golang menggunakan `underscore` jika lebih dari dua kata. Contoh `article_test.go`,`article_data.go` (Snake Case)
* Nama fungsi harus menggunakan `CamelCase`. Contoh: `GetAll()`, `GetArticleById()`
* Nama variabel harus menggunakan `CamelCase` dengan huruf pertama kecil. Contoh: `isSuccess`, `dataResult`, `results`
* Nama constant harus menggunakan `AllCaps` dan jika lebih dari satu kata tambahkan pemisahnya dengan `underscore`. Contoh: `HOSTNAME`, `POSTGRES_HOST`.

## Contributor
Maintenance By @santekno
