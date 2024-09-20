CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    second_name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    is_master boolean DEFAULT false
);

CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL
);

CREATE TABLE masters (
    id SERIAL PRIMARY KEY,
    user_id integer UNIQUE,
    position_id integer,
    bio text NOT NULL,
    slug varchar(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (position_id) REFERENCES positions (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description text NOT NULL,
    slug varchar(255) NOT NULL
);

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    category_id integer,
    title varchar(255) NOT NULL,
    duration time NOT NULL,
    price decimal (6, 2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT valid_price CHECK(price > 0)
);

CREATE TABLE appointments (
	id SERIAL PRIMARY KEY,
	appointment_start   timestamp NOT NULL,
	appointment_end      timestamp NOT NULL,
	user_id   integer,
	master_id integer,
	status varchar(255) NOT NULL,
	comment text,
	total_sum decimal(7, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (master_id) REFERENCES masters (id) ON DELETE RESTRICT,
    CONSTRAINT valid_total_sum CHECK(total_sum > 0),
    CONSTRAINT valid_time CHECK(appointment_end > appointment_start)
);