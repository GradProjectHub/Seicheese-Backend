-- +goose Up
-- +goose StatementBegin
CREATE TABLE admin (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT FOREIGN KEY REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE admin;
-- +goose StatementEnd
