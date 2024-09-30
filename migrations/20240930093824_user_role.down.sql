ALTER TABLE users
    DROP COLUMN role;

DROP TYPE user_role;

ALTER TABLE users
    ADD COLUMN is_master BOOLEAN NOT NULL DEFAULT false;
