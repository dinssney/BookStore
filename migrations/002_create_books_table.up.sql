CREATE TABLE IF NOT EXISTS books (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT,
    publisher_id BIGINT REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_books_deleted_at ON books(deleted_at);

CREATE TABLE IF NOT EXISTS user_books (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    book_id BIGINT REFERENCES books(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, book_id)
);