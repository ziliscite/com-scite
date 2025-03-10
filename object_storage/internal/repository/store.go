package repository

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ziliscite/com-scite/object_storage/pkg/encryptor"
	"os"
	"path/filepath"
)

type Write interface {
	Save(filename, types string, imageData bytes.Buffer) (string, error)
	Delete(signedUrl string) error
}

type Read interface {
	Get(signedUrl string) (string, error)
}

type ImageStore interface {
	Write
	Read
}

type store struct {
	en *encryptor.Encryptor
}

func NewStore(en *encryptor.Encryptor) ImageStore {
	return &store{
		en: en,
	}
}

func (s *store) Save(filename, types string, imageData bytes.Buffer) (string, error) {
	imagePath := fmt.Sprintf("%s/%s", types, filename)

	signedUrl, err := s.en.Encrypt(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot encrypt image url: %w", err)
	}

	filePath := s.createFilePath(imagePath)
	if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return "", fmt.Errorf("cannot create directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	return signedUrl, nil
}

func (s *store) Get(signedUrl string) (string, error) {
	filePath, err := s.en.Decrypt(signedUrl)
	if err != nil {
		return "", fmt.Errorf("cannot decrypt image url: %w", err)
	}

	return s.createFilePath(string(filePath)), nil
}

func (s *store) Delete(signedUrl string) error {
	// first insert
	if signedUrl == "" {
		return nil
	}
	
	fileString, err := s.en.Decrypt(signedUrl)
	if err != nil {
		return fmt.Errorf("cannot decrypt image url: %w", err)
	}

	filePath := s.createFilePath(string(fileString))

	if err = os.Remove(filePath); err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return fmt.Errorf("cannot delete image file: %w", err)
		}
	}

	return nil
}

func (s *store) createFilePath(filePath string) string {
	return "./store/" + filePath
}
