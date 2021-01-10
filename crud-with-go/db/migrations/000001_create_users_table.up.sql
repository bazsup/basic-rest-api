CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp DEFAULT '0000-00-00 00:00:00',
    updated_at timestamp NOT NULL,
    PRIMARY KEY (id)
);