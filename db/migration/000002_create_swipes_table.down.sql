DO $$
BEGIN

    IF EXISTS( SELECT 1 FROM pg_constraint WHERE conname = 'swipe_uq_index')
        THEN
            ALTER TABLE swipes DROP CONSTRAINT swipe_uq_index;
    END IF;

    IF EXISTS( SELECT * FROM information_schema.columns WHERE table_name='swipes' )
        THEN
            DROP TABLE swipes;
    END IF;

    IF EXISTS( SELECT 1 FROM pg_type WHERE typname = 'swipe_direction_enum' )
        THEN
            DROP TYPE swipe_direction_enum;
    END IF;

END $$;