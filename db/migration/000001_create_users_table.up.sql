DO $$
BEGIN
 
     IF NOT EXISTS( SELECT 1 FROM pg_type WHERE typname = 'gender_enum' )
          THEN
               CREATE TYPE gender_enum AS enum (
                    'male',
                    'female',
                    'undisclosed'
               );
     END IF;

     IF NOT EXISTS( SELECT * FROM information_schema.columns WHERE table_name='users' )
          THEN
               CREATE TABLE users (
                    id varchar not null primary key,
                    username varchar not null unique,
                    fullname varchar not null,
                    email varchar not null unique,
                    password varchar not null,
                    image varchar,
                    gender gender_enum not null default 'undisclosed',
                    dob timestamptz,
                    is_premium smallint,
                    created_at timestamptz not null default now(),
                    updated_at timestamptz not null default now(),
                    deleted_at timestamptz
               );
     END IF;

 END $$;