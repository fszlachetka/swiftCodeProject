-- +goose Up
CREATE TABLE swift_codes (
    swift_code CHAR(11) PRIMARY KEY,
    countryISO2 CHAR(2) NOT NULL,
    bankName VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    country_name VARCHAR(255) NOT NULL,
    is_headquarter BOOLEAN NOT NULL
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE swift_codes;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
