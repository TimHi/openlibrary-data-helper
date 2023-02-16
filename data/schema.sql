CREATE TABLE IF NOT EXISTS reading (
    id INTEGER PRIMARY KEY,
    workkey text NOT NULL,
    editionkey text,
    shelf text,
    datestamp text
);

CREATE TABLE IF NOT EXISTS rating (
    id INTEGER PRIMARY KEY,
    workkey text NOT NULL,
    editionkey text,
    ratingvalue REAL,
    datestamp text
);