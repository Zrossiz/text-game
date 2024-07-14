CREATE TABLE IF NOT EXISTS command (
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS location (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS session (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    location_id INTEGER,
    FOREIGN KEY (location_id) REFERENCES location(id)
);

CREATE TABLE IF NOT EXISTS item (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    location_id INTEGER,
    FOREIGN KEY (location_id) REFERENCES location(id)
);

CREATE TABLE IF NOT EXISTS session_item (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    session_id INTEGER,
    FOREIGN KEY (session_id) REFERENCES session(id)
);
