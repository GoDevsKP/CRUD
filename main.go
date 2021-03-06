package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
	crud "./crud"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8888"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	dataStorage := crud.DB{}
	dataStorage.Init()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		fmt.Println("New client")

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn, &dataStorage)
	}
}

func handleConnection(conn net.Conn, dataStorage *crud.DB) {
	for {
		connbuf := bufio.NewReader(conn)
		for {
			result, err := connbuf.ReadString('\n')
			if err != nil {
				conn.Write([]byte("Error\n"))
			}
			go handleRequest(result, conn, dataStorage)
		}
	}
}

func handleRequest(input string, conn net.Conn, dataStorage *crud.DB) {
	query := strings.Fields(input)

	if (!crud.Check(query)) {
		conn.Write([]byte("Bad query\n"))
	} else {
		conn.Write([]byte(crud.Handle(query, dataStorage) + "\n" ))
	}
}
