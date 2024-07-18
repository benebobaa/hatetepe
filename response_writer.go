package hatetepe

import (
	"encoding/json"
	"fmt"
	"net"
)

type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
	Header() map[string]string
	SetHeader(key, value string)
	WriteJSON(v interface{}) error
}

func (r *response) SetHeader(key, value string) {
	if r.wroteHeader {
		return // Can't set headers after they've been written
	}
	r.headers[key] = value
}

type response struct {
	conn        net.Conn
	headers     map[string]string
	wroteHeader bool
}

func (r *response) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(200)
	}
	return r.conn.Write(b)
}

func (r *response) WriteHeader(statusCode int) {
	if r.wroteHeader {
		return
	}
	r.wroteHeader = true
	fmt.Fprintf(r.conn, "HTTP/1.1 %d %s\r\n", statusCode, statusText(statusCode))
	for k, v := range r.headers {
		fmt.Fprintf(r.conn, "%s: %s\r\n", k, v)
	}
	fmt.Fprintf(r.conn, "\r\n")
}

func (r *response) Header() map[string]string {
	return r.headers
}

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	default:
		return "Unknown"
	}
}

func (r *response) WriteJSON(v interface{}) error {
	r.SetHeader("Content-Type", "application/json")
	if !r.wroteHeader {
		r.WriteHeader(200)
	}
	return json.NewEncoder(r.conn).Encode(v)
}
