package domain

import (
	"fmt"
	"github.com/gosimple/slug"
	"time"
)

type ComicStatus int

const (
	Ongoing ComicStatus = iota
	Completed
	Dropped
	Hiatus
	ComingSoon
	SeasonEnd
)

func (cs ComicStatus) String() string {
	return [...]string{"Ongoing", "Completed", "Dropped", "Hiatus", "Coming Soon", "Season End"}[cs]
}

func NewComicStatus(status string) (ComicStatus, error) {
	if cs, exists := map[string]ComicStatus{
		"Ongoing":     Ongoing,
		"Completed":   Completed,
		"Dropped":     Dropped,
		"Hiatus":      Hiatus,
		"Coming Soon": ComingSoon,
		"Season End":  SeasonEnd,
	}[status]; exists {
		return cs, nil
	}

	return 0, fmt.Errorf("invalid status: %s", status)
}

type ComicType int

const (
	Manga ComicType = iota
	Manhwa
	Manhua
)

func (ct ComicType) String() string {
	return [...]string{"Manga", "Manhwa", "Manhua"}[ct]
}

func NewComicType(status string) (ComicType, error) {
	if cs, exists := map[string]ComicType{
		"Manga":  Manga,
		"Manhwa": Manhwa,
		"Manhua": Manhua,
	}[status]; exists {
		return cs, nil
	}

	return 0, fmt.Errorf("invalid status: %s", status)
}

type Comic struct {
	ID          int64
	Title       string
	Slug        string
	Description string
	Author      string
	Artist      string
	Status      ComicStatus
	Type        ComicType

	Genres   []string // Will be inserted separately
	CoverUrl string   // Will also be inserted separately

	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int64
}

func NewComic(title, description, author, artist, comicStatus, comicType string) (*Comic, error) {
	stats, err := NewComicStatus(comicStatus)
	if err != nil {
		return nil, err
	}

	types, err := NewComicType(comicType)
	if err != nil {
		return nil, err
	}

	return &Comic{
		Title:       title,
		Slug:        slug.Make(title),
		Description: description,
		Author:      author,
		Artist:      artist,
		Status:      stats,
		Type:        types,
	}, nil
}
