CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE wallets (
    user_id INT PRIMARY KEY,
    balance BIGINT NOT NULL DEFAULT 0,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    event_type TEXT NOT NULL,
    from_user INT,
    to_user INT,
    amount BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);