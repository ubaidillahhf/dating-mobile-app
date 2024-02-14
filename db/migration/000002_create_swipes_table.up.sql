CREATE TABLE IF NOT EXISTS swipes (
    id bigserial not null primary key,
    sender_id varchar not null,
    receiver_id varchar not null,
    direction smallint not null,
    created_at timestamptz not null default now(),
    deleted_at timestamptz
)