package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8085")

	// accept connection on port
	conn, _ := ln.Accept()

	fmt.Println("Server launched!")

	go messageListener(conn)
	messegeSender(conn)
}

func messageListener(conn net.Conn) {

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// erase the current line and output message received underneath
		fmt.Print("\033[2K", "\nMessage Received:", string(message))

	}
}

func messegeSender(conn net.Conn) {

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		if strings.Compare(text, "\n") != 0 {
			// send to socket
			conn.Write([]byte(text + "\n"))
		}

	}
}
