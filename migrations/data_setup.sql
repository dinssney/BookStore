-- Enable the pgcrypto extension if not already enabled
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert a publisher with a hashed password
WITH inserted_publisher AS (
  INSERT INTO users (created_at, updated_at, username, email, password, role)
  VALUES (
    NOW(), 
    NOW(), 
    'publisher', 
    'publisher@gmail.com', 
    crypt('Password1', gen_salt('bf')), 
    'publisher'
  )
  RETURNING id
)
-- Use that ID to insert books
INSERT INTO books (title, author, description, publisher_id, cover_image_url)
SELECT 
  bd.title, 
  bd.author, 
  bd.description, 
  inserted_publisher.id, 
  bd.cover_image_url
FROM 
  inserted_publisher,
  (VALUES
    ('The Great Gatsby', 'F. Scott Fitzgerald', 'A story of the Jazz Age', 'https://covers.openlibrary.org/b/isbn/9780743273565-L.jpg'),
    ('To Kill a Mockingbird', 'Harper Lee', 'A novel about racial injustice', 'https://covers.openlibrary.org/b/isbn/9780061120084-L.jpg'),
    ('1984', 'George Orwell', 'A dystopian social science fiction novel', 'https://covers.openlibrary.org/b/isbn/9780451524935-L.jpg'),
    ('Pride and Prejudice', 'Jane Austen', 'A romantic novel of manners', 'https://covers.openlibrary.org/b/isbn/9780141439518-L.jpg'),
    ('The Catcher in the Rye', 'J. D. Salinger', 'A story of teenage rebellion', 'https://covers.openlibrary.org/b/isbn/9780316769488-L.jpg'),
    ('The Lord of the Rings', 'J.R.R. Tolkien', 'An epic high-fantasy novel', 'https://covers.openlibrary.org/w/id/12685829-L.jpg'),
    ('Animal Farm', 'George Orwell', 'A satirical allegorical novella', 'https://covers.openlibrary.org/b/isbn/9780451526342-L.jpg'),
    ('The Hobbit', 'J.R.R. Tolkien', 'A fantasy novel and children''s book', 'https://covers.openlibrary.org/b/isbn/9780547928227-L.jpg'),
    ('Brave New World', 'Aldous Huxley', 'A dystopian social science fiction novel', 'https://covers.openlibrary.org/b/isbn/9780060850524-L.jpg'),
    ('The Grapes of Wrath', 'John Steinbeck', 'A novel set during the Great Depression', 'https://covers.openlibrary.org/b/isbn/9780143039433-L.jpg')
  ) AS bd(title, author, description, cover_image_url);