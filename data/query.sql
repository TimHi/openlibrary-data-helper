-- name: GetRereading :one
SELECT * FROM reading
WHERE id = ? LIMIT 1;
-- name: ListReadings :many
SELECT * FROM reading;
-- name: InsertReading :exec
INSERT INTO reading (workkey, editionkey, shelf, datestamp) VALUES (?, ?, ?, ?);

-- name: GetRating :one
SELECT * FROM rating
WHERE id = ? LIMIT 1;
-- name: ListRatings :many
SELECT * FROM rating;
-- name: InsertRating :exec
INSERT INTO rating (workkey, editionkey, ratingvalue, datestamp) VALUES (?, ?, ?, ?);
