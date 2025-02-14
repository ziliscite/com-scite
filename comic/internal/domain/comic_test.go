package domain

import (
	"testing"
)

func TestNewComicStatus_Success(t *testing.T) {
	statusDummy := "Ongoing"
	stts, err := NewComicStatus(statusDummy)
	if err != nil {
		t.Errorf("Error creating new comic status: %v", err)
	}

	if Ongoing != stts {
		t.Errorf("NewComicStatus() returned %v; want %v", stts, Ongoing)
	}
}

func TestNewComicStatus_Failed(t *testing.T) {
	statusDummy := "nae nae nae"
	_, err := NewComicStatus(statusDummy)
	if err == nil {
		t.Errorf("Not error? creating new comic status")
	}
}

func TestNewComicType_Success(t *testing.T) {
	typeDummy := "Manga"
	stts, err := NewComicType(typeDummy)
	if err != nil {
		t.Errorf("Error creating new comic type: %v", err)
	}

	if Manga != stts {
		t.Errorf("NewComicType() returned %v; want %v", stts, Ongoing)
	}
}

func TestNewComicType_Failed(t *testing.T) {
	typeDummy := "nae nae nae"
	_, err := NewComicType(typeDummy)
	if err == nil {
		t.Errorf("Not error? creating new comic type")
	}
}

func TestNewComic_Success(t *testing.T) {

}
