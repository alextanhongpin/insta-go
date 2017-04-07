CREATE TABLE users (
	user_id SERIAL PRIMARY KEY,
	username text,
	email text,
	salted_password text,
	first_name text,
	last_name text,
	last_ip text,
	date_created date NOT NULL DEFAULT CURRENT_DATE,
	date_updated date NOT NULL DEFAULT CURRENT_DATE
)
