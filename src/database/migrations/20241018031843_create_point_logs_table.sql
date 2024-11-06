-- +goose Up
-- +goose StatementBegin
CREATE TABLE point_logs (
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    point INT NOT NULL,
    PRIMARY KEY (created_at, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE point_logs;
-- +goose StatementEnd
