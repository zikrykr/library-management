BEGIN;

CREATE TABLE IF NOT EXISTS books (
  id VARCHAR(255) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  isbn VARCHAR(20) UNIQUE NOT NULL,
  author_id VARCHAR(255) NOT NULL,
  category_id VARCHAR(255) NOT NULL,
  published_year INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

--  Index for faster book title search
CREATE INDEX idx_books_title ON books USING GIN (to_tsvector('english', title));

-- Index for faster ISBN lookup
CREATE INDEX idx_books_isbn ON books (isbn);

-- Index for author-based filtering
CREATE INDEX idx_books_author ON books (author_id);

-- Index for category-based filtering
CREATE INDEX idx_books_category ON books (category_id);

COMMIT;