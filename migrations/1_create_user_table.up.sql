CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255),
    username VARCHAR(255),
    email VARCHAR(255),
    roles ENUM('admin', 'user')
);