package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
	"log"
)

func GeneratePasswordSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatalln("Failed to generate salt, err:", err)
	}

	return salt
}

func GenerateAesKey(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

func AesEncryptBytes(toEncrypt []byte, key []byte) []byte {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln("Failed to create new AES cipher, err:", err)
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalln("Failed to create new gcm, err:", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatalln("Failed to put random values in to nonce ???, err:", err)
	}

	return gcm.Seal(nonce, nonce, toEncrypt, nil)
}
