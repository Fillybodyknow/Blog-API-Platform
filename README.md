# 📘 Blog API Platform

**Blog API Platform** คือระบบ API สำหรับสร้างและจัดการ Blog ที่มีฟีเจอร์ครบครัน ไม่ว่าจะเป็นระบบผู้ใช้งาน (Auth), ยืนยันตัวตนด้วย OTP, การกด Like และการแสดงความคิดเห็น  
พัฒนาโดยใช้ **Go (Gin)**, **MongoDB**, **Docker** และมีเอกสาร API ผ่าน **Swagger UI**

---

## 🚀 Features

- ✅ ลงทะเบียน / เข้าสู่ระบบ (Register / Login)
- 🔐 Authentication ด้วย JWT
- 🔑 ยืนยันตัวตนด้วย OTP (OTP Verification)
- 📝 สร้าง / แก้ไข / ลบบทความ (CRUD Posts)
- ❤️ กด Like / ยกเลิก Like
- 💬 แสดงความคิดเห็น / แก้ไข / ลบ
- 📖 เอกสาร API ด้วย Swagger UI

---

## 🛠️ Tech Stack

| Technology          | Description              |
|---------------------|--------------------------|
| **Go** (Gin)        | Backend API Framework    |
| **MongoDB**         | ฐานข้อมูล NoSQL         |
| **Docker**          | Containerization         |
| **Swagger**         | API Documentation        |
| **JWT**             | Token-based Authentication |
| **godotenv**        | จัดการ environment file |

---

## 📦 วิธีติดตั้งและใช้งาน (Local)

### ✅ สิ่งที่ต้องติดตั้งก่อน

- [Docker Desktop](https://www.docker.com/products/docker-desktop)

### 🧪 เริ่มต้นใช้งาน

```bash
# 1. Clone โปรเจกต์
git clone https://github.com/Fillybodyknow/blog-api.git
cd blog-api

# 2. สร้างไฟล์ .env จาก template
cp .env.example .env

# 3. แก้ไขค่า .env (ดูรายละเอียดการตั้งค่า SMTP ด้านล่าง)

# 4. รันโปรเจกต์ด้วย Docker
docker-compose up --build
```

## 📧 การตั้งค่า SMTP สำหรับส่ง OTP

เพื่อให้ระบบสามารถส่งอีเมล OTP ได้ คุณต้องตั้งค่า SMTP โดยแนะนำให้ใช้ Gmail ผ่าน **App Password** (ไม่ใช่รหัสผ่านบัญชีโดยตรง)

### วิธีสร้าง SMTP Email และ Password จาก Gmail:

1. ไปที่หน้า [Google Account Security](https://myaccount.google.com/security)
2. เปิดใช้งาน **2-Step Verification (2FA)** หากยังไม่ได้เปิด
3. หลังจากเปิด 2FA แล้ว จะมีเมนู **App Passwords** ปรากฏขึ้น หากยังไม่เคยสร้าง App Passwords จะไม่แสดงให้เห็น ให้เข้าลิงค์นี้โดยตรงแทน [App Passwords](https://myaccount.google.com/apppasswords)
4. เข้าไปที่ **App Passwords**
5. สร้างรหัสใหม่ โดย:
   - App: เลือก "Other (Custom name)" → พิมพ์ `Blog API`
   - กด "Generate"
6. จะได้ **App Password 16 หลัก** เช่น `abcd efgh ijkl mnop`

### เพิ่มค่าที่ได้ลงในไฟล์ `.env`

```env
SMTP_EMAIL=your-email@gmail.com
SMTP_PASS=abcdefghijklmnop
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

## 🌐 API Documentation
Swagger UI: http://localhost:8080/swagger/index.html
คุณสามารถดูรายละเอียด endpoint, request/response และทดสอบ API ได้ที่นี่

## 🛡 วิธีใช้ Token ใน Swagger UI
1. เรียก /api/auth/login เพื่อเข้าสู่ระบบ → ได้ token จาก response
2. กดปุ่ม Authorize บน Swagger UI
3. วาง Token โดยมี Bearer นำหน้า

### ตัวอย่าง:

```Authorize
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
หากไม่เติมคำว่า Bearer จะไม่สามารถเข้าถึง endpoint ที่ต้องยืนยันตัวตนได้

📄 License
MIT License © 2025 Fillybodyknow
