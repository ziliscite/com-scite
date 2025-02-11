-- Drop foreign key relationships and the ComicGenre table
DROP INDEX IF EXISTS idx_comicgenre_genre;
DROP TABLE IF EXISTS ComicGenre;

-- Drop Genre table
DROP TABLE IF EXISTS Genre;

-- Drop CoverUrl related indexes
DROP INDEX IF EXISTS idx_cover_comic;
DROP INDEX IF EXISTS idx_current_cover;

-- Drop CoverUrl table
DROP TABLE IF EXISTS CoverUrl;

-- Drop UserBookmark related index
DROP INDEX IF EXISTS idx_bookmark_comic;

-- Drop Ratings related index
DROP INDEX IF EXISTS idx_ratings_comic;

-- Drop Ratings table
DROP TABLE IF EXISTS Ratings;

-- Drop ComicInteraction related indexes
DROP INDEX IF EXISTS idx_comic_interaction;
DROP INDEX IF EXISTS idx_interaction_time;

-- Drop ComicInteraction table
DROP TABLE IF EXISTS ComicInteraction;

-- Drop Comic table
DROP INDEX IF EXISTS idx_comic_status;
DROP INDEX IF EXISTS idx_comic_type;
DROP TABLE IF EXISTS Comic;

-- Drop Enum types
DROP TYPE IF EXISTS comic_status;
DROP TYPE IF EXISTS comic_type;
