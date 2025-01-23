package req

import (
	"strings"
)

func Body(reqStr string) string{

	// separate the parameters by \r\n 
	separated_params := strings.Split(reqStr, "\r\n")
	
	// grab the \r\n at the last of the params array 
	body := separated_params[len(separated_params)-1]

	return body

}