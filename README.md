# 🚀 CRUD Golang & Vue 3 Application

A full-stack web application implementing **CRUD (Create, Read, Update, Delete)** operations built with **Golang** as the backend REST API and **Vue.js 3** as the frontend interactive user interface.

> 📚 **Catatan:** Proyek ini dibuat untuk keperluan latihan dan studi kasus pengembangan web fullstack berdasarkan materi tutorial dari [SantriKoding.com](https://santrikoding.com).

---

## 📋 Daftar Isi

- [Fitur Utama](#-fitur-utama)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Struktur Proyek](#-struktur-proyek)
- [Prasyarat Sistem](#-prasyarat-sistem)
- [Panduan Instalasi & Cara Jalankan](#-panduan-instalasi--cara-jalankan)
  - [1. Backend (Golang)](#1-backend-golang)
  - [2. Frontend (Vue 3)](#2-frontend-vue-3)
- [Dokumentasi API](#-dokumentasi-api)
- [Variabel Lingkungan (.env)](#-variabel-lingkungan-env)
- [Referensi & Kredit](#-referensi--kredit)
- [Kontribusi](#-kontribusi)
- [Lisensi](#-lisensi)

---

## ✨ Fitur Utama

- 🔄 **Full CRUD Operations**:
  - **Create**: Tambah data baru melalui form interaktif.
  - **Read**: Tampilkan daftar data dengan tampilan tabel/card yang rapi & fitur detail.
  - **Update**: Edit data yang sudah ada secara realtime.
  - **Delete**: Hapus data dengan konfirmasi aman.
- ⚡ **Fast & Lightweight Backend**: Dibangun menggunakan Golang untuk performa tinggi dan eksekusi cepat.
- 🎨 **Modern Reactive Frontend**: Dibangun dengan Vue 3 (Composition API) untuk UI yang responsif dan terstruktur.
- 🔌 **RESTful API**: Komunikasi antara frontend dan backend menggunakan JSON.
- 🌐 **CORS Handled**: Akses API lintas domain sudah dikonfigurasi dengan aman.

---

## 🛠️ Teknologi yang Digunakan

### Backend
- **Bahasa**: [Go (Golang)](https://go.dev/) (v1.20+)
- **Framework / Router**: [Fiber](https://gofiber.io/) / [Gin](https://gin-gonic.com/) / Net/HTTP
- **Database**: MySQL / PostgreSQL / SQLite
- **ORM / Driver**: GORM / `database/sql`

### Frontend
- **Framework**: [Vue.js 3](https://vuejs.org/) (Composition API)
- **Build Tool**: [Vite](https://vitejs.dev/)
- **HTTP Client**: [Axios](https://axios-http.com/)
- **Routing**: Vue Router
- **Styling**: Tailwind CSS / Bootstrap / CSS Custom

---

## 📁 Struktur Proyek

```text
CRUD-Golang-and-Vue3/
├── backend/                  # REST API Server (Golang)
│   ├── config/               # Konfigurasi Database & Env
│   ├── controllers/          # Handler Logic / Controllers
│   ├── models/               # Struct & Schema Database
│   ├── routes/               # Routing Endpoints API
│   ├── .env.example          # Contoh variabel lingkungan backend
│   ├── go.mod                # Modul Go & Dependensi
│   └── main.go               # Entry point server Go
│
├── frontend/                 # Client App (Vue.js 3)
│   ├── src/
│   │   ├── assets/           # Gambar, Icon, & Styles
│   │   ├── components/       # Komponen UI Reusable
│   │   ├── services/         # API Service (Axios Config)
│   │   ├── views/            # Halaman Tampilan (Home, Edit, Add)
│   │   ├── App.vue           # Root Component
│   │   └── main.js           # Entry point Vue.js
│   ├── index.html            # HTML Template
│   ├── package.json          # Dependensi Node.js
│   └── vite.config.js        # Konfigurasi Vite
│
└── README.md                 # Dokumentasi Proyek
```

---

## ⚙️ Prasyarat Sistem

Sebelum menjalankan proyek ini, pastikan komputer Anda telah terinstall:
- **Go**: versi `1.20` atau yang lebih baru ([Download Go](https://go.dev/dl/))
- **Node.js**: versi `18.x` atau yang lebih baru ([Download Node.js](https://nodejs.org/))
- **Package Manager**: `npm` atau `pnpm` / `yarn`
- **Database Server**: MySQL / PostgreSQL / SQLite

---

## 🚀 Panduan Instalasi & Cara Jalankan

### Clone Repository
```bash
git clone [https://github.com/rhanadam32/CRUD-Golang-and-Vue3.git](https://github.com/rhanadam32/CRUD-Golang-and-Vue3.git)
cd CRUD-Golang-and-Vue3