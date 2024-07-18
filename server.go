package hatetepe

import (
	"fmt"
	"net"
)

type Server struct {
	Addr    string
	Handler Handler
}

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	req, err := parseRequest(conn)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	resp := &response{
		conn:    conn,
		headers: make(map[string]string),
	}

	if s.Handler != nil {
		s.Handler.ServeHTTP(resp, req)
	} else {
		resp.WriteHeader(404)
		resp.Write([]byte("404 Not Found"))
	}
}
