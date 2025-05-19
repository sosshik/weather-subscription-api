CREATE TABLE subscriptions (
                               email TEXT PRIMARY KEY,
                               city TEXT NOT NULL,
                               frequency TEXT NOT NULL,
                               token TEXT NOT NULL UNIQUE,
                               is_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
                               created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
