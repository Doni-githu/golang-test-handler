-- 001_create_person_table.up.sql
CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    age INTEGER,
    gender TEXT,
    nationality TEXT
);