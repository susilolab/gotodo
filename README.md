# gotodo
Aplikasi todo sederhana menggunakan framework gofiber


## Setup

Clone repo ini dan jalankan 
```
go mod tidy
```

## Auto reload

Agar setiap kali terjadi perubahan kode kita perlu melakukan `go build` berulang ulang dapat menggunakan 
aplikasi `air`

```
go get -u github.com/cosmtrek/air
cd gotodo
air init
air
```
