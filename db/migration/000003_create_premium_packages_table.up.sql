CREATE TABLE IF NOT EXISTS premium_packages (
    id bigserial not null primary key,
    name varchar not null,
    description text,
    price decimal(16,2) not null,
    duration_in_days int not null,
    repeat smallint not null default 0,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);

-- insert data
INSERT INTO premium_packages (name,description,price,duration_in_days,repeat) VALUES 
('A Week','One time purchase for 1 weeks', 10, 7, 0), 
('Subscribe Daily','Daily subscription', 1, 1, 1),
('Subscribe Weekly','Weekly subscription, save $1!', 6, 7, 1),
('Subscribe Monthly','Monthly subscription, save $5!', 25, 30, 1), 
('Subscribe Yearly','Yearly subscription, save $65!', 300, 365, 1);