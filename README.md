# 📝 NoteTaker

![Go](https://img.shields.io/badge/Go-1.21-blue?style=for-the-badge&logo=go)
![Clean Architecture](https://img.shields.io/badge/Clean--Architecture-%F0%9F%92%AA-blueviolet?style=for-the-badge)
![SOLID](https://img.shields.io/badge/SOLID-Principles-ff69b4?style=for-the-badge)
![License](https://img.shields.io/github/license/RoyRoki/NoteTaker?style=for-the-badge)

A simple and modular 🧱 **Go-based Note Taking API** built with:

- 💎 **Clean Architecture**
- 💡 **SOLID Principles**
- 🧪 **Testable structure**
- 🗃️ SQLite (can be easily swapped)
- ⚙️ Dependency Injection for better maintainability
- 🌐 Simple REST API

---

## 📽️ Demo

Check out a quick walkthrough of the project setup and code structure:

[![NoteTaker Demo](https://img.youtube.com/vi/C6deIp21r5E/0.jpg)](https://youtu.be/C6deIp21r5E?si=LDtkRswd7m7gnUbH)

---

## 📦 Project Structure

```bash
.
├── cmd                # App entrypoint
│   └── server         # Main server file
├── commons            # Shared utilities (config, logging, DB, etc.)
├── modules
│   └── notes
│       ├── data       # Data layer (models, datasource, impls)
│       ├── domain     # Domain layer (interfaces, usecases)
│       ├── di         # Dependency injection
│       └── presentation # HTTP controllers and routers
├── test.html          # Quick test UI (excluded from stats)
├── notes.db           # SQLite database file
└── README.md          # You're here!
```

---

## 🧠 Key Concepts

### 🧼 Clean Architecture

- Separated concerns across data, domain, and presentation
- Easily testable and scalable

### 🧱 SOLID Principles

- Each layer has a single responsibility
- Interfaces are implemented explicitly
- High-level code is decoupled from low-level details

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/RoyRoki/NoteTaker.git && cd NoteTaker
```

### 2. Run the server

```bash
go run cmd/server/main.go
```

### 3. Access the API

```bash
GET     /notes
POST    /notes
PUT     /notes/:id
DELETE  /notes/:id
```

You can use `test.html` locally to quickly test endpoints.

---

## 🛠️ Tech Stack

- Go 🐹
- SQLite 🛢️
- chi router 🧭
- sqlx (optionally pluggable)

---

## 🤝 Contributions

PRs and stars are welcome! ⭐ If you're interested in Clean Architecture or Go backend design, this project is a great starting point!

---

## 📄 License

ROCKET © [RoyRoki](https://github.com/RoyRoki)
