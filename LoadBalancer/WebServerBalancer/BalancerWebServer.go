//package main allow balance 3 active servers
package main

import (
	"io"
	"log"
	"net"
)

var (
	activeServers    = 2
	balanceCondition = 0
	hostsServers     = []string{"localhost:9100", "localhost:9200"}
)

const (
	hostBalancer = "localhost:9000"
)

//copy Buffer like mirror
func copy(w io.WriteCloser, r io.Reader) {
	defer w.Close()
	io.Copy(w, r)
}

//handleConnection this method select the next server(+-condition) to send the data
func handleConnection(us *net.Conn, servers *[]string) {

	//if balanceCondition+1 equals activeServers, then reset the balanceCondition, else balanceCondition++
	if (balanceCondition + 1) == activeServers {
		balanceCondition = 0
	} else {
		balanceCondition++
	}

	//Get path of the next server to send the data
	ds, err := net.Dial("tcp", (*servers)[balanceCondition])
	if err != nil {
		(*us).Close()
		log.Printf("failed to dial %s: %s", (*servers)[balanceCondition], err)
		return
	}

	//Send data to next server, then response data, like mirror
	go copy(ds, (*us))
	go copy((*us), ds)
}

func main() {

	//Listen all entry points
	ln, err := net.Listen("tcp", hostBalancer)
	if err != nil {
		log.Fatalf("Failed to connect to HostBalancer: %s", err)
	}
	log.Printf("listening on %s, balancing %s and %s", hostBalancer, hostsServers[0], hostsServers[1])

	//Accept and Handle new connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}

		go handleConnection(&conn, &hostsServers)
	}
}
