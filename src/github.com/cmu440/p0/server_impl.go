// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"bufio"
)


type keyValueServer struct {
	/* the # of connections */
	conns int
	/* the kv store to actually store data */
	kvstore KVStore
}

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
	BUFFER_SIZE = 1024
)


/**
	create and initialize a new kv server

	the function New is a go convention for packages that create a core type
or different types for use by the application developer

	different between new (Thing) and &Thing{}:
	&Thing{} works only for Thing being a struct type, map type,
array type or slice type;
	new(Thing) works for Thing of any type
 */
func New() *keyValueServer{
	fmt.Println("Allocating a new keyValueServer")
	kvServer := new (keyValueServer)
	kvServer.kvstore.init_db()
	kvServer.conns = 0
	return kvServer
}

// method on the keyValueServer to start the server with the given port #
func (kvs *keyValueServer) Start(port int) error {
	fmt.Println("Creating socket on port: ", port)
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + strconv.Itoa(port))
	if err != nil {
		/* log network error to stdio*/
		fmt.Println("Error listening:", err.Error())
		return err
	}
	fmt.Println("Listening on " + CONN_HOST  + ":" + strconv.Itoa(port))
	for { // Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			/* Go has the continue keyword */
			continue
		}
		fmt.Printf("Connected to socket %d\n", conn)
		/* Start is a class method, now it calls a non-class method.
		How does it work? */
		go handleRequest(conn)
	}

	/* when existing from this function, close the port */
	defer l.Close()
	return nil
}


func (kvs *keyValueServer) Close() {
	// TODO: release the kv store; how to terminate all goroutines?
}

func (kvs *keyValueServer) Count() int {
	// TODO: return the # of connected clients
	return -1
}

/*
The server should not assume that the key-value API function listed are
thread-safe. You will be responsible for ensuring that there are no race
conditions while accessing the database? How to guard critical region with
goroutine? If this is known,then problem solved

All synchronization must be done using goroutines,
channels and Go's channel-based select statement The server must implement a
Count() function that returns the # of connected clients
*/
func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	text, _ := reader.ReadString('\n')
	fmt.Printf("Received command %v from connection %d\n", text, conn)
	tokens:= strings.Split(text, ",")
	switch strings.ToLower(strings.TrimRight(tokens[0], " ")){
	case "get": doGet(tokens, conn)
	case "put": doPut(tokens, conn)
	default:
	}
}

/*
work horse for the GET request

GET request format:
put, key, value

response format of GET:
key, value

*/
func doGet(tokens []string, conn net.Conn){
	//todo: complete the logic here
	fmt.Printf("Processing a get request %v, %v\n", tokens[0],tokens[1])
	fmt.Fprintf(conn, "Received command %v, %v\n", tokens[0], tokens[1])
}

/*
work horse for the PUT request

PUT request format:
get, key

No reponse should be sent to any of the clients for a put request

*/
func doPut(tokens []string, conn net.Conn){
	//todo: complete the logic here
	fmt.Printf("Processing a put request %v, %v, %v\n", tokens[0],
		tokens[1], tokens[2])
}
