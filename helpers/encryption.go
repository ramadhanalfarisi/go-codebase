package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var (
	keyString = []byte("tokkoH22893uja294Jhjk9iioafae2s4")
)

func Encrypt(stringToEncrypt string) (encryptedString string) {
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(keyString)
	if err != nil {
		Error(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		Error(err)
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		Error(err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(encryptedString string) (decryptedString string) {
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(keyString)
	if err != nil {
		Error(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		Error(err)
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		Error(err)
	}

	return fmt.Sprintf("%s", plaintext)
}