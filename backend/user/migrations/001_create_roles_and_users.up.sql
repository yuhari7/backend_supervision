CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role_id INT NOT NULL REFERENCES roles(id) ON DELETE RESTRICT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles (name) VALUES ('admin'), ('contributor');

ALTER TABLE users
ADD COLUMN is_active BOOLEAN DEFAULT TRUE;

ALTER TABLE users
ADD COLUMN deleted_at TIMESTAMP;

INSERT INTO users (name, email, password, role_id)
VALUES ('Super Admin', 'admin@example.com', '$2a$10$86r/dZye2Ge45jqF4hptkeF07GD0AghyHdcqGNxbjANs29Ro0oge.', 1);
