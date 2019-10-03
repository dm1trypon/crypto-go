package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"log"
	"filesOperations"
)

func ToDecrypt(pathToFile string, key []byte, cryptoPrefix string) {
	if decrypted, err := decrypt(key, filesOperations.ReadFile(pathToFile)); err != nil {
		log.Println(err)
	} else {
		filesOperations.Clear(pathToFile)
		filesOperations.WriteFile(pathToFile, []byte(decrypted), cryptoPrefix)
	}
}

func decrypt(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)

	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}
