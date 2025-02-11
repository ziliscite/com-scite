-- Enum types for comic status and type
CREATE TYPE comic_status AS ENUM (
    'Ongoing',
    'Completed',
    'Dropped',
    'Hiatus',
    'Coming Soon',
    'Season End'
);

CREATE TYPE comic_type AS ENUM (
    'Manga',
    'Manhwa',
    'Manhua'
);

-- Comic table to store basic comic details
CREATE TABLE Comic (
    comic_id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,  -- Increased length for title
    description TEXT,
    author VARCHAR(255),
    artist VARCHAR(255),
    status comic_status NOT NULL,
    type comic_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes to speed up queries on comic status and type
CREATE INDEX idx_comic_status ON Comic(status);
CREATE INDEX idx_comic_type ON Comic(type);

-- Comic interaction tracking table
CREATE TABLE ComicInteraction (
interaction_id BIGSERIAL PRIMARY KEY,
comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
event_type VARCHAR(20) NOT NULL CHECK (
        event_type IN ('view', 'follow', 'unfollow', 'share')
    ),
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for fast lookups of interactions by event type and time
CREATE INDEX idx_comic_interaction ON ComicInteraction(comic_id, event_type, created_at);
CREATE INDEX idx_interaction_time ON ComicInteraction(created_at);

-- Ratings table to track user ratings for comics
CREATE TABLE Ratings (
    user_id BIGINT NOT NULL,  -- Reference to external user service
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    rating SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    rated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, comic_id)
);

-- Index to speed up comic-related ratings lookups
CREATE INDEX idx_ratings_comic ON Ratings(comic_id);

-- UserBookmark table to track user bookmarks for comics
CREATE TABLE UserBookmark (
    user_id BIGINT NOT NULL,
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    bookmarked_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, comic_id)
);

-- Index for fast lookups on bookmarked comics
CREATE INDEX idx_bookmark_comic ON UserBookmark(comic_id);

-- Cover URLs for comics
CREATE TABLE CoverUrl (
    cover_id BIGSERIAL PRIMARY KEY,
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    url VARCHAR(2048) NOT NULL,
    version INT NOT NULL DEFAULT 1,
    is_current BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Ensure only one current cover per comic
CREATE UNIQUE INDEX idx_current_cover ON CoverUrl(comic_id) WHERE is_current = TRUE;
CREATE INDEX idx_cover_comic ON CoverUrl(comic_id);

-- Genre table to store comic genres
CREATE TABLE Genre (
    genre_id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- ComicGenre to relate comics with their genres
CREATE TABLE ComicGenre (
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    genre_id INT NOT NULL REFERENCES Genre(genre_id) ON DELETE CASCADE,
    PRIMARY KEY (comic_id, genre_id)
);

-- Index for fast lookups of comics by genre
CREATE INDEX idx_comicgenre_genre ON ComicGenre(genre_id);

-- Ensure efficient querying for updates by timestamp
CREATE INDEX idx_comic_updated_at ON Comic(updated_at);
