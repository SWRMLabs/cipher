package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

func Decrypt(keyString, filename string) ([]byte, error) {
	key := []byte(keyString)
	ciphered, err := ReadFromFile(filename)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphered[:aes.BlockSize]
	ciphered = ciphered[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphered, ciphered)
	return ciphered, nil
}

func Encrypt(keyString, filename string) (string, error) {
	key := []byte(keyString)
	plaintext, err := ReadFromFile(filename)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return filename+"_encrypted", WriteToFile(string(ciphertext), filename+"_encrypted")
}

func WriteToFile(data, file string) error {
	return ioutil.WriteFile(file, []byte(data), 0755)
}

func ReadFromFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}
