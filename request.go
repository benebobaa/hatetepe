package hatetepe

import (
	"bufio"
	"encoding/json"
	// "errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
}

// TODO: Need handle error if the body json is nil
func (r *Request) ParseJSON(v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func parseRequest(conn net.Conn) (*Request, error) {
	reader := bufio.NewReader(conn)

	// Read the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	parts := strings.Fields(requestLine)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line")
	}

	req := &Request{
		Method:  parts[0],
		URL:     parts[1],
		Headers: make(map[string]string),
	}

	// Read headers
	for {
		line, err := reader.ReadString('\n')

		if err != nil || line == "\r\n" {
			break
		}

		parts := strings.SplitN(strings.TrimSpace(line), ":", 2)

		if len(parts) == 2 {
			req.Headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	req.Body = reader

	return req, nil
}
