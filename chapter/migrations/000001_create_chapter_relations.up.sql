CREATE TABLE Chapter (
    chapter_id BIGSERIAL PRIMARY KEY,
    comic_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    chapter_number INT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_chapter_comic ON Chapter(comic_id);
CREATE INDEX idx_chapter_number ON Chapter(chapter_number);

CREATE TABLE ChapterPage (
    page_id BIGSERIAL PRIMARY KEY,
    chapter_id BIGINT NOT NULL REFERENCES Chapter(chapter_id) ON DELETE CASCADE,
    page_number INT NOT NULL,
    url VARCHAR(2048) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_chapter_page ON ChapterPage(chapter_id);
CREATE INDEX idx_page_number ON ChapterPage(page_number);

