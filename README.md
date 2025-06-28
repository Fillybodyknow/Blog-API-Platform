# 📘 Blog API Platform

Blog API Platform คือระบบสำหรับสร้างและจัดการ Blog ที่รองรับระบบผู้ใช้งาน (Auth), ยืนยันตัวตนด้วย OTP, การกด Like และ Comment บนบทความ  
พัฒนาโดยใช้ **Go (Gin)**, **MongoDB** และ **Docker** พร้อมเอกสาร API ผ่าน **Swagger UI**

---

## 🚀 Features

- ✅ ลงทะเบียน / เข้าสู่ระบบ (Register / Login)
- 🔒 ยืนยันตัวตนด้วย OTP (OTP Verification)
- 📝 สร้าง / แก้ไข / ลบ บทความ (CRUD Posts)
- ❤️ กด Like / ยกเลิก Like
- 💬 แสดงความคิดเห็น / แก้ไข / ลบ
- 🔍 API Document ด้วย Swagger UI

---

## 🛠️ Tech Stack

| Technology | Description        |
|------------|--------------------|
| Go         | Backend API (Gin) |
| MongoDB    | ฐานข้อมูล NoSQL   |
| Docker     | Containerization   |
| Swagger    | API Documentation  |

---

## 📦 วิธีติดตั้งและใช้งาน (Local)

### ✅ สิ่งที่ต้องติดตั้งก่อน

- [Docker Desktop](https://www.docker.com/products/docker-desktop)

### 🧪 เริ่มใช้งานโปรเจกต์

```bash
# 1. Clone โปรเจกต์
git clone https://github.com/Fillybodyknow/blog-api.git
cd blog-api

# 2. สร้างไฟล์ .env
cp .env.example .env
# หรือแก้ค่าใน .env ตามต้องการ

# 3. สั่งรันด้วย Docker
docker-compose up --build

🌐 API Documentation
Swagger UI: http://localhost:8080/swagger/index.html
คุณสามารถดูรายละเอียด endpoint, request/response และทดสอบ API ได้ที่นี่

📄 License
MIT License © 2025 Fillybodyknow
