package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8085")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		go listen(conn)
	}
}

func listen(conn net.Conn) {

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("\033[2K Message from server: " + message)
	// \033[2K stands for erase line in ANSI

}
