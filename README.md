# Bookstore CLI (Go + GORM + PostgreSQL)

A simple command-line interface (CLI) bookstore management system built in Go using GORM ORM and PostgreSQL.  
This project allows you to perform **CRUD operations** on books and authors directly from your terminal.

## Project Structure

```
bookstore-cli/
├── cmd/                     # CLI commands and main application entry point
├── config/                  # Database configuration and connection setup
├── model/                   # GORM models for database entities (Book, Author, etc.)
├── repository/              # Data access layer with GORM implementations
├── service/                 # Business logic layer
├── test/                    # Unit and integration tests
├── utils/                   # Utility functions and helpers
├── .gitignore              # Git ignore rules
└── README.md               # Project documentation
```

### Directory Descriptions

- **`cmd/`** - Contains the main application entry point and CLI command definitions
- **`config/`** - Database configuration, environment setup, and application settings
- **`model/`** - GORM models representing database entities (Book, Author, etc.)
- **`repository/`** - Data access layer implementing repository pattern with GORM
- **`service/`** - Business logic layer that orchestrates repository operations
- **`test/`** - Test files for unit and integration testing
- **`utils/`** - Shared utility functions and helper methods

## Tech Stack

- **Go** - Programming language for the CLI application
- **GORM** - ORM library for database operations
- **PostgreSQL** - Database for persistent storage
- **Go testing** - Built-in testing framework for unit tests
