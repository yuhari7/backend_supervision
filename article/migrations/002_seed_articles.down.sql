-- Rollback seeder untuk tabel posts
DELETE FROM posts WHERE title IN (
    'How to Build a RESTful API in Go',
    'Understanding Golang Interfaces',
    'Top 10 Web Development Trends in 2025'
);
