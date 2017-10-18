package main

import (
	"io"
	"log"
	"net"
)

var (
	activeServers    = 2
	balanceCondition = 0
	hostBalancer     = "localhost:9000"
	hostsServers     = []string{"localhost:9100", "localhost:9200"}
)

func copy(w io.WriteCloser, r io.Reader) {
	defer w.Close()
	io.Copy(w, r)
}

func handleConnection(us net.Conn, servers *[]string) {

	//Para q sea el primero del arreglo balancecondition=0
	if (balanceCondition + 1) == activeServers {
		balanceCondition = 0
	} else {
		balanceCondition++
	}

	ds, err := net.Dial("tcp", (*servers)[balanceCondition])

	if err != nil {
		us.Close()
		log.Printf("failed to dial %s: %s", (*servers)[balanceCondition], err)
		return
	}

	go copy(ds, us)
	go copy(us, ds)
}

func main() {

	ln, err := net.Listen("tcp", hostBalancer)
	if err != nil {
		log.Fatalf("Failed to connect to HostBalancer: %s", err)
	}

	log.Printf("listening on %s, balancing %s and %s", hostBalancer, hostsServers[0], hostsServers[1])

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}

		go handleConnection(conn, &hostsServers)
	}

}
