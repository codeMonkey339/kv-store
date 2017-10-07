// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import {
	"fmt"
	"net"
	"bufio"
	"bytes"
	"strconv"
}

type keyValueServer struct {
	var kvstore KVStore
	var conns int
}

const (
	CONN_HOST = "localhost"
	CONN_TYPE = "tcp"
	BUFFER_SIZE = 1024
)



// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer *keyValueServer{
	kvServer := &{conns: 0, kvstore: make(KVStore)}
	kvServer.kvstore.init_db()
	return kvServer
}

func (kvs *keyValueServer) Start(port int) error {
	l, err = net.Listen(CONN_TYPE, CONN_HOST + ":" + string(port))
	if err != nil {
		/* log network error to stdio*/
		fmt.Println("Error listening:", err.Error())
		return err
	}
	/* when existing from this function, close the port */
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST  + ":" + string(port))
	for { // Listen for an incoming connection
		conn, err = l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			/* does this keyword exist ?*/
			continue
		}
		// Handle the new connection in a new goroutine
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

// TODO: add additional methods/functions below!

/* this function will handle newly accepted connections in a go */
/* routine */
func handleRequest(conn net.Conn) {
	buf := make([]byte, BUFFER_SIZE)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	/* todo: need to parse request and send back response */
}
