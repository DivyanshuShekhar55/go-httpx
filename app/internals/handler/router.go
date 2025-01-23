package handler 

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/path"
)

func Router(route string, fullString string) (msg string) {

	switch {
	case route == "/":
		msg = "HTTP/1.1 200 OK\r\n\r\n"

	case strings.HasPrefix(route, "/echo"):
		content := path.NestedPath(route, 1)
		content_len := strconv.Itoa(len(content))
		msg = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", content_len, content)

	case route == "/user-agent":

		// use the command in curl : Invoke-WebRequest -Uri http://localhost:4221/user-agent -Headers @{"User-Agent" = "foobar/1.2.3"}
		content := path.GetUserAgent(fullString)
		content_len := strconv.Itoa(len(content))

		msg = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", content_len, content)

	case strings.HasPrefix(route, "/file"):
		

	default:
		msg = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return msg
}
