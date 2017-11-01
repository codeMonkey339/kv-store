// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)


type keyValueServer struct {
	conns int
	kvstore KVStore
}

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
	BUFFER_SIZE = 1024
)



// New creates and returns (but does not start) a new KeyValueServer.
func New() *keyValueServer{
	kvServer := &keyValueServer{
		conns: 0,
	}
	kvServer.kvstore.init_db()
	return kvServer
}

func (kvs *keyValueServer) Start(port int) error {
	fmt.Println("Creating socket on port: ", port)
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + strconv.Itoa(port))
	if err != nil {
		/* log network error to stdio*/
		fmt.Println("Error listening:", err.Error())
		return err
	}
	/* when existing from this function, close the port */
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST  + ":" + strconv.Itoa(port))
	for { // Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			/* Go has the continue keyword */
			continue
		}
		go handleRequest(conn)
	}
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
The server should not assume that the key-value API function listed are thread-safe.
You will be responsible for ensuring that there are no race conditions while accessing the database?

How to guard critical region with goroutine? If this knows, then problem solved

All synchronization must be done using goroutines, channels and Go's channel-based select statement

The server must implement a Count() function that returns the # of connected clients
*/
func handleRequest(conn net.Conn) {
	buf := make([]byte, BUFFER_SIZE)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println(conn)
	fmt.Println("Reading len: ", reqLen)
	/* todo: need to parse request and send back response */
	tokens:= strings.Split(string(buf[:]), ",")
	switch strings.ToLower(strings.TrimRight(tokens[0], " ")){
	case "get": doGet()
	case "put": doPut()
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
func doGet(){

}

/*
work horse for the PUT request

PUT request format:
get, key

No reponse should be sent to any of the clients for a put request

*/
func doPut(){

}
