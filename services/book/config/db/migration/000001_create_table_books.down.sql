BEGIN;

DROP INDEX IF EXISTS idx_books_title;
DROP INDEX IF EXISTS idx_books_isbn;
DROP INDEX IF EXISTS idx_books_author;
DROP INDEX IF EXISTS idx_books_category;

DROP TABLE IF EXISTS books;

COMMIT;