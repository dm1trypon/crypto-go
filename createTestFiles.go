package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"os"
)

func generingText(count int) string {
	var str string

	for word := 0; word < count; word ++ {
		str += "test"
	}

	return str
}

func main() {
	args := os.Args

	if len(args) != 4 {
		log.Fatal("Invalid args!")
	}

	path := args[1]
	fileName := args[2]
	count, _ := strconv.Atoi(args[3])
	data := generingText(100000)

	for i := 0; i < count; i++ {
		err := ioutil.WriteFile(path+fileName+strconv.Itoa(i), []byte(data), 0644)

		if err != nil {
			log.Fatal(err)
		}
	}
}
