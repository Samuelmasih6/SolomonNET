CREATE TABLE advisor_recommendations (

    id BIGSERIAL PRIMARY KEY,

    case_id BIGINT NOT NULL,

    recommendation TEXT NOT NULL,

    confidence NUMERIC(5,2),

    reasoning TEXT,

    model_version VARCHAR(50),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    FOREIGN KEY(case_id)
        REFERENCES cases(id)
        ON DELETE CASCADE
);
