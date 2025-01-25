# Moduline Controller Inventory

## 🛠️ Technology Stack

### Tools (Go)
- **SQLC (SQL Compiler)** - [📄 Documentation](https://docs.sqlc.dev/en/stable/index.html)
- **Goose (Database Migration Tool)** - [🛠️ GitHub](https://github.com/pressly/goose)
- **Webi (install and update Go version)** - [📚 Documentation](https://webinstall.dev/golang/)

#### Framework 
- **Fiber (Go HTTP framework)** - [🌐 Documentation](https://docs.gofiber.io/)

### 🌟 Frontend
- **Vue.js** with **Vuetify** component framework

---

## 🚀 Running the Application

### Run Only Frontend and Backend
```bash
# Clone repositories
git clone https://github.com/DoeMoor/moci_backend
git clone https://github.com/DoeMoor/moci_front

# Set up frontend end build frontend module
cd ./moci_front
npm install
npm run buildModule

# Build end run backend
cd ./../moci_backend
(cd ./src && go build -o ./../_build/moci && ./../_build/moci)
```

---

### Backend 🔚
#### Build and Run the Go Application:

**Run from the project root directory!**
```bash
(cd ./src && go build -o ./../_build/moci && ./../_build/moci)
```

### Frontend 🖥️
#### Start the Development Server:

**Run from the project root directory!**
```bash
npm run dev
```

#### Build Module for Backend:

**Run from the project root directory!**

This will execute `vite build --outDir ./../moci_back/src/web/ --emptyOutDir`
```bash
npm run buildModule
```

---

## ❓ QnA

- **Q:** Are you ahuel?
  
  **A:** Yes, I will add the frontend as a Git submodule to the backend... eventually. Also, I plan to add one script for building, launching the app, and initializing the database.

---

## 🗄️ Database Setup
The following SQL scripts are available:

- **db_init.sql** - Creates and initializes the database schema
- **insert_scratch.sql** - Populates the database with test data
- **select_scratch.sql** - Contains examples of SELECT queries

---

## 🔧 Tools Usage

### Webi
```bash
webi webi # To update Webi
webi go@stable # To update or switch Go versions
```

### Goose
```bash
# TODO
```

### SQLC
```bash
sqlc generate # To generate .go files from .sql queries
