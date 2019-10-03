package main

import (
	"filesOperations"
	"log"
	"os"
	"worker"
)

func start(args []string, hackMode bool) {
	pathToFiles := args[1]
	typeOperation := args[2]
	cryptoKey := args[3]
	cryptoPrefix := args[4]

	if _, err := os.Stat(cryptoKey); os.IsNotExist(err) {
		log.Fatal("Decryption's key is not found!")
	}

	CIPHER_KEY := []byte(filesOperations.ReadFile(cryptoKey))

	if len(CIPHER_KEY) != 16 {
		log.Fatal("Line length must be 16 characters!")
	}

	worker.Operations(typeOperation, pathToFiles, CIPHER_KEY, cryptoPrefix, hackMode)
}

func isHelp(args []string) bool {
	return len(args) < 3 && len(args) > 1 && args[1] == "help"
}

func isHack(args []string) bool {
	return len(args) > 5 && args[5] == "hack"
}

func showHelp() {
	log.Println("\n\n --------- ARGUMENTS --------- \n",
		"1) Path to file or directory for encrypt/decrypt.(REQUIRED)\n",
		"2) Select type operation: encrypt or decrypt.(REQUIRED)\n",
		"3) Path to key.(REQUIRED)\n",
		"4) Prefix for append/remove to name of file for encrypt/decrypt.(REQUIRED)\n",
		"5) Parameter that determines whether to delete the source files(flag 'hack') or not.\n",
		"\n\nExample:\n",
		"1) encrypt: go run main.go /path_to_folder_or_file/ encrypt crypto.key mask_prefix hack\n",
		"2) decrypt: go run main.go /path_to_folder_or_file/ decrypt crypto.key mask_prefix")
}

func main() {
	args := os.Args

	if isHelp(args) {
		showHelp()
		return
	}

	if len(args) < 5 {
		log.Fatal("Invalid args, try help!")
	}

	hackMode := false

	if isHack(args) {
		log.Println("Hack mode enabled =)")
		hackMode = true
	}

	start(args, hackMode)
}
