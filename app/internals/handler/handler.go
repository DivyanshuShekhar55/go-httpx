package handler

import (
	"bytes"
	"fmt"
	"net"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/req"
)

func HandleReq(conn net.Conn) {

	// Next, lets get the url requested, a GET looks like this :GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n

	buffer := make([]byte, 1024)
	defer conn.Close()

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error while reading the request buffer")
	}

	// buffer is a long sequence of bytes (like 12 17 ... 0 0 ...0)
	// convert it to readable string, using the bytes package
	buf := bytes.NewBuffer(buffer)
	fullString := buf.String() //returns full req string
	// fmt.Println(fullString)

	route := req.GetPath(buf.String())
	method := req.Method(fullString)

	var res string // response message 

	switch method {
	case "GET":
		res = Get(route, fullString)
	case "POST" :
		res = Post(route, fullString)
	}

	_, err = conn.Write([]byte(res))
	if err != nil {
		fmt.Println("error writing over connection")
	}

}
