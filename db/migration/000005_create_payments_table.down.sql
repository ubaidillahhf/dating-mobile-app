DO $$
BEGIN

    IF EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'payments_index')
        THEN
            ALTER TABLE payments DROP CONSTRAINT payments_index;
    END IF;

    IF EXISTS( SELECT * FROM information_schema.columns WHERE table_name='payments' )
        THEN
            DROP TABLE payments;
    END IF;

    IF EXISTS( SELECT 1 FROM pg_type WHERE typname = 'payments_status_enum' )
        THEN
            DROP TYPE payments_status_enum;
    END IF;

    IF EXISTS( SELECT 1 FROM pg_type WHERE typname = 'payments_context_enum' )
        THEN
            DROP TYPE payments_context_enum;
    END IF;

END $$;