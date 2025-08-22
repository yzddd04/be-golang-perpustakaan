# Library Management System API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### 1. Health Check

#### GET /health
Check if the server is running.

**Response:**
```json
{
  "status": "success",
  "message": "Library Management System is running"
}
```

### 2. Authentication

#### POST /api/auth/register
Register a new user.

**Request Body:**
```json
{
  "username": "string",
  "email": "string",
  "password": "string",
  "role": "string" // optional, defaults to "user"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "role": "user",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### POST /api/auth/login
Login user and get JWT token.

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@library.com",
      "role": "admin",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  }
}
```

### 3. Books

#### GET /api/books
Get all books.

**Response:**
```json
{
  "status": "success",
  "message": "Books retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "isbn": "978-0743273565",
      "publisher": "Scribner",
      "year": 1925,
      "category": "Fiction",
      "description": "A story of the fabulously wealthy Jay Gatsby...",
      "stock": 5,
      "available": 3,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### GET /api/books/{id}
Get a specific book by ID.

**Response:**
```json
{
  "status": "success",
  "message": "Book retrieved successfully",
  "data": {
    "id": 1,
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "isbn": "978-0743273565",
    "publisher": "Scribner",
    "year": 1925,
    "category": "Fiction",
    "description": "A story of the fabulously wealthy Jay Gatsby...",
    "stock": 5,
    "available": 3,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### POST /api/books
Create a new book.

**Request Body:**
```json
{
  "title": "string",
  "author": "string",
  "isbn": "string",
  "publisher": "string",
  "year": 0,
  "category": "string",
  "description": "string",
  "stock": 0
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Book created successfully",
  "data": {
    "id": 1,
    "title": "Sample Book",
    "author": "Sample Author",
    "isbn": "978-1234567890",
    "publisher": "Sample Publisher",
    "year": 2023,
    "category": "Fiction",
    "description": "A sample book for testing",
    "stock": 10,
    "available": 10,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### PUT /api/books/{id}
Update a book.

**Request Body:**
```json
{
  "title": "string",
  "author": "string",
  "isbn": "string",
  "publisher": "string",
  "year": 0,
  "category": "string",
  "description": "string",
  "stock": 0
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Book updated successfully",
  "data": {
    "id": 1,
    "title": "Updated Book Title",
    "author": "Updated Author",
    "isbn": "978-1234567890",
    "publisher": "Sample Publisher",
    "year": 2023,
    "category": "Fiction",
    "description": "A sample book for testing",
    "stock": 15,
    "available": 15,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### DELETE /api/books/{id}
Delete a book.

**Response:**
```json
{
  "status": "success",
  "message": "Book deleted successfully"
}
```

### 4. Members

#### GET /api/members
Get all members.

**Response:**
```json
{
  "status": "success",
  "message": "Members retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@email.com",
      "phone": "+6281234567890",
      "address": "Jl. Sudirman No. 123, Jakarta",
      "member_code": "MEM000001",
      "status": "active",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### GET /api/members/{id}
Get a specific member by ID.

**Response:**
```json
{
  "status": "success",
  "message": "Member retrieved successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@email.com",
    "phone": "+6281234567890",
    "address": "Jl. Sudirman No. 123, Jakarta",
    "member_code": "MEM000001",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### POST /api/members
Create a new member.

**Request Body:**
```json
{
  "name": "string",
  "email": "string",
  "phone": "string",
  "address": "string"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Member created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+6281234567890",
    "address": "Jl. Sample No. 123, Jakarta",
    "member_code": "MEM000001",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### PUT /api/members/{id}
Update a member.

**Request Body:**
```json
{
  "name": "string",
  "email": "string",
  "phone": "string",
  "address": "string",
  "status": "string"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Member updated successfully",
  "data": {
    "id": 1,
    "name": "John Updated Doe",
    "email": "john.doe@example.com",
    "phone": "+6281234567891",
    "address": "Jl. Sample No. 123, Jakarta",
    "member_code": "MEM000001",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### DELETE /api/members/{id}
Delete a member.

**Response:**
```json
{
  "status": "success",
  "message": "Member deleted successfully"
}
```

### 5. Loans

#### GET /api/loans
Get all loans.

**Response:**
```json
{
  "status": "success",
  "message": "Loans retrieved successfully",
  "data": [
    {
      "id": 1,
      "book_id": 1,
      "member_id": 1,
      "loan_date": "2024-01-01T00:00:00Z",
      "due_date": "2024-02-01T00:00:00Z",
      "return_date": null,
      "status": "borrowed",
      "fine": 0,
      "notes": "Test loan",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "book": {
        "id": 1,
        "title": "The Great Gatsby",
        "author": "F. Scott Fitzgerald",
        "isbn": "978-0743273565"
      },
      "member": {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@email.com",
        "member_code": "MEM000001"
      }
    }
  ]
}
```

#### GET /api/loans/{id}
Get a specific loan by ID.

**Response:**
```json
{
  "status": "success",
  "message": "Loan retrieved successfully",
  "data": {
    "id": 1,
    "book_id": 1,
    "member_id": 1,
    "loan_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-02-01T00:00:00Z",
    "return_date": null,
    "status": "borrowed",
    "fine": 0,
    "notes": "Test loan",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z",
    "book": {
      "id": 1,
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "isbn": "978-0743273565"
    },
    "member": {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@email.com",
      "member_code": "MEM000001"
    }
  }
}
```

#### POST /api/loans
Create a new loan.

**Request Body:**
```json
{
  "book_id": 1,
  "member_id": 1,
  "due_date": "2024-02-01T00:00:00Z",
  "notes": "string"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Loan created successfully",
  "data": {
    "id": 1,
    "book_id": 1,
    "member_id": 1,
    "loan_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-02-01T00:00:00Z",
    "return_date": null,
    "status": "borrowed",
    "fine": 0,
    "notes": "Test loan",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### PUT /api/loans/{id}/return
Return a borrowed book.

**Response:**
```json
{
  "status": "success",
  "message": "Book returned successfully",
  "data": {
    "id": 1,
    "book_id": 1,
    "member_id": 1,
    "loan_date": "2024-01-01T00:00:00Z",
    "due_date": "2024-02-01T00:00:00Z",
    "return_date": "2024-01-15T00:00:00Z",
    "status": "returned",
    "fine": 0,
    "notes": "Test loan",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-15T00:00:00Z"
  }
}
```

## Error Responses

### Validation Error (400)
```json
{
  "status": "error",
  "message": "Validation failed",
  "error": "Field validation error details"
}
```

### Unauthorized (401)
```json
{
  "status": "error",
  "message": "Authorization header is required",
  "error": "Authorization header is required"
}
```

### Not Found (404)
```json
{
  "status": "error",
  "message": "Resource not found",
  "error": "Resource not found"
}
```

### Conflict (409)
```json
{
  "status": "error",
  "message": "Resource already exists",
  "error": "Resource already exists"
}
```

### Internal Server Error (500)
```json
{
  "status": "error",
  "message": "Internal server error",
  "error": "Internal server error"
}
```

## Testing with Postman

1. Import the `docs/postman_collection.json` file into Postman
2. Set the environment variable `base_url` to `http://localhost:8080`
3. Run the "Login User" request to get a JWT token
4. The token will be automatically set in the `auth_token` variable
5. Test all other endpoints

## Sample Data

The system comes with sample data:
- **Users**: admin/password, librarian/password
- **Books**: The Great Gatsby, To Kill a Mockingbird, 1984, Pride and Prejudice
- **Members**: John Doe, Jane Smith, Bob Johnson
