package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

var (
	activeServers    = 2
	balanceCondition = 0
	hostBalancer     = "localhost:8000"
	hostsServers     = []string{"localhost:8100", "localhost:8200", "localhost:8300"}
)

//parallelize ..
func parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

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

func turnThirdServer() {
	if activeServers >= 3 {
		return
	}
	activeServers++
	cmdStr := "./../APIRest3/util/RunAPIRest.sh"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	log.Println("SE ACTIVO EL 3ER SERVIDOR")
}

func main() {

	serversUP := func() {

		ln, err := net.Listen("tcp", hostBalancer)
		if err != nil {
			log.Fatalf("Failed to connect to HostBalancer: %s", err)
		}

		log.Printf("listening on %s, balancing %s and %s like overload", hostBalancer, hostsServers[0], hostsServers[1])

		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("failed to accept: %s", err)
				continue
			}

			go handleConnection(conn, &hostsServers)
		}
	}

	serverResponseTime := func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				timestart := time.Now()
				resp, _ := http.Get("http://" + hostsServers[balanceCondition])
				log.Print(balanceCondition)
				if time.Since(timestart) > time.Duration(2)*time.Millisecond {
					log.Print("tiempo mayor a 2ms: ")
					go turnThirdServer()
				} else {
					log.Print("tiempo menor a 2ms: ")
				}
				log.Println(time.Since(timestart))
				resp.Body.Close()
			}
		}
	}
	parallelize(serversUP, serverResponseTime)
}
