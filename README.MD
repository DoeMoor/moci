# Moduline Controller Inventory

## 🛠️ Technology Stack

### Tools (Go)
- **SQLC (SQL Compiler)** - [📄 Documentation](https://docs.sqlc.dev/en/stable/index.html)
- **Goose (Database Migration Tool)** - [🛠️ GitHub](https://github.com/pressly/goose)
- **Webi (install and update Go version)** - [📚 Documentation](https://webinstall.dev/golang/)
- **Golangci-lint (Go linter)** - [📚 Documentation](https://golangci-lint.run/)

### Framework 
- **Fiber (Go HTTP framework)** - [🌐 Documentation](https://docs.gofiber.io/)

### 🌟 Frontend
- **Vue.js** with **Vuetify** component framework

### Dev Tools
- **Task (Task runner)** - [📚 Documentation](https://taskfile.dev/usage/)
- **Podman** - [📚 Documentation](https://podman.io/) 

---

## 📋 Tasks

### init_app
- **Directory**: `./moci_front`
- **Purpose**: Initializes the frontend application
- **Commands**:
  - Updates and initializes git submodules
  - Installs npm dependencies

### build_front
- **Directory**: `./moci_front`
- **Purpose**: Builds the frontend module
- **Command**: Runs the `buildModule` npm script

### clean_run_app
- **Purpose**: Cleans, rebuilds, and runs the entire application
- **Steps**:
  1. Clears the `_build` directory
  2. Builds the frontend module
  3. Copies the `.env` file to the build directory
  4. Copies the `web` directory to the build directory
  5. Builds the Go application and runs it

### run_go_lint
- **Directory**: `./src`
- **Purpose**: Runs the Go linter
- **Command**: Executes `golangci-lint run`

### init_db
- **Purpose**: Initializes a PostgreSQL database using Podman
- **Command**: Runs a Podman container with PostgreSQL, setting up the database for the application

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

## 🗄️ Database Setup
The following SQL scripts are available:

- **db_init.sql** - Creates and initializes the database schema
- **insert_scratch.sql** - Populates the database with test data
- **select_scratch.sql** - Contains examples of SELECT queries

Initializes a PostgreSQL database using Podman for **DEV** environment
- **DO NOT USE FOR PRODUCTION! Or change the password for db**.
```bash
task init_db
```

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
