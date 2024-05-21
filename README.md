# Overview
This documentation provides a guide to the web service project implemented in Go (Golang) using the Chi framework for routing and Redis for data storage. The web service includes basic CRUD operations for managing orders.

### Project Structure
```plaintext
/papi
│
├── /cmd
│   └── /server
│       └── main.go
│
├── /internal
│   ├── /app
│   │   └── app.go
│   ├── /handler
│   │   └── order.go
│   ├── /model
│   │   └── order.go
│   ├── /repository
│   │   └── redis.go
│   └── /router
│       └── router.go
│
├── /vendor
│
├── go.mod
└── go.sum
```

# Installation

### Clone the Repository:

```bash
git clone https://github.com/yourusername/web-service.git
```
### Install Dependencies:
```sh
go mod tidy
```
