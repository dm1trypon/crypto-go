package worker

import (
	"decrypt"
	"encrypt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func Operations(wg *sync.WaitGroup, typeOperation string, pathToFiles string, CIPHER_KEY []byte, cryptoPrefix string, hackMode bool, isAsync bool) {
	if _, err := os.Stat(pathToFiles); os.IsNotExist(err) {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(pathToFiles)

	if err != nil {
		switch typeOperation {
		case "encrypt":
			if isAsync {
				wg.Add(1)
				go encrypt.ToEncrypt(wg, pathToFiles, CIPHER_KEY, cryptoPrefix, hackMode, isAsync)
				break
			}
			encrypt.ToEncrypt(wg, pathToFiles, CIPHER_KEY, cryptoPrefix, hackMode, isAsync)
		case "decrypt":
			if isAsync {
				wg.Add(1)
				go decrypt.ToDecrypt(wg, pathToFiles, CIPHER_KEY, cryptoPrefix, isAsync)
				break
			}
			decrypt.ToDecrypt(wg, pathToFiles, CIPHER_KEY, cryptoPrefix, isAsync)
		default:
			log.Fatal("Unknown type!")
		}

		return
	}

	for _, file := range files {
		if file.IsDir() {
			Operations(wg, typeOperation, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, hackMode, isAsync)
			continue
		}

		switch typeOperation {
		case "encrypt":
			if isAsync {
				wg.Add(1)
				go encrypt.ToEncrypt(wg, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, hackMode, isAsync)
				continue
			}

			encrypt.ToEncrypt(wg, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, hackMode, isAsync)
			
		case "decrypt":
			if isAsync {
				wg.Add(1)
				go decrypt.ToDecrypt(wg, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, isAsync)
				continue
			}
			
			decrypt.ToDecrypt(wg, pathToFiles+"/"+file.Name(), CIPHER_KEY, cryptoPrefix, isAsync)

		default:
			log.Fatal("Unknown type!")
		}
	}
}
