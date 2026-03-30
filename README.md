# Lacy Lorel

Full-stack eCommerce platform for the Lacy Lorel clothing brand.

Built with a scalable backend using Go and the Chi router, and a modern frontend (coming soon) using Next.js and TypeScript.

---

## 🚀 Tech Stack

**Backend**

* Go
* Chi Router
* PostgreSQL
* Docker

**Frontend (Planned)**

* Next.js
* TypeScript
* Tailwind CSS

---

## 📦 Features

* RESTful API with clean architecture (handler → service → repository)
* PostgreSQL database integration
* Dockerized development environment
* Product listing endpoint (`/api/v1/products`)

---

## 🛠️ Development Setup

### Backend

```bash
cd backend
docker compose up -d
go run cmd/server/main.go
```

---

## 📋 Roadmap

### Backend

* [x] Project structure & architecture
* [x] Products API endpoint
* [x] PostgreSQL integration
* [ ] SQLC integration
* [ ] Authentication (JWT)
* [ ] Orders & checkout

### Frontend

* [ ] Initialize Next.js app (TypeScript)
* [ ] Product listing page
* [ ] Product detail page
* [ ] Shopping cart
* [ ] Checkout flow

---

## 📌 Notes

This project is being built with a focus on scalable architecture and real-world production practices.
