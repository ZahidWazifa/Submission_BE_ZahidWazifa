# Dokumentasi REST API Toko Online

## Deskripsi
Project ini merupakan implementasi REST API untuk sistem toko online menggunakan Go dengan framework GORM. 
API ini dilengkapi dengan sistem autentikasi JWT untuk keamanan endpoint.

## Fitur Utama
- Autentikasi menggunakan JWT Token
- CRUD Produk
- Validasi Request
- Dokumentasi API Lengkap

## Struktur Project
```
Task3/
├── config/
│   └── database.go
├── controllers/
│   ├── auth_controller.go
│   └── product_controller.go
├── middlewares/
│   └── auth_middleware.go
├── models/
│   ├── user.go
│   └── product.go
├── routes/
│   └── routes.go
└── main.go
```

## Endpoint API

### Autentikasi

#### Register User
- Method: POST
- Endpoint: `/api/auth/register`
- Request Body:
```json
{
    "username": "string",
    "password": "string",
    "email": "string"
}
```
- Response Success:
```json
{
    "status": "success",
    "message": "User berhasil didaftarkan"
}
```

#### Login
- Method: POST
- Endpoint: `/api/auth/login`
- Request Body:
```json
{
    "username": "string",
    "password": "string"
}
```
- Response Success:
```json
{
    "status": "success",
    "token": "jwt_token_string"
}
```

### Produk

#### Get All Products
- Method: GET
- Endpoint: `/api/products`
- Headers: 
  - Authorization: Bearer {token}
- Response Success:
```json
{
    "status": "success",
    "data": [
        {
            "id": 1,
            "name": "string",
            "price": number,
            "stock": number
        }
    ]
}
```

#### Create Product
- Method: POST
- Endpoint: `/api/products`
- Headers: 
  - Authorization: Bearer {token}
- Request Body:
```json
{
    "name": "string",
    "price": number,
    "stock": number
}
```
- Response Success:
```json
{
    "status": "success",
    "data": {
        "id": 1,
        "name": "string",
        "price": number,
        "stock": number
    }
}
```

## Instalasi & Menjalankan Aplikasi

1. Clone repository
2. Install dependensi:
```bash
go mod tidy
```
3. Sesuaikan konfigurasi database di config/database.go
4. Jalankan aplikasi:
```bash
go run main.go
```

## Teknologi yang Digunakan
- Go 1.16+
- GORM
- JWT-Go
- Gin Framework
- MySQL/PostgreSQL

## Keamanan
- Implementasi JWT untuk autentikasi
- Password hashing menggunakan bcrypt
- Validasi input request
- Middleware authentication untuk protected endpoints

## Referensi
- [GORM Documentation](https://gorm.io/docs/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [JWT-Go Documentation](https://github.com/golang-jwt/jwt)
