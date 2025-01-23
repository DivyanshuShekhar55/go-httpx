package req

import (
	"strings"
)

func GetUserAgent(fullString string) (userAgent string) {
	space_separated_params := strings.Split(fullString, "\r\n")

	// the above var is like : User-Agent: foobar
	// again we need to split 
	userAgent = strings.Split(space_separated_params[1], " ")[1]
	return userAgent
}