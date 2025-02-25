package service

import (
	"bytes"
	"fmt"
	"github.com/ziliscite/com-scite/object_storage/pkg/encryptor"
	"os"
)

type Saver interface {
	Save(filename, types string, imageData bytes.Buffer) (string, error)
}

type Getter interface {
	Get(signedUrl string) (string, error)
}

type ImageStore interface {
	Saver
	Getter
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

	file, err := os.Create("./store/" + imagePath)
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

	fullFilePath := "./store/" + string(filePath)

	return fullFilePath, nil
}
