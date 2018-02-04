package main

import (
    "fmt"
	"net"
	"strconv"
	"os"
	"bufio"
	"strings"
)

const (
    defaultHost = "localhost"
    defaultPort = 9999
    defaultProtocol = "tcp"
)

// To test your server implementation, you might find it helpful to implement a
// simple 'client runner' program. The program could be very simple, as long as
// it is able to connect with and send messages to your server and is able to
// read and print out the server's response to standard output. Whether or
// not you add any code to this file will not affect your grade.
func main() {
    fmt.Printf("Creating a test Client\n")
    conn, _:= net.Dial(defaultProtocol, defaultHost + ":" + strconv.Itoa(
    	defaultPort))

    for{
    	reader := bufio.NewReader(os.Stdin)
    	fmt.Printf("Text to send: ")
    	text, _ := reader.ReadString('\n')
    	//todo: ideally read is non-blocking
    	fmt.Printf("Sending to server\n")
		fmt.Fprintf(conn, text + "\n")

		if strings.Contains(text,"get") {
			// only log for get request
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Printf("Message received from server: %v\n", message)
		}
	}
}

