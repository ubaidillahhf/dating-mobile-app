DO $$
BEGIN

    IF EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'subscriptions_index')
        THEN
            ALTER TABLE subscriptions DROP CONSTRAINT subscriptions_index;
    END IF;

    IF EXISTS( SELECT * FROM information_schema.columns WHERE table_name='subscriptions' )
        THEN
            DROP TABLE subscriptions;
    END IF;

    IF EXISTS( SELECT 1 FROM pg_type WHERE typname = 'subscription_status_enum' )
        THEN
            DROP TYPE subscription_status_enum;
    END IF;

END $$;
