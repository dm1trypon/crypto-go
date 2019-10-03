package worker

import (
	"decrypt"
	"encrypt"
	"io/ioutil"
	"log"
	"os"
)

func Operations(typeOperation string, pathToFiles string, CIPHER_KEY []byte, cryptoPrefix string, hackMode bool) {
	if _, err := os.Stat(pathToFiles); os.IsNotExist(err) {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(pathToFiles)

	if err != nil {
		switch typeOperation {
		case "encrypt":
			encrypt.ToEncrypt(pathToFiles, CIPHER_KEY, cryptoPrefix, hackMode)
		case "decrypt":
			decrypt.ToDecrypt(pathToFiles, CIPHER_KEY, cryptoPrefix)
		default:
			log.Fatal("Unknown type!")
		}

		return
	}

	for _, file := range files {
		if file.IsDir() {
			Operations(typeOperation, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, hackMode)
			continue
		}

		switch typeOperation {
		case "encrypt":
			encrypt.ToEncrypt(pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, hackMode)
		case "decrypt":
			decrypt.ToDecrypt(pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix)
		default:
			log.Fatal("Unknown type!")
		}
	}
}
