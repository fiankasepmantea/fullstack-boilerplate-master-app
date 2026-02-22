# Frontend - Payment Dashboard

Frontend aplikasi payment dashboard yang dibangun menggunakan **Nuxt 3** dengan **Vue 3** dan **Tailwind CSS**.

## 🚀 Tech Stack

- **Framework**: Nuxt 3 (^4.3.1)
- **Language**: TypeScript
- **UI**: Vue 3 (^3.5.28)
- **Styling**: Tailwind CSS
- **HTTP Client**: Native Fetch API
- **State Management**: Vue Composition API
- **Authentication**: JWT (JSON Web Token)

## 📋 Prerequisites

Pastikan sudah terinstall:

- Node.js (v24.x atau lebih baru)
- npm atau yarn
- Backend API berjalan di http://localhost:8080

## 🛠️ Installation

```bash
# Install dependencies
npm install

# Atau jika menggunakan yarn
yarn install

🏃 Run Development Server
# Jalankan development server
npm run dev
# Server akan berjalan di http://localhost:3000

🔨 Build untuk Production
# Build aplikasi
npm run build
# Preview hasil build
npm run preview

📁 Struktur Folder
frontend/
├── app/
│   ├── app.vue              # Root component
│   ├── assets/
│   │   └── css/
│   │       └── main.css     # Global styles (Tailwind)
│   ├── components/
│   │   ├── Header.vue       # Header component
│   │   ├── Footer.vue       # Footer component
│   │   └── PaymentList.vue  # Payment list component
│   ├── composables/
│   │   ├── useApi.ts        # API client helper
│   │   ├── useAuth.ts       # Authentication logic
│   │   └── usePayment.ts    # Payment logic
│   ├── layout/
│   │   └── default.vue      # Default layout
│   ├── middleware/
│   │   └── auth.global.ts   # Global auth middleware
│   └── pages/
│       ├── index.vue        # Homepage
│       ├── login.vue        # Login page
│       ├── dashboard.vue    # Protected dashboard
│       └── payments.vue     # Payments page
├── public/                   # Static assets
├── nuxt.config.ts           # Nuxt configuration
├── tailwind.config.js       # Tailwind configuration
└── tsconfig.json            # TypeScript configuration

🔐 Authentication
Aplikasi menggunakan JWT untuk autentikasi:
1. Login: User login dengan email dan password
2. Token Storage: Token disimpan di localStorage
3. Protected Routes: Middleware auth.global.ts melindungi halaman yang membutuhkan autentikasi
4. Auto Redirect: User yang sudah login akan di-redirect dari halaman login ke dashboard
Email: cs@test.com
Password: password
Email: operation@test.com
Password: password

🌐 API Configuration
API base URL dikonfigurasi di nuxt.config.ts:
runtimeConfig: {
  public: {
    apiBase: "http://localhost:8080"
  }
}

API Endpoints
POST /dashboard/v1/auth/login - Login
GET /dashboard/v1/payments - List payments (protected)
PUT /dashboard/v1/payment/{id}/review - Review payment (protected)

🎨 Styling
Menggunakan Tailwind CSS untuk styling. Custom CSS global ada di app/assets/css/main.css.

📦 Composables
useAuth()
Handling autentikasi:
const { token, login, logout, isAuthenticated } = useAuth()

// Login
await login('email@test.com', 'password')

// Logout
logout()

// Check auth status
if (isAuthenticated()) {
  // User is logged in
}

usePayment()
Handling payment data:
const { payments, fetchPayments } = usePayment()

// Fetch payments
await fetchPayments()

// Access payments
console.log(payments.value)

useApi()
HTTP client helper:
const { get, post } = useApi()

// GET request
const data = await get('/endpoint')

// POST request
const result = await post('/endpoint', { data })

Environment Variables
Untuk development, API URL sudah di-set di nuxt.config.ts. Untuk production, Anda bisa override menggunakan environment variables:
# .env file
NUXT_PUBLIC_API_BASE=http://your-api-url.com

Troubleshooting
CORS Issues
Pastikan backend sudah mengizinkan CORS untuk origin http://localhost:3000.
Authentication Issues
- Pastikan token tersimpan di localStorage
- Periksa apakah middleware auth berfungsi dengan benar
- Pastikan backend API berjalan dan accessible

Build Errors
- Hapus folder .nuxt dan node_modules
- Install ulang dependencies: npm install
- Jalankan build ulang: npm run build

📝 Development Notes
Hot Reload: Nuxt 3 mendukung hot reload secara otomatis
TypeScript: Semua file menggunakan TypeScript untuk type safety
Composition API: Menggunakan Vue 3 Composition API dengan <script setup>
Auto Imports: Nuxt 3 auto-import components dan composables

🤝 Contributing
