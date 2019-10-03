package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const COUNT_OF_REQ = 10

func getGooglePage(waiter chan int, i int) {
	fmt.Println("start: ", i)
	resp, _ := http.Get("https://google.com")

	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
	fmt.Println("end: ", i)

	waiter <- i
}

type ServerHandler struct{}

func (sh ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		waiter := make(chan int)
		count := 0

		for i := 0; i < COUNT_OF_REQ; i++ {
			go getGooglePage(waiter, i)
		}

	Loop:
		for {
			select {
			case msg := <-waiter:
				count++
				fmt.Println(msg)
			case <-time.After(500 * time.Millisecond):
				fmt.Println("Timeout. Cancel")
				break Loop
			}
		}

		fmt.Println(
			fmt.Sprintf("Count of request before cancel %d: ", count))

		elapsed := time.Since(start)

		fmt.Fprintf(w, elapsed.String())
	})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not Found")
	})

	http.ListenAndServe(":80", h)
}
