package main

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/path"
)

func main() {

	// create a tcp-listener
	lis, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	// accept connections by keeping the server on...will shut once a connection is over
	conn, err := lis.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	// close the connection when not needed
	defer conn.Close()

	// Next, lets get the url requested, a GET looks like this :GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n
	// send 404 for any other path than the home route
	buffer := make([]byte, 10240)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error while reading the request buffer")
	}


	// buffer is a long sequence of bytes (like 12 17 ... 0 0 ...0)
	// convert it to readable string, using the bytes package
	buf := bytes.NewBuffer(buffer)
	//fmt.Println(buf.String()) returns full req string


	route := path.GetPath(buf.String())
	msg := Router(route)

	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("error writing over connection")
	}

}
