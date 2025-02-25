package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type Encryptor struct {
	key string
	iv  string
}

func NewEncryptor(key, iv string) *Encryptor {
	return &Encryptor{
		key: key,
		iv:  iv,
	}
}

func (en Encryptor) Decrypt(encrypted string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(en.key))
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(en.iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = en.padding(ciphertext)

	return ciphertext, nil
}

func (en Encryptor) padding(src []byte) []byte {
	length := len(src)
	unpad := int(src[length-1])

	return src[:(length - unpad)]
}

func (en Encryptor) Encrypt(plaintext string) (string, error) {
	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)

	block, err := aes.NewCipher([]byte(en.key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(en.iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
