package domain

import (
	"errors"
	"sync"
	"unicode"
)

func invalidGenre(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsNumber(r) || !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

type Genre struct {
	ID   int64
	Name string
}

func NewGenre(name string) (*Genre, error) {
	if name == "" {
		return nil, errors.New("genre must not be empty")
	}

	if invalidGenre(name) {
		return nil, errors.New("genre names must not contain any numbers or special characters")
	}

	return &Genre{
		Name: name,
	}, nil
}

func NewMassGenre(names []string) ([]*Genre, error) {
	uniqueValues := make(map[string]bool)

	var mu sync.Mutex

	mu.Lock()
	for _, value := range names {
		uniqueValues[value] = true
	}
	mu.Unlock()

	if len(names) != len(uniqueValues) {
		return nil, errors.New("genre names must be unique")
	}

	genres := make([]*Genre, 0)
	for _, name := range names {
		gen, err := NewGenre(name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, gen)
	}

	return genres, nil
}
