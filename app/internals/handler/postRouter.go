package handler

import (
	"strings"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/filehandler"
	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/req"
)

func Post(route string, reqStr string) (msg string) {
	switch {

	// use: Invoke-WebRequest -Uri "http://localhost:4221/files/abc" -Method Post -ContentType "application/octet-stream" -Body 'hello devs'
	case strings.HasPrefix(route, "/files"):

		// extract /abc from /file/abc
		file_name := req.NestedPath(route, 1)

		// next extract body...
		body := req.Body(reqStr)

		// create a file and put the data there ...
		if err := filehandler.CreateFile(file_name, body); err != nil {
			// I am not actually sure about this status code !
			msg = "HTTP/1.1 424 Couldn't Perform Operation\r\n\r\n"
			return
		}
		msg = "HTTP/1.1 201 Created\r\n\r\n"

	}

	return msg
}
