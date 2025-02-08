BEGIN;

DROP INDEX IF EXISTS idx_borrowed_books_book_id;
DROP INDEX IF EXISTS idx_borrowed_books_user_id;

DROP TABLE IF EXISTS borrowed_books;

COMMIT;