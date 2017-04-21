# insta-go
Mock API Instagram with Golang

- [ ] Config best practices
- [ ] Connect to database
- [ ] Designing models
- [ ] Chaining middlewares
- [ ] Deploy using docker
- [ ] JWT authentication
- [ ] Testing each layers individually


```
$ psql

// List down all database
$ \l 
// Create a new database with the name instadb;
CREATE DATABASE instadb;

// Connect to the database
$ \c instadb;

// Create a new table
CREATE TABLE photo_tbl (
    photo_id INTEGER NOT NULL,
    src TEXT,
    CONSTRAINT "PRIM_KEY" PRIMARY KEY (photo_id)

    FOREIGN_KEY (user_id)
    REFERENCES user_tbl (user_id)
)

// Create another table
CREATE TABLE user_tbl (
    user_id INTEGER NOT NULL,
)

// View list of tables
$ \d

// View specific table
$ \d photo_tbl

// Alter the table by adding columns
ALTER TABLE photo_tbl
ADD COLUMN alt TEXT;


// Clear the table
TRUNCATE TABLE photo_tbl;

// Inserting data
INSERT INTO photo_tbl
(photo_id, src) VALUES (1, 'SOMETHING');

// Query the sql
SELECT * FROM photo_tbl;
```