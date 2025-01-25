package types

import (
	"fmt"
)

type Response struct {
	StatusCode int
	Headers    Header
	Body       string
	Protocol   string
}

type Header struct {
	Fields map[string]string
}

func statusMsg(statusCode int) string {
	var msg string
	switch statusCode {
	case 200:
		msg = "Ok"
	case 201:
		msg = "Created"
	case 404:
		msg = "Not Found"
	case 500:
		msg = "Server Error"
	}

	return msg

}

func NewResponse(statusCode int, header Header, body string) string {

	// get the status message
	statusMsg := statusMsg(statusCode)

	// looksm like HTTP/1.1 200 Ok
	statusLine := fmt.Sprintf("%s %d %s\r\n", "HTTP/1.1", statusCode, statusMsg)

	headerStr := ""
	if header.Fields != nil {
		for key, value := range header.Fields {
			headerStr += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}

	// Always add Content-Length header
	headerStr += fmt.Sprintf("Content-Length: %d\r\n", len(body))
	headerStr += fmt.Sprintf("Server: %s\r\n", "go-httpx")

	// End of headers
	headerStr += "\r\n"

	// add the body too ...
	return statusLine + headerStr + body
}

func NewTextHeader() Header {
	textHeader := Header{
		Fields: map[string]string{
			"Content-Encoding": "text/plain",
		},
	}
	return textHeader
}

func AddHeader(key, value string, header *Header) *Header{
	header.Fields[key] = value
	return header
}
