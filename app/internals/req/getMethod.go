package req

import "strings"

func Method(reqStr string)string{
	
	// separate the parameters by space
	// reqStr is like: POST /files/number HTTP/1.1\r\nHost: ...
	separated_params := strings.Split(reqStr, " ")
	
	// get the first item ...
	method := separated_params[0]
	return method
}