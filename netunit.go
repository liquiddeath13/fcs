package main

import (
	"net"
	"os"
)

func sendFileTransferHeader(fileInfo os.FileInfo, mode string, connection net.Conn) {
	sendContentLength(fileInfo.Size(), connection)
	sendString(fileInfo.Name(), connection)
	sendString(mode, connection)
}

func sendContentLength(contentLength int64, connection net.Conn) {
	sendString(intToStr(contentLength), connection)
}

func sendContent(content []byte, connection net.Conn) {
	connection.Write(content)
}

func sendString(str string, connection net.Conn) {
	str = fillString(str, 64)
	sendContent([]byte(str), connection)
}

func sendOpType(opType string, connection net.Conn) {
	//sendString(opType, connection)
}
