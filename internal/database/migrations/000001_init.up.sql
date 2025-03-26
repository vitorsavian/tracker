CREATE TABLE novel (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    name VARCHAR(512) NOT NULL,
    page SMALLINT,
    finished BOOLEAN DEFAULT FALSE
);
