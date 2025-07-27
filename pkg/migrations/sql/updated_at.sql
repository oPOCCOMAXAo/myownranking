CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW())::bigint;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

DO $$
DECLARE
    r RECORD;
BEGIN
    FOR r IN
    SELECT
        table_schema,
        table_name
    FROM
        information_schema.columns
    WHERE
        column_name = 'updated_at'
        AND table_schema = 'public' LOOP
            IF NOT EXISTS (
                SELECT
                    1
                FROM
                    pg_trigger
                WHERE
                    tgname = 'trigger_set_updated_at'
                    AND tgrelid = format('%I.%I', r.table_schema, r.table_name)::regclass) THEN
            EXECUTE format('CREATE TRIGGER trigger_set_updated_at BEFORE UPDATE ON %I.%I FOR EACH ROW EXECUTE FUNCTION set_updated_at();', r.table_schema, r.table_name);
        END IF;
END LOOP;
END;
$$
LANGUAGE plpgsql;

