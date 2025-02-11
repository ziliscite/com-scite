-- Chapter Table to store chapter details (without foreign key to Comic)
CREATE TABLE Chapter (
    chapter_id BIGSERIAL PRIMARY KEY,
    comic_id BIGINT NOT NULL,  -- Regular column for comic_id
    title VARCHAR(255) NOT NULL,
    chapter_number INT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for fast lookups of chapters by comic
CREATE INDEX idx_chapter_comic ON Chapter(comic_id);
CREATE INDEX idx_chapter_number ON Chapter(chapter_number);

-- ChapterPage Table to store chapter pages
CREATE TABLE ChapterPage (
    page_id BIGSERIAL PRIMARY KEY,
    chapter_id BIGINT NOT NULL REFERENCES Chapter(chapter_id) ON DELETE CASCADE,
    page_number INT NOT NULL,
    url VARCHAR(2048) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for fast lookups of chapter pages by chapter
CREATE INDEX idx_chapter_page ON ChapterPage(chapter_id);
CREATE INDEX idx_page_number ON ChapterPage(page_number);

