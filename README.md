# Hatetepe

Hatetepe is a minimalist and efficient HTTP server for Go, inspired by the `net/http` package. It aims to provide a simple, yet powerful, foundation for building HTTP services. This project is primarily for experimental learning purposes.

## Features

* **Simple API:** Intuitive and easy-to-use API inspired by `net/http`.
* **Lightweight:** Minimal dependencies and overhead.
* **Compliant:** Follows HTTP/1.1 standards.

## Installation

To install Hatetepe, use `go get`:

```sh
go get -u github.com/benebobaa/hatetepe
```

## Quick Start

Here's a quick example to get you started:

```go
package main

import (
	"fmt"

	http "github.com/benebobaa/hatetepe"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	// Create Router
	router := http.NewRouter()

	// Add Routes GET
	router.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteJSON(map[string]string{"message": "Hello, World!"})
	})

	// Add Routes POST
	router.HandleFunc("POST", "/hello", func(w http.ResponseWriter, r *http.Request) {

		var hello HelloWorld

		if err := r.ParseJSON(&hello); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Error parsing JSON"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Your message: %s", hello.Message)))
		fmt.Println("Message:", hello.Message)
	})
	
	// Create Server
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Start Server
	fmt.Println("Server is running on port 8081")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
```

## Usage

### Creating a Server

To create a new server instance:

```go
server := hatetepe.NewServer()
```

### Handling Requests

You can handle requests by defining handler functions and registering them with the server:

```go
// Define Handler Function
helloFunc := func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

// Register Handler Function
handler := http.HandlerFunc(helloFunc)

server := http.Server{
    Addr:    ":8081", 
	// Register Handler
    Handler: handler,
}

```

### Starting the Server

To start the server on a specific address:

```go
// Create Server
server := http.Server{
    Addr:    ":8081",
    Handler: router,
}

// Start Server
fmt.Println("Server is running on port 8081")
err := server.ListenAndServe()
```

### Router

Hatetepe supports basic router for request processing. Router can be added as follows:

```go
// Create Router
router := http.NewRouter()

// Add Routes GET
router.HandleFunc("GET", "/hello", func(w http.ResponseWriter, r *http.Request) {
    w.WriteJSON(map[string]string{"message": "Hello, World!"})
})
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

Hatetepe is licensed under the MIT License. See [LICENSE](LICENSE) for more information.

## Acknowledgements

Inspired by the excellent `net/http` package in the Go standard library.

## Note

This project is for experimental learning purposes and may not be suitable for production use.

Would you like me to explain or break down any part of this README?