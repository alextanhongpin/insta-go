CREATE TABLE photos (
	photo_id SERIAL PRIMARY KEY,
	src text,
	caption text,
	user_id REFERENCES user,
	date_created date NOT NULL DEFAULT CURRENT_TIMESTAMP
	date_updated date NOT NULL DEFAULT CURRENT_TIMESTAMP
)