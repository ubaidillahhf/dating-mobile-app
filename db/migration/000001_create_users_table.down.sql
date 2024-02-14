DO $$
BEGIN

    IF EXISTS( SELECT * FROM information_schema.columns WHERE table_name='users' )
        THEN
            DROP TABLE users;
    END IF;

    IF EXISTS( SELECT 1 FROM pg_type WHERE typname = 'gender_enum' )
        THEN
            DROP TYPE gender_enum;
    END IF;

END $$;