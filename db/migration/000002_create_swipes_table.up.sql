DO $$
BEGIN

    IF NOT EXISTS( SELECT 1 FROM pg_type WHERE typname = 'swipe_direction_enum' )
        THEN
            CREATE TYPE swipe_direction_enum AS enum (
                'left',
                'right'
            );
    END IF;

   IF NOT EXISTS( SELECT * FROM information_schema.columns WHERE table_name='swipes' )
        THEN
            CREATE TABLE swipes (
                id bigserial not null primary key,
                sender_id varchar not null,
                receiver_id varchar not null,
                direction swipe_direction_enum not null,
                created_at timestamptz not null default now(),
                deleted_at timestamptz
            );
    END IF;

    IF NOT EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'swipe_uq_index')
        THEN
            CREATE UNIQUE INDEX "swipe_uq_index" ON swipes(sender_id,receiver_id);
    END IF;

END $$;


