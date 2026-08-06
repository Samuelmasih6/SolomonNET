CREATE TABLE testimonies (
    id BIGSERIAL PRIMARY KEY,

    case_id BIGINT NOT NULL,

    witness_id BIGINT NOT NULL,

    suspect TEXT NOT NULL,

    available BOOLEAN NOT NULL,

    response_time_ms INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY (case_id)
        REFERENCES cases(id)
        ON DELETE CASCADE,

    FOREIGN KEY (witness_id)
        REFERENCES witnesses(id)
);
