CREATE TABLE photo_tags (
	photo_id text references photos
	tag_id text references tags
)