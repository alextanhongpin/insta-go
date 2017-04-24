CREATE TABLE likes (
	user_id integer REFERENCES users (user_id),
	photo_id integer REFERENCES photos (photo_id),
	date_created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
	date_updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
)
