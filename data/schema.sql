CREATE TABLE reading (
    id INTEGER PRIMARY KEY,
    workkey text NOT NULL,
    editionkey text,
    shelf text,
    datestamp text
);

CREATE TABLE rating (
    id INTEGER PRIMARY KEY,
    workkey text NOT NULL,
    editionkey text,
    ratingvalue INTEGER,
    datestamp text
);