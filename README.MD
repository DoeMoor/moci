# Moduline Controller Inventory

## Technology Stack

### Database Tools (Go)
- SQLC (SQL Compiler) - [Documentation](https://docs.sqlc.dev/en/stable/index.html)
- Goose (Database Migration Tool) - [GitHub](https://github.com/pressly/goose)

### Frontend
- Vue.js with Vuetify component framework

## Running the Application

### Backend
Build and run the Go application:
```bash
go build -o ./_build/moci && ./_build/moci
```

### Frontend
Start the development server:
```bash
npm run dev
```

##### Database Setup
The following SQL scripts are available:

- db_init.sql - Creates and initializes the database schema
- insert_scratch.sql - Populates the database with test data
- select_scratch.sql - Contains example SELECT queries
