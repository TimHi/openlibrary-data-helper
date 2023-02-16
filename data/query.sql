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
-- name: ListTop100 :many
SELECT *
FROM rating
WHERE workkey IN (
  SELECT workkey
  FROM reading
  WHERE shelf = 'Already Read'
  GROUP BY workkey
  HAVING COUNT(*) > 100
)
ORDER BY ratingvalue DESC
LIMIT 100;
