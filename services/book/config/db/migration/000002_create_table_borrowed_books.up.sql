BEGIN;

CREATE TABLE IF NOT EXISTS borrowed_books (
  id VARCHAR(255) PRIMARY KEY,
  book_id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  borrowed_at TIMESTAMP NOT NULL,
  due_at TIMESTAMP NOT NULL,
  returned_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
);

-- index for book_id
CREATE INDEX idx_borrowed_books_book_id ON borrowed_books (book_id);

-- index for user_id
CREATE INDEX idx_borrowed_books_user_id ON borrowed_books (user_id);

COMMIT;