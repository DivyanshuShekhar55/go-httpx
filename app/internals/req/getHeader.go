package req

import (
	"strings"
)

type Header struct {
	Fields map[string]string
}

// the cut method returns before, after and isFound values for the occurence of a separator 
// we can also get the body and status line from here
func Headers(reqStr string) *Header {
	_, header_body, _ := strings.Cut(reqStr, "\r\n")

	header, _, _ := strings.Cut(header_body, "\r\n\r\n")

	headerMap := &Header{
		Fields: make(map[string]string),
	}

	for _, field := range strings.Split(header, "\r\n") {
		key, value, _ := strings.Cut(field, ": ")
		headerMap.Fields[key] = value
	}

	return headerMap
}
