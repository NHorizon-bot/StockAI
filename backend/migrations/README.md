# Database Migrations

V0.1 has no business tables yet. Versioned migration files belong in this
directory and should follow the convention below once the first schema is
introduced:

- `NNNNNN_name.up.sql`
- `NNNNNN_name.down.sql`

The backend reads the migration directory from `MYSQL_MIGRATIONS_DIR`. Its
default value is `./migrations` when the server is started from `backend/`.

