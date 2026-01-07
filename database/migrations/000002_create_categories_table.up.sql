CREATE TABLE IF NOT EXISTS "categories" (
    id SERIAL PRIMARY KEY,
    title varchar(100) NOT NULL,
    slug varchar(200) UNIQUE NOT NULL,
    created_by_id INT REFERENCES "users"(id) ON DELETE CASCADE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_categories_created_by_id ON "categories"(created_by_id);
