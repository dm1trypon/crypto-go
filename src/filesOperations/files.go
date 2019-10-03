package filesOperations

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	s "strings"
)

func ReadFile(path string) string {
	contents, _ := ioutil.ReadFile(path)
	return string(contents)
}

func WriteFile(pathToFile string, data []byte, cryptoPrefix string) {
	var path string

	if s.HasSuffix(pathToFile, cryptoPrefix) {
		path = strings.Replace(pathToFile, "."+cryptoPrefix, "", -1)
		log.Printf("DECRYPT: %s ---> %s", pathToFile, path)
	} else {
		path = pathToFile + "." + cryptoPrefix
		log.Printf("ENCRYPT: %s ---> %s", pathToFile, path)
	}

	err := ioutil.WriteFile(path, data, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func Clear(pathToFile string) {
	err := os.Remove(pathToFile)

	if err != nil {
		log.Println(err)
		return
	}
}
