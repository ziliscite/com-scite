DROP INDEX IF EXISTS idx_comicgenre_genre;
DROP TABLE IF EXISTS ComicGenre;

DROP TABLE IF EXISTS Genre;

DROP INDEX IF EXISTS idx_cover_comic;
DROP INDEX IF EXISTS idx_current_cover;

DROP TABLE IF EXISTS CoverUrl;

DROP INDEX IF EXISTS idx_bookmark_comic;

DROP INDEX IF EXISTS idx_ratings_comic;

DROP TABLE IF EXISTS Ratings;

DROP INDEX IF EXISTS idx_comic_interaction;
DROP INDEX IF EXISTS idx_interaction_time;

DROP TABLE IF EXISTS ComicInteraction;

DROP INDEX IF EXISTS idx_comic_status;
DROP INDEX IF EXISTS idx_comic_type;
DROP TABLE IF EXISTS Comic;

DROP TYPE IF EXISTS comic_status;
DROP TYPE IF EXISTS comic_type;
