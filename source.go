package main

import (
	"net"
	"os"
)

//think about gzip!!!
func main() {
	args := os.Args
	if len(args) < 4 {
		panic("usage: client <origin:port> <filepath> <flag>, where flag is new or append and file path is absolute file path or relative, if it's located near the client")
	}
	conn, err := net.Dial("tcp", args[1])
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	sendFileToServer(args[2], args[3], conn)
}
