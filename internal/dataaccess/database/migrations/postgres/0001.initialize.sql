-- +migrate Up
CREATE TABLE IF NOT EXISTS ais_account (
    account_id BIGSERIAL PRIMARY KEY,
    account_name VARCHAR(256) NOT NULL,
    account_type INT,
    account_status INT
);

-- +migrate Down
DROP TABLE IF EXISTS ais_account;
