CREATE TABLE IF NOT EXISTS users (
     id varchar not null primary key,
     username varchar not null,
     fullname varchar not null,
     email varchar not null,
     password varchar not null,
     created_at timestamptz not null default now(),
     updated_at timestamptz not null default now(),
     deleted_at timestamptz
);