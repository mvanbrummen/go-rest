CREATE SCHEMA imdb;

CREATE TABLE imdb.titles (
	tconst CHAR(20) PRIMARY KEY,
	titleType CHAR(50),
	primaryTitle VARCHAR(500),
	originalTitle VARCHAR(500),
	isAdult BOOLEAN,
	startYear INTEGER,
	endYear INTEGER,
	runtimeMinutes INTEGER,
	genres VARCHAR(250)
);

CREATE ROLE imdb_ro WITH LOGIN ENCRYPTED PASSWORD 'Password@1';
GRANT USAGE ON SCHEMA imdb TO imdb_ro; 
GRANT SELECT ON ALL TABLES IN SCHEMA imdb TO imdb_ro;