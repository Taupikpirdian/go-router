[WHAT_TASK]
* Melanjutkan Assignment 1, tambahkan relasi Article Category
  1. Dengan Category mempunyai multiple Language
  2. Implement router dengan method Get untuk menampilkan data Article (All). bisa di filter berdasarkan parameter Category slug
  3. Implement router dengan method Get untuk menampilkan data Article berdasarkan Article slugnya dan ditampilkan sesuai dengan bahasa yang dipilih
  contoh url : localhost/slug-title-article/en atau localhost/slug-title-article?lang=en
  1. Implement poin diatas menggunakan redis cache di setiap use case nya
  2. Mapping Response ke bentuk JSON

[HOW_TO_RUN]
* Installation
  - remove go.mod and go.sum
  - go mod init
  - go mod tidy
  - go mod vendor
  - import db to database (database sudah dimasukan ke dalam zip, salt_academy_exam_2.sql)
  - go run main.go

[DOCS]
* For Install Package
  - Go get github.com/go-redis/redis/v8
  - Go get -u github.com/gorilla/mux

[CURL]
  - curl --location --request GET 'http://localhost:8080/article/en?category=food'
  - curl --location --request GET 'http://localhost:8080/article/detail/the-bad-impact-of-instant-noodles'