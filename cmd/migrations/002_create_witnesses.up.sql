CREATE TABLE witnesses (
    id BIGSERIAL PRIMARY KEY,

    name TEXT NOT NULL,

    address TEXT NOT NULL UNIQUE,

    status VARCHAR(20) NOT NULL,

    reliability_score NUMERIC(5,2)
        NOT NULL DEFAULT 100.00,

    last_seen TIMESTAMPTZ,

    created_at TIMESTAMPTZ
        NOT NULL DEFAULT NOW()
);
