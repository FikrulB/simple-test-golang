# SuperIndo Product API

## 📌 Deskripsi
SuperIndo Product API adalah RESTful API berbasis Golang yang digunakan untuk mengelola data produk di SuperIndo. API ini menyediakan fitur untuk menambahkan, menampilkan, mencari, memfilter, dan mengurutkan data produk.

---

## 🚀 Fitur
- **Menambahkan produk baru**
- **Menampilkan daftar produk**
- **Mencari produk berdasarkan ID atau nama**
- **Memfilter produk berdasarkan kategori** (Sayuran, Protein, Buah, Snack)
- **Sorting produk berdasarkan tanggal, harga, dan nama**

---

## 🛠️ Teknologi yang Digunakan
- **Bahasa**: Golang
- **Database**: PostgreSQL (Docker) + Migration & Seeder
- **Cache**: Redis
- **Dependency Injection**: Wire
- **Docker**: Digunakan untuk menjalankan seluruh aplikasi

## 🏗️ Instalasi dan Menjalankan API dengan Docker

### **1. Clone Repository**
```sh
git clone https://github.com/FikrulB/simple-test-golang.git
```

### **2. Jalankan Docker Compose**
```sh
docker-compose up --build -d
```

### **3. Jalankan Database Migration & Seeder dalam Container**
```sh
docker exec -it golang_app ./main migrate
```
```sh
docker exec -it golang_app ./main seeder
```

### **4. API akan berjalan di**
```
http://localhost:5000
```

---

## 📌 Endpoint API

### **1. Tambah Produk**
- **Endpoint**: `POST /product`
- **Request Body (JSON)**:
```json
{
  "code": "LPT-333",
  "name": "Test 11",
  "price": 4000,
  "category_id": 2
}
```
- **Response (201 OK)**:
```json
{
  "code": 201,
  "message": "Product added successfully"
}
```

### **2. Tampilkan Semua Produk**
- **Endpoint**: `GET /product`
- **Response (200 OK)**:
```json
{
  "code": 200,
  "message": "Products retrieved successfully",
  "data": [
    {
      "code": "LPT-333",
      "name": "Test 11",
      "category": {
        "id": 2,
        "name": "Protein",
        "create_at": "2025-02-16T11:35:21.916182Z"
      },
      "price": 4000,
      "created_at": "2025-02-16T11:37:35.286597Z"
    }
  ]
}
```

### **3. Cari Produk berdasarkan Nama atau ID**
- **Endpoint**: `GET /product?name=Test 1`
- **Response (200 OK)**:
```json
{
  "code": 200,
  "message": "Products retrieved successfully",
  "data": [
    {
      "code": "LPT-333",
      "name": "Test 11",
      "category": {
        "id": 2,
        "name": "Protein",
        "create_at": "2025-02-16T11:35:21.916182Z"
      },
      "price": 4000,
      "created_at": "2025-02-16T11:37:35.286597Z"
    }
  ]
}
```

### **4. Filter Produk berdasarkan Kategori**
- **Endpoint**: `GET /product?category_id=2`

### **5. Sorting Produk**
- **Endpoint**: `GET /product?sortBy=created_at&order=asc`

---

## 📬 Kontak
Jika ada pertanyaan, silakan hubungi:
- **Nama**: [M. Fikrul Bachtiar]
- **Email**: [mfikrulb@gmail.com]
- **GitHub**: [https://github.com/FikrulB]