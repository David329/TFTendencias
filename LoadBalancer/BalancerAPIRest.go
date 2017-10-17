package main

import (
	"io"
	"log"
	"net"
)

var balanceCondition = true

func copy(w io.WriteCloser, r io.Reader) {
	defer w.Close()
	io.Copy(w, r)
}

func handleConnection(us net.Conn, server1 *string, server2 *string) {

	var selectedServer string

	if balanceCondition {
		selectedServer = *server1
	} else {
		selectedServer = *server2
	}

	balanceCondition = !balanceCondition

	ds, err := net.Dial("tcp", selectedServer)

	if err != nil {
		us.Close()
		log.Printf("failed to dial %s: %s", selectedServer, err)
		return
	}

	go copy(ds, us)
	go copy(us, ds)
}

func main() {

	//0: HostBalancer. 1, 2: HostServers
	var hosts = []string{"localhost:8000", "localhost:8100", "localhost:8200"}

	ln, err := net.Listen("tcp", hosts[0])
	if err != nil {
		log.Fatalf("Failed to connect to HostBalancer: %s", err)
	}

	log.Printf("listening on %s, balancing %s, %s", hosts[0], hosts[1], hosts[2])

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}

		go handleConnection(conn, &hosts[1], &hosts[2])
	}
}
