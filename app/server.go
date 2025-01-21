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

	HandleReq(conn)

}
