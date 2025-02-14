-- Comic domain
CREATE TYPE comic_status AS ENUM (
    'Ongoing', 'Completed', 'Dropped',
    'Hiatus', 'Coming Soon', 'Season End'
);

CREATE TYPE comic_type AS ENUM (
    'Manga', 'Manhwa', 'Manhua'
);

CREATE TABLE Comic (
    comic_id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    author VARCHAR(255),
    artist VARCHAR(255),
    status comic_status NOT NULL,
    type comic_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    version integer NOT NULL DEFAULT 1
);

CREATE INDEX idx_comic_status ON Comic(status);
CREATE INDEX idx_comic_type ON Comic(type);

-- Cover domain
CREATE TABLE Cover (
    cover_id BIGSERIAL PRIMARY KEY,
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    url VARCHAR(2048) UNIQUE NOT NULL,
    is_current BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_current_cover ON Cover(comic_id) WHERE is_current = TRUE;
CREATE INDEX idx_cover_comic ON Cover(comic_id);

-- Genre domain
CREATE TABLE Genre (
    genre_id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE Comic_Genre (
    comic_id BIGINT NOT NULL REFERENCES Comic(comic_id) ON DELETE CASCADE,
    genre_id INT NOT NULL REFERENCES Genre(genre_id) ON DELETE CASCADE,
    PRIMARY KEY (comic_id, genre_id)
);

CREATE INDEX idx_comic_genre_genre ON Comic_Genre(genre_id);
CREATE INDEX idx_comic_updated_at ON Comic(updated_at);
