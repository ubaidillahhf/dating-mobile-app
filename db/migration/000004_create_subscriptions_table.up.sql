DO $$
BEGIN

    IF NOT EXISTS( SELECT 1 FROM pg_type WHERE typname = 'subscription_status_enum' )
        THEN
            CREATE TYPE subscription_status_enum AS enum (
                'active',
                'grace_period',
                'end',
                'pending',
                'cancel'
            );
    END IF;

    IF NOT EXISTS( SELECT * FROM information_schema.columns WHERE table_name='subscriptions' )
        THEN
            CREATE TABLE subscriptions (
                id bigserial not null primary key,
                user_id varchar not null,
                premium_packages_id bigint not null,
                status subscription_status_enum not null default 'pending',
                end_at timestamptz not null,
                created_at timestamptz not null default now(),
                updated_at timestamptz not null default now(),
                deleted_at timestamptz
            );
    END IF;

    IF NOT EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'subscriptions_index')
        THEN
            CREATE INDEX "subscriptions_index" on subscriptions using btree (status);
    END IF;

END $$;