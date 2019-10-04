package main

import (
	"filesOperations"
	"log"
	"os"
	"sync"
	"worker"
	"flag"
)

func start(pathToFiles string, isEncrypt bool, isDecrypt bool, pathToKey string, namePrefix string, isHack bool, isAsync bool) {
	if _, err := os.Stat(pathToKey); os.IsNotExist(err) {
		log.Fatal("Decryption's key is not found!")
	}

	CIPHER_KEY := []byte(filesOperations.ReadFile(pathToKey))

	if len(CIPHER_KEY) != 16 {
		log.Fatal("Line length must be 16 characters!")
	}

	if isAsync {
		log.Println("Program execution asynchronously")
	}

	var typeOperation string

	if isEncrypt {
		typeOperation = "encrypt"
	}

	if isDecrypt {
		typeOperation = "decrypt"
	}

	var wg sync.WaitGroup

	worker.Operations(&wg, typeOperation, pathToFiles, CIPHER_KEY, namePrefix, isHack, isAsync)

	if !isAsync {
		return
	}
	
	wg.Wait()
}

func main() {
	pathToFiles := flag.String("path", "", "Path to file or directory for encrypt/decrypt.(REQUIRED)")
	isEncrypt := flag.Bool("encrypt", false, "Select type operation: encrypt. (REQUIRED)")
	isDecrypt := flag.Bool("decrypt", false, "Select type operation: decrypt.(REQUIRED)")
	pathToKey := flag.String("key", "", "Path to key.(REQUIRED)")
	namePrefix := flag.String("prefix", "", "Prefix for append/remove to name of file for encrypt/decrypt.(REQUIRED)")
	isHack := flag.Bool("hack", false, "Parameter that determines whether to delete the source files(flag 'hack') or not.")
	isAsync := flag.Bool("async", false, "Run app as async.")

	flag.Parse()

	if *pathToFiles == "" {
		log.Fatal("Missing parameter 'path'!")
	}

	if *isEncrypt == *isDecrypt {
		log.Fatal("Wrong parameter 'encrypt' or 'decrypt'!")
	}

	if *pathToKey == "" {
		log.Fatal("Missing parameter 'key'!")
	}

	if *namePrefix == "" {
		log.Fatal("Missing parameter 'prefix'!")
	}

	start(*pathToFiles, *isEncrypt, *isDecrypt, *pathToKey, *namePrefix, *isHack, *isAsync)

	log.Println("Done!")
}
