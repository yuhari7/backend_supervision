-- Seeder untuk tabel posts
INSERT INTO posts (title, content, category, status, created_date, updated_date)
VALUES
    ('How to Build a RESTful API in Go', 
     'In this tutorial, we will walk through the steps to build a RESTful API using Go and GORM.',
     'Tech', 
     'Published', 
     CURRENT_TIMESTAMP, 
     CURRENT_TIMESTAMP),
     
    ('Understanding Golang Interfaces', 
     'Golang interfaces allow us to define behaviors that types can implement. In this article, we will dive into how interfaces work.',
     'Programming', 
     'Draft', 
     CURRENT_TIMESTAMP, 
     CURRENT_TIMESTAMP),

    ('Top 10 Web Development Trends in 2025', 
     'Web development is ever-evolving. Let’s look at the top 10 trends that will dominate web development in 2025.',
     'Web Development', 
     'Published', 
     CURRENT_TIMESTAMP, 
     CURRENT_TIMESTAMP);
