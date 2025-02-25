DROP INDEX IF EXISTS idx_comic_genre_genre;
DROP TABLE IF EXISTS Comic_Genre;

DROP INDEX IF EXISTS idx_cover_comic;
DROP INDEX IF EXISTS idx_current_cover;

DROP TABLE IF EXISTS CoverUrl;

DROP INDEX IF EXISTS idx_current_cover;
DROP INDEX IF EXISTS idx_cover_comic;

DROP TABLE IF EXISTS Cover;

DROP INDEX IF EXISTS idx_comic_status;
DROP INDEX IF EXISTS idx_comic_type;
DROP TABLE IF EXISTS Comic CASCADE ;

DROP TYPE IF EXISTS comic_status;
DROP TYPE IF EXISTS comic_type;

DROP TABLE IF EXISTS Genre;