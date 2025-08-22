

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(50) UNIQUE NOT NULL,
    publisher VARCHAR(255),
    year INTEGER,
    category VARCHAR(100),
    description TEXT,
    stock INTEGER DEFAULT 0,
    available INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS members (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    member_code VARCHAR(20) UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL REFERENCES books(id),
    member_id INTEGER NOT NULL REFERENCES members(id),
    loan_date TIMESTAMP NOT NULL,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    status VARCHAR(20) DEFAULT 'borrowed',
    fine DECIMAL(10,2) DEFAULT 0,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn);
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_members_email ON members(email);
CREATE INDEX IF NOT EXISTS idx_members_member_code ON members(member_code);
CREATE INDEX IF NOT EXISTS idx_loans_book_id ON loans(book_id);
CREATE INDEX IF NOT EXISTS idx_loans_member_id ON loans(member_id);
CREATE INDEX IF NOT EXISTS idx_loans_status ON loans(status);
CREATE INDEX IF NOT EXISTS idx_loans_due_date ON loans(due_date);

INSERT INTO users (username, email, password, role) VALUES 
('admin', 'admin@library.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin'),
('librarian', 'librarian@library.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user')
ON CONFLICT (username) DO NOTHING;

INSERT INTO books (title, author, isbn, publisher, year, category, description, stock, available) VALUES 
('The Great Gatsby', 'F. Scott Fitzgerald', '978-0743273565', 'Scribner', 1925, 'Fiction', 'A story of the fabulously wealthy Jay Gatsby and his love for the beautiful Daisy Buchanan.', 5, 5),
('To Kill a Mockingbird', 'Harper Lee', '978-0446310789', 'Grand Central Publishing', 1960, 'Fiction', 'The story of young Scout Finch and her father Atticus in a racially divided Alabama town.', 3, 3),
('1984', 'George Orwell', '978-0451524935', 'Signet Classic', 1949, 'Dystopian', 'A dystopian novel about totalitarianism and surveillance society.', 4, 4),
('Pride and Prejudice', 'Jane Austen', '978-0141439518', 'Penguin Classics', 1813, 'Romance', 'The story of Elizabeth Bennet and Mr. Darcy in Georgian-era England.', 2, 2)
ON CONFLICT (isbn) DO NOTHING;

INSERT INTO members (name, email, phone, address, member_code, status) VALUES 
('John Doe', 'john.doe@email.com', '+6281234567890', 'Jl. Sudirman No. 123, Jakarta', 'MEM000001', 'active'),
('Jane Smith', 'jane.smith@email.com', '+6281234567891', 'Jl. Thamrin No. 456, Jakarta', 'MEM000002', 'active'),
('Bob Johnson', 'bob.johnson@email.com', '+6281234567892', 'Jl. Gatot Subroto No. 789, Jakarta', 'MEM000003', 'active')
ON CONFLICT (email) DO NOTHING;
