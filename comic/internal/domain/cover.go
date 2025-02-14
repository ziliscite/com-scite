package domain

import "time"

type Cover struct {
	ID        int64
	ComicID   int64
	URL       string
	IsCurrent bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCover(comicId int64, url string) Cover {
	// Some url things with AWS S3 will be done on the handler, then it'll be sent here
	return Cover{
		ComicID: comicId,
		URL:     url,
		// Newly created cover will be displayed
		IsCurrent: true,
	}
}
