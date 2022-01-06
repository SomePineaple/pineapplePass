package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
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

func AesEncryptBytes(plaintext []byte, key []byte) []byte {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln("(AesEncryptBytes): Failed to create new AES cipher, err:", err)
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalln("(AesEncryptBytes): Failed to create new gcm, err:", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatalln("(AesEncryptBytes): Failed to put random values in to nonce ???, err:", err)
	}

	return gcm.Seal(nonce, nonce, plaintext, nil)
}

func AesDecryptBytes(ciphertext []byte, key []byte) []byte {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln("(AesDecryptBytes): Failed to create new AES cipher, err:", err)
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		log.Fatalln("(AesDecryptBytes): Failed to create new gcm, err:", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalln("(AesDecryptBytes): Failed to decrypt ciphertext, err:", err)
	}

	return plaintext
}
