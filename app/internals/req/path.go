package req

import (
	"strings"
)

// returns the route / url path from the request string

func GetPath(fullString string) (reqPath string) {

	// following method splits all words with spaces around into an array
	// GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n
	// in above the array[1] will return us /index or the path
	space_separated_params := strings.Split(fullString, " ")
	return space_separated_params[1]

}

// following func works like this ...
// if full route is /echo/abc/hello and num of Paths to be ignored is 2
// then returns /hello as /echo and /abc are ignored
func NestedPath(fullRoute string, numOfPathsIgnored int) (nestedRoute string) {
	slash_separated_route := strings.Split(fullRoute, "/")

	var parent_route_ignored string

	// we range from first index [1:] as there is a leading slash too
	// like /echo, so during splitting we have to ignore the 0-th index anyway
	for index, nestedItem := range slash_separated_route[1:] {
		if index < numOfPathsIgnored {
			continue
		} else {
			parent_route_ignored = parent_route_ignored + "/" + nestedItem
		}
	}
	return parent_route_ignored
}
