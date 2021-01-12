package main

import (
	"io"
	"log"
	"net"
	"os"
)

func sendFileToServer(filePath string, mode string, connection net.Conn) {
	sendOpType("ft_start", connection)
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
		sendOpType("ft_end", connection)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err.Error())
		sendOpType("ft_end", connection)
		return
	}
	sendFileTransferHeader(fileInfo, mode, connection)
	sendBuffer := make([]byte, 4096)
	log.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			sendOpType("ft_end", connection)
			break
		}
		sendContent(sendBuffer, connection)
	}
	log.Println("Closing connection!")
}
