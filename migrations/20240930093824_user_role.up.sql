CREATE TYPE user_role AS ENUM ('user', 'master', 'admin');

ALTER TABLE users
    DROP COLUMN is_master,
    ADD COLUMN role user_role NOT NULL DEFAULT 'user';
