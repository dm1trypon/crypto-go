package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	path := "/home/dmitry/Desktop/11"

	for i := 0; i < 10000; i++ {
		err := ioutil.WriteFile(path+"/"+"file"+strconv.Itoa(i), []byte("Data file"), 0644)

		if err != nil {
			log.Fatal(err)
		}
	}
}
