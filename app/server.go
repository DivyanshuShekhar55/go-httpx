package main

import (
	"fmt"
	"net"
	"os"
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

	// write a message to this connection
	// any simple msg string will not work because the testing tools like curl or postman will read http response only, so we need to send a HTTP compliant response
	msg := "HTTP/1.1 200 OK\r\n\r\n"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("error writing over connection")
	}

}
