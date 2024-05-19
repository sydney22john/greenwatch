-- +goose Up
CREATE TABLE greenwashers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company_name TEXT NOT NULL UNIQUE
);

CREATE TABLE domains (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    domain TEXT NOT NULL UNIQUE,

    greenwasher_id INTEGER REFERENCES greenwashers (id) ON DELETE CASCADE
);

CREATE TABLE offenses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    offence TEXT NOT NULL,

    greenwasher_id INTEGER REFERENCES greenwashers (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE greenwashers;
DROP TABLE domains;
DROP TABLE offenses;
