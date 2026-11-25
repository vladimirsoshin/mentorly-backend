package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	println("Client connected:", conn.RemoteAddr())
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	response := "HTTP/1.1 404 Not Found\r\nContent-Length: 0\r\n\r\n"
	if n > 0 {
		message := string(buf[:n])
		//fmt.Println("Received message:", message)
		parts := strings.Fields(message)
		if len(parts) > 1 {
			path := parts[1]

			if path == "/health" {
				response = "HTTP/1.1 200 OK\r\nContent-Length: 2\r\n\r\nOK"
			}

		}
	}

	//response := "HTTP/1.1 200 OK\r\nContent-Length: 2\r\n\r\nOK"
	conn.Write([]byte(response))

	conn.Close()
}
