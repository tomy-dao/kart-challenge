# Kart Challenge Backend

Backend service built with Go (Golang) using Clean Architecture, providing APIs for product and order management.

**IMPORTANT**: Before running the application, you need to extract the compressed coupon data files located in `./backend-challenge/repository/data`

## 🏗️ Backend Structure

### Overall Architecture
```
backend-challenge/
├── cmd/                    # Entry point and application initialization
├── config/                 # Application configuration
├── endpoint/               # Business logic layer (endpoints)
├── model/                  # Data models and structs
├── repository/             # Data access layer
├── service/                # Business logic layer
├── transport/              # HTTP transport layer
├── utils/                  # Utility functions
├── client/                 # External client integrations
├── main.go                 # Main application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── Dockerfile              # Production Docker image
└── Dockerfile.dev          # Development Docker image
```

### Detailed Layer Breakdown

#### 1. **cmd/** - Application Entry Point
- `cmd.go`: Initialize and run HTTP server
- Manage application lifecycle
- Handle graceful shutdown

#### 2. **config/** - Configuration Management
- `config.go`: Manage configuration from environment variables
- Settings: HTTP port, host, database connection

#### 3. **model/** - Data Models
- `product.go`: Product model
- `order.go`: Order model
- `params.go`: Request/response parameters
- `response.go`: Response structures

#### 4. **repository/** - Data Access Layer
- `repo.go`: Interface and implementation for data access
- `product.go`: Product repository operations
- `data/`: Contains coupon data (couponbase1.gz, couponbase2.gz, couponbase3.gz)

#### 5. **service/** - Business Logic Layer
- `service.go`: Service interface and implementation
- `product.go`: Product business logic
- `order.go`: Order business logic

#### 6. **endpoint/** - API Endpoints
- `initial.go`: Initialize endpoints
- `product.go`: Product-related endpoints
- `order.go`: Order-related endpoints

#### 7. **transport/http/** - HTTP Transport Layer
- `router.go`: HTTP routing with Chi router
- `handler.go`: HTTP request handlers
- `middleware.go`: HTTP middleware
- `transport.go`: HTTP transport setup

#### 8. **client/** - External Integrations
- `init.go`: Client initialization and configuration

### 2.5. Extract coupon data (Required)
Before running the application, you need to extract the compressed coupon data files located in `./backend-challenge/repository/data`:




## 🚀 How to Run Source Code

### System Requirements
- Go 1.23.4 or higher
- Git

### 1. Clone repository
```bash
git clone <repository-url>
cd kart-challenge/backend-challenge
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Configure environment variables
Create a `.env` file or set environment variables:

```bash
export HTTP_PORT=8080
export HOST=0.0.0.0
```

### 4. Run the application

#### Run directly
```bash
go run main.go
```

#### Build and run
```bash
go build -o main .
./main
```

### 5. Run with Docker

#### Development mode
```bash
docker build -f Dockerfile.dev -t kart-challenge-dev .
docker run -p 8080:8080 kart-challenge-dev
```

#### Production mode
```bash
docker build -t kart-challenge .
docker run -p 8080:8080 kart-challenge
```

#### Docker Compose (if available)
```bash
docker-compose up -d
```

## 📡 API Endpoints

### Health Check
- `GET /health` - Check service status

### Products
- `GET /api/product` - Get all products
- `GET /api/product/{id}` - Get product by ID

### Orders
- `POST /api/order` - Create new order

## 🔧 Development

### Project Structure following Clean Architecture
```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP Transport Layer                     │
├─────────────────────────────────────────────────────────────┤
│                    Endpoint Layer                           │
├─────────────────────────────────────────────────────────────┤
│                    Service Layer                            │
├─────────────────────────────────────────────────────────────┤
│                   Repository Layer                          │
└─────────────────────────────────────────────────────────────┘
```

### Main Dependencies
- **Chi Router**: HTTP routing
- **JWT**: Authentication
- **BoltDB**: Embedded key-value store for coupon data

## 🐛 Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   # Check which process is using the port
   lsof -i :8080
   
   # Kill the process using the port
   kill -9 <PID>
   ```

2. **Database connection failed**
   - Check if MySQL service is running
   - Verify database connection information
   - Check firewall settings

3. **Permission denied**
   ```bash
   # Grant execution permission
   chmod +x main
   ```

## 📚 References

- [Go Documentation](https://golang.org/doc/)
- [Chi Router](https://github.com/go-chi/chi)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
