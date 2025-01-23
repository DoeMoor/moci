# Moduline Controller Inventory

## Technology Stack

### Database Tools (Go)
- Fiber (http) - [Documentation](https://docs.gofiber.io/)
- SQLC (SQL Compiler) - [Documentation](https://docs.sqlc.dev/en/stable/index.html)
- Goose (Database Migration Tool) - [GitHub](https://github.com/pressly/goose)

### Frontend
- Vue.js with Vuetify component framework

## Running the Application

run only frontend end backend
```bash
git copy https://github.com/DoeMoor/moci_backend
git copy https://github.com/DoeMoor/moci_front

cd ./moci_front
npm install
npm run buildModule

cd ./../moci_backend
(cd ./src && go build -o ./../_build/moci && ./../_build/moci)
```

---
### Backend
Build and run the Go application: 

**from root of project !** 
```bash
(cd ./src && go build -o ./../_build/moci && ./../_build/moci)
```

### Frontend
##### Start the development server:

**from root of project !** 
```bash
npm run dev
```
##### Build module for backend:

**from root of project !** 

This is will execute `"vite build --outDir ./../moci_back/src/web/ --emptyOutDir"`
```bash
npm run buildModule
```
---
QnA

*Q* -- Ar you ahuel?

*A* -- Yes, i will add frontend as git submodule to backend.. eventually. Also i will add one script for build and launch app and for DB_init .

---

### Database Setup
The following SQL scripts are available:

- db_init.sql - Creates and initializes the database schema
- insert_scratch.sql - Populates the database with test data
- select_scratch.sql - Contains a examples of SELECT queries
