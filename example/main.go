package main

import (
	"fmt"
	"log"

	"github.com/benebobaa/hatetepe"
)

type HelloWorld struct {
	Message string
}

func main() {

	// Single Route
	// handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello brow"))
	// })

	router := hatetepe.NewRouter()

	router.HandleFunc("GET", "/", func(w hatetepe.ResponseWriter, r *hatetepe.Request) {

		// w.Write([]byte("Hello, World!"))

		hello := HelloWorld{
			Message: "Hello, World!",
		}

		w.WriteJSON(hello)
	})

	router.HandleFunc("POST", "/hello", func(w hatetepe.ResponseWriter, r *hatetepe.Request) {
		log.Println("hello :: ", r.URL)

		var hello HelloWorld

		if err := r.ParseJSON(&hello); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error parsing JSON"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Your message: %s", hello.Message)))
		fmt.Println("Message:", hello.Message)
	})

	server := &hatetepe.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on :8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
