CREATE TABLE events (
    id BIGSERIAL PRIMARY KEY,

    case_id BIGINT NOT NULL,

    event_type VARCHAR(50) NOT NULL,

    payload JSONB,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY (case_id)
        REFERENCES cases(id)
        ON DELETE CASCADE
);
