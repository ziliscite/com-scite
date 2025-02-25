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
	return Cover{
		ComicID:   comicId,
		URL:       url,
		IsCurrent: true,
	}
}
