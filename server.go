package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"time"
)

func initialize_source(s string) net.Listener {
	if s == "gob" {
		ln, err := net.Listen("tcp", "127.0.0.1:8082")
		if err != nil {
			log.Fatal(err)
		}
		return ln
	}
	if s == "json" {
		ln, err := net.Listen("tcp", "127.0.0.1:8083")
		if err != nil {
			log.Fatal(err)
		}
		return ln
	}
	return nil
}

func gob_recieve(source net.Listener) {
	for {

		// accept incoming connections
		conn, err := source.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// pass connection into subproccess to handle incoming messages
		go func(conn net.Conn) {
			dec := gob.NewDecoder(conn)
			var shni []byte
			for {
				dec.Decode(&shni)
				fmt.Println(time.Now())
			}
		}(conn)

	}
}

func json_recieve(source net.Listener) {
	for {
		conn, err := source.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			dec := json.NewDecoder(conn)
			var shni []byte

				dec.Decode(&shni)
				fmt.Println(time.Now())
	
		}(conn)
	}
}


func main() {
	args := os.Args
	self := args[1]
	server := initialize_source(self)
	if self == "gob" {
		gob_recieve(server)
	}
	if self == "json" {
		json_recieve(server)
	}
	defer server.Close()
}
