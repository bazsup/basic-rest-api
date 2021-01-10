CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);