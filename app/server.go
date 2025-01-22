package main

import (
	"fmt"
	"net"
	"os"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/workerpool"
)

func main() {

	// create a tcp-listener
	lis, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}


	// create a worker-pool with max 2 workers
	pool := workerpool.NewPool(2)
	defer pool.Stop()

	// main server loop ... accept connections by keeping the server on
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		pool.Submit(conn)
	}
}
