-- Drop ComicChapter table (mapping between comic and chapter)
DROP INDEX IF EXISTS idx_comic_chapter;
DROP TABLE IF EXISTS ComicChapter;

-- Drop ChapterPage table (pages for chapters)
DROP INDEX IF EXISTS idx_chapter_page;
DROP INDEX IF EXISTS idx_page_number;
DROP TABLE IF EXISTS ChapterPage;

-- Drop Chapter table (stores chapters)
DROP INDEX IF EXISTS idx_chapter_comic;
DROP INDEX IF EXISTS idx_chapter_number;
DROP TABLE IF EXISTS Chapter;
