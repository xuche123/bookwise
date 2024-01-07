CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    author VARCHAR(500) NOT NULL,
    image_url VARCHAR(500),
    description TEXT,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    version INT NOT NULL DEFAULT 1
);