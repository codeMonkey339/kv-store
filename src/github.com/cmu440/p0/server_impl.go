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
	conns_num int
	/* the kv store to actually store data */
	kvstore KVStore
	/* currently active connections */
	conns []net.Conn
	/* channel put request will go to */
	dataChan chan string
}

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
	BUFFER_SIZE = 1024
	MAX_CONNS = 100
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
	kvServer.conns_num = 0
	/* []net.Conn is a different type from [2]net.Conn */
	kvServer.conns = make([]net.Conn, 0)
	kvServer.dataChan = make(chan string)
	return kvServer
}

// method on the keyValueServer to start the server with the given port #
/* able to use instance symbols with kvs */
func (kvs *keyValueServer) Start(port int) error {
	fmt.Println("Creating socket on port: ", port)
	l, _:= net.Listen(CONN_TYPE, CONN_HOST + ":" + strconv.Itoa(port))
	fmt.Println("Listening on " + CONN_HOST  + ":" + strconv.Itoa(port))
	go handleKvsAccess(kvs.dataChan, kvs)
	for { // Listen for an incoming connection
		conn, _ := l.Accept()
		kvs.conns = append(kvs.conns, conn)
		kvs.conns_num++
		fmt.Printf("Connected to socket %d\n", conn)
		go handleRequest(conn, kvs)
	}

	fmt.Printf("Closing the server\n")
	defer l.Close()
	return nil
}

// a goroutine that will handle put requests from clients
func handleKvsAccess(writeChan chan string, kvs *keyValueServer){
	fmt.Printf("Start handling put request\n")

	for{
		cmd := <- writeChan
		tokens := strings.Split(cmd, ",")
		fmt.Printf("Serving %v\n", cmd)
		switch strings.ToLower(strings.Trim(tokens[0], " ")){
		case "get": handleGetAccess(tokens, kvs)
		case "put": handlePutAccess(tokens, kvs)
		default:
			fmt.Printf("Invalid operation %v\n", cmd)
		}
		fmt.Printf("Finished serving %v\n", cmd)

	}
}

func handleGetAccess(tokens []string, kvs *keyValueServer){
	fmt.Printf("Processed get cmd %v %v\n", strings.Trim(tokens[0], " "),
		strings.Trim(tokens[1], " "))
	res := string(kvs.kvstore.get(strings.Trim(tokens[1], " ")))
	for i := 0; i < kvs.conns_num; i++ {
		conn := kvs.conns[i]
		fmt.Fprintf(conn, "%v,%v\n", strings.Trim(tokens[1], " "), res)
	}
}

func handlePutAccess(tokens []string, kvs *keyValueServer){
	fmt.Printf("Processed put cmd %v %v %v\n", tokens[0], tokens[1], tokens[2])
	kvs.kvstore.put(strings.Trim(tokens[1], " "),
		[]byte(strings.Trim(tokens[2], " ")))
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
func handleRequest(conn net.Conn, kvs *keyValueServer) {
	for{
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil{
			// break from the connection upon error
			fmt.Printf("encountered an error %v when reading from connection" +
				"\n", err)
			break;
		}
		fmt.Printf("Received command %v from connection %d\n", text[:len(text)- 1],
			conn)
		tokens:= strings.Split(text, ",")
		switch strings.ToLower(strings.TrimRight(tokens[0], " ")){
		case "get": doGet(text, conn, kvs)
		case "put": doPut(text, conn, kvs)
		default:
		}
	}
}

/*
work horse for the GET request

GET request format:
put, key, value

response format of GET:
key, value

*/
func doGet(cmd string, conn net.Conn, kvs *keyValueServer){
	fmt.Printf("Processing a get request %v\n", cmd)
	kvs.dataChan <- cmd[:len(cmd) - 1]
}

/*
work horse for the PUT request

PUT request format:
get, key

No response should be sent to any of the clients for a put request

*/
func doPut(cmd string, conn net.Conn, kvs *keyValueServer){
	fmt.Printf("Processing a put request: %v\n", cmd)
	kvs.dataChan <- cmd[:len(cmd) - 1]
}
