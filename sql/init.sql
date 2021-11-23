DROP TABLE IF EXISTS user;

CREATE TABLE user (
    id              SERIAL,
    name            VARCHAR(80) NOT NULL,
    email           VARCHAR(250) UNIQUE,
    password        VARCHAR(250) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);