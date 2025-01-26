package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/compression"
	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/filehandler"
	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/req"
	"github.com/DivyanshuShekhar55/go-htttpx/app/types"
)

func Get(route string, fullString string) (msg string) {

	// create an empty header, this common header will be modified for each route
	resHeader := types.Header{
		Fields: map[string]string{},
	}

	var resBody string

	reqHeader := req.Headers(fullString)
	encoding, ok := reqHeader.Fields["Accept-Encoding"]
	if ok && encoding == "gzip" {
		resHeader = *types.AddHeader("Content-Encoding", "gzip", &resHeader)
	}

	switch {
	case route == "/":
		resBody = compression.AddGZip("ok") 
		msg = types.NewResponse(200, types.NewTextHeader(), resBody)

	case strings.HasPrefix(route, "/echo"):

		content := req.NestedPath(route, 1)
		content_len := strconv.Itoa(len(content))

		resHeader = *types.AddHeader("Content-Type", "text/plain", &resHeader)
		resHeader = *types.AddHeader("Content-Length", string(content_len), &resHeader)

		resBody = compression.AddGZip(content)
		msg = types.NewResponse(200, resHeader, resBody)

	case route == "/user-agent":

		// use the command in curl : Invoke-WebRequest -Uri http://localhost:4221/user-agent -Headers @{"User-Agent" = "foobar/1.2.3"}
		content := req.GetUserAgent(fullString)
		content_len := strconv.Itoa(len(content))

		msg = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %s\r\n\r\n%s", content_len, content)

	case strings.HasPrefix(route, "/file"):
		file_name := req.NestedPath(route, 1)
		fmt.Println(file_name)
		content, err := filehandler.GetFile(file_name)
		if err != nil {
			msg = "HTTP/1.1 404 Not Found\r\n\r\n"
			return
		}
		content_len := strconv.Itoa(len(content))
		msg = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %s\r\n\r\n%s", content_len, content)

	default:
		msg = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return msg
}
