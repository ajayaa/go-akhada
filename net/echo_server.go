package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic("port bound already")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("some error!")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()
	fmt.Fprintf(os.Stdout, "conn opened from: %s\n", remoteAddr)
	var buffer = make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stdout, "conn closed from: %s\n", remoteAddr)
			return
		}
		message := string(buffer[:n])
		fmt.Fprintf(os.Stdout, "%s%s%s", remoteAddr, ": ", message)
		conn.Write([]byte(message))
	}
}
