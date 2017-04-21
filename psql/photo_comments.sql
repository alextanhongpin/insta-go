CREATE TABLE photo_comments {
	photo_id text REFERENCES photos
	comment_id text REFERENCES comments
}