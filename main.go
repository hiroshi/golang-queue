package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func main() {
	// fmt.Printf("Hello, world.\n")
	messages := make(chan string, 100)

	go func () {
		for {
			select {
			case msg := <- messages:
				fmt.Println("msg: ", msg)
				time.Sleep(3 * time.Second)
				fmt.Println("done.")
			// default:
			// 	fmt.Println("no msg")
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		messages <- "hi"
		fmt.Fprintf(w, "ok!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
