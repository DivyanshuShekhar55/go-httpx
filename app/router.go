package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/path"
)

func Router(route string) (msg string) {

	switch {
	case route == "/":
		msg = "HTTP/1.1 200 OK\r\n\r\n"

	case strings.HasPrefix(route, "/echo"):
		content := path.NestedPath(route, 1)
		content_len := strconv.Itoa(len(content))
		msg = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", content_len, content)

	default:
		msg = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return msg
}
