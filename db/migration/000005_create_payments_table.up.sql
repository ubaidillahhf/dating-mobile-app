DO $$
BEGIN

    IF NOT EXISTS( SELECT 1 FROM pg_type WHERE typname = 'payments_status_enum' )
        THEN
            CREATE TYPE payments_status_enum AS enum (
                'waiting',
                'success',
                'failed'
            );
    END IF;

    IF NOT EXISTS( SELECT 1 FROM pg_type WHERE typname = 'payments_context_enum' )
        THEN
            CREATE TYPE payments_context_enum AS enum (
                'subscription'
            );
    END IF;

    IF NOT EXISTS( SELECT * FROM information_schema.columns WHERE table_name='payments' )
        THEN
            CREATE TABLE payments (
                id bigserial not null primary key,
                user_id varchar not null,
                ref_context payments_context_enum not null default 'subscription',
                ref_id varchar not null,
                amount decimal(16,2) not null,
                external_id varchar not null,
                method varchar not null,
                status payments_status_enum not null default 'waiting',
                created_at timestamptz not null default now(),
                updated_at timestamptz not null default now(),
                deleted_at timestamptz
            );
    END IF;

    IF NOT EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'payments_index')
        THEN
            CREATE INDEX "payments_index" on payments using btree (status);
    END IF;

END $$;