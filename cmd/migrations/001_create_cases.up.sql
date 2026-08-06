CREATE TABLE cases (
    id BIGSERIAL PRIMARY KEY,

    question TEXT NOT NULL,

    status VARCHAR(30) NOT NULL,

    verdict VARCHAR(100),

    votes_received INTEGER NOT NULL DEFAULT 0,

    votes_total INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    closed_at TIMESTAMPTZ
);
