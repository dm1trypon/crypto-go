package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	path := `E:\PROJECT\Go\crypto-go\testfiles\`

	for i := 0; i < 1000; i++ {
		err := ioutil.WriteFile(path+"file"+strconv.Itoa(i), []byte("Data file"), 0644)

		if err != nil {
			log.Fatal(err)
		}
	}
}
