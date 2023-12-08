DO
$do$
BEGIN
    IF NOT EXISTS(
        SELECT FROM pg_catalog.pg_roles
        where rolname = 'api_readonly')  then
        CREATE ROLE api_readonly LOGIN PASSWORD 'boms';
    END IF;

    IF NOT EXISTS(
            SELECT FROM pg_catalog.pg_roles
            where rolname = 'api_readwrite')  then
            CREATE ROLE api_readwrite LOGIN PASSWORD 'boms';
        END IF;
END

$do$;