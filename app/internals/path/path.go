package path

import (
	"strings"
)

// returns the route / url path from the request string

func GetPath(fullPath string) (reqPath string) {

	// following method splits all words with spaces around into an array
	// GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n
	// in above the array[1] will return us /index or the path
	space_separated_params := strings.Split(fullPath, " ")
	return space_separated_params[1]

}
