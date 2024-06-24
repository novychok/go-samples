-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    name TEXT NOT NULL,
    author TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS events (
    id BIGSERIAL PRIMARY KEY,
    event_type VARCHAR(255) NOT NULL,
    payload TEXT,
    status TEXT NOT NULL DEFAULT 'new' CHECK(status IN ('new', 'done')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reserved_to TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
