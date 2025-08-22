# Library Management System - Setup Guide

## Prerequisites

Before setting up the Library Management System, ensure you have the following installed:

- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **PostgreSQL 12+** - [Download here](https://www.postgresql.org/download/)
- **Postman** - [Download here](https://www.postman.com/downloads/) (for API testing)

## Step-by-Step Setup

### 1. Clone and Navigate to Project

```bash
git clone <repository-url>
cd library-management-system
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Database Setup

#### 3.1 Create PostgreSQL Database

1. Open PostgreSQL command line or pgAdmin
2. Create a new database:

```sql
CREATE DATABASE library_db;
```

#### 3.2 Run Database Migration

1. Open the `migrations/schema.sql` file
2. Execute the SQL commands in your PostgreSQL database
3. This will create all necessary tables and insert sample data

Alternatively, you can run the migration using psql:

```bash
psql -U postgres -d library_db -f migrations/schema.sql
```

### 4. Environment Configuration

1. Copy the `config.env` file and update it with your database credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=library_db
JWT_SECRET=your-secret-key-here
SERVER_PORT=8080
```

**Important Notes:**
- Replace `your_password_here` with your actual PostgreSQL password
- Replace `your-secret-key-here` with a strong secret key for JWT tokens
- The `JWT_SECRET` should be at least 32 characters long

### 5. Run the Application

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### 6. Verify Installation

1. Open your browser or Postman
2. Send a GET request to: `http://localhost:8080/health`
3. You should receive:

```json
{
  "status": "success",
  "message": "Library Management System is running"
}
```

## Testing with Postman

### 1. Import Collection

1. Open Postman
2. Click "Import" button
3. Select the `docs/postman_collection.json` file
4. The collection will be imported with all endpoints

### 2. Set Environment Variables

1. In Postman, click on the "Environment" dropdown
2. Create a new environment or use the default
3. Set the following variables:
   - `base_url`: `http://localhost:8080`
   - `auth_token`: (leave empty, will be set automatically)

### 3. Test Authentication

1. Run the "Login User" request first
2. Use the sample credentials:
   - Username: `admin`
   - Password: `password`
3. The JWT token will be automatically saved to the `auth_token` variable

### 4. Test All Endpoints

Now you can test all other endpoints. The authentication token will be automatically included in protected requests.

## Sample Data

The system comes with pre-loaded sample data:

### Users
- **Admin**: username: `admin`, password: `password`, role: `admin`
- **Librarian**: username: `librarian`, password: `password`, role: `user`

### Books
- The Great Gatsby (F. Scott Fitzgerald)
- To Kill a Mockingbird (Harper Lee)
- 1984 (George Orwell)
- Pride and Prejudice (Jane Austen)

### Members
- John Doe (MEM000001)
- Jane Smith (MEM000002)
- Bob Johnson (MEM000003)

## Troubleshooting

### Common Issues

#### 1. Database Connection Error
```
Failed to connect to database: dial tcp [::1]:5432: connect: connection refused
```

**Solution:**
- Ensure PostgreSQL is running
- Check if the database credentials in `config.env` are correct
- Verify the database `library_db` exists

#### 2. Port Already in Use
```
Failed to start server: listen tcp :8080: bind: address already in use
```

**Solution:**
- Change the `SERVER_PORT` in `config.env` to another port (e.g., 8081)
- Or stop the process using port 8080

#### 3. JWT Token Issues
```
Invalid or expired token
```

**Solution:**
- Make sure you're logged in and have a valid token
- Check if the `JWT_SECRET` in `config.env` is set correctly
- Try logging in again to get a fresh token

#### 4. Validation Errors
```
Validation failed
```

**Solution:**
- Check the request body format
- Ensure all required fields are provided
- Verify data types (e.g., numbers for IDs, valid email format)

### Database Reset

If you need to reset the database:

1. Drop and recreate the database:
```sql
DROP DATABASE library_db;
CREATE DATABASE library_db;
```

2. Run the migration again:
```bash
psql -U postgres -d library_db -f migrations/schema.sql
```

## Development

### Project Structure

```
library-management-system/
├── cmd/main.go                 # Application entry point
├── internal/
│   ├── config/database.go      # Database configuration
│   ├── models/                 # Data models
│   ├── handlers/               # HTTP handlers
│   ├── middleware/             # Middleware functions
│   ├── repository/             # Database operations
│   └── utils/                  # Utility functions
├── migrations/schema.sql       # Database schema
├── docs/                       # Documentation
├── go.mod                      # Go module file
├── config.env                  # Environment variables
└── README.md                   # Project documentation
```

### Adding New Features

1. **Models**: Add new models in `internal/models/`
2. **Repository**: Create repository methods in `internal/repository/`
3. **Handlers**: Add HTTP handlers in `internal/handlers/`
4. **Routes**: Update routes in `cmd/main.go`

### Code Style

- Follow Go conventions and best practices
- Use meaningful variable and function names
- Add comments for complex logic
- Handle errors properly
- Use consistent formatting (run `go fmt`)

## Production Deployment

For production deployment, consider:

1. **Environment Variables**: Use proper environment variable management
2. **Database**: Use a production-grade PostgreSQL instance
3. **Security**: 
   - Use strong JWT secrets
   - Enable HTTPS
   - Implement rate limiting
   - Add input validation
4. **Monitoring**: Add logging and monitoring
5. **Backup**: Implement database backup strategies

## Support

If you encounter any issues:

1. Check the troubleshooting section above
2. Review the API documentation in `docs/API_DOCUMENTATION.md`
3. Check the logs for error messages
4. Ensure all prerequisites are properly installed
