-- Users таблицасы
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name TEXT UNIQUE NOT NULL,
                                     email TEXT UNIQUE NOT NULL,
                                     password TEXT NOT NULL,
                                     age INTEGER,
                                     role TEXT NOT NULL DEFAULT 'user'
);

-- Habits таблицасы
CREATE TABLE IF NOT EXISTS habits (
                                      id SERIAL PRIMARY KEY,
                                      title TEXT NOT NULL,
                                      user_id INTEGER NOT NULL,
                                      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
