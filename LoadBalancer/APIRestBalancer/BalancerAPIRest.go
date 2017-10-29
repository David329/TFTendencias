//package main allow balance 2 active servers and 1 sleep server
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
	hostsServers     = []string{"localhost:8100", "localhost:8200", "localhost:8300"}
)

const (
	hostBalancer = "localhost:8000"
)

//parallelize functions with goRoutines
func parallelize(functions ...func()) {

	//New WaitGroup for handle goRoutines
	var waitGroup sync.WaitGroup

	//Add length of goRoutines to run
	waitGroup.Add(len(functions))

	//End of Functions wait for goRoutines to finish=(functionsGoRoutines.length==0)
	defer waitGroup.Wait()

	//Foreach of functions to execute in goRoutines
	for _, function := range functions {
		go func(copy func()) {

			//functionsGoRoutines.length -1
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

//copy Buffer like mirror
func copy(w io.WriteCloser, r io.Reader) {
	defer w.Close()
	io.Copy(w, r)
}

//handleConnection this method select the next server(+-condition) to send the data
func handleConnection(us net.Conn, servers *[]string) {

	//if balanceCondition+1 equals activeServers, then reset the balanceCondition, else balanceCondition++
	if (balanceCondition + 1) == activeServers {
		balanceCondition = 0
	} else {
		balanceCondition++
	}

	//Get path of the next server to send the data
	ds, err := net.Dial("tcp", (*servers)[balanceCondition])
	if err != nil {
		us.Close()
		log.Printf("failed to dial %s: %s", (*servers)[balanceCondition], err)
		return
	}

	//Send data to next server, then response data, like mirror
	go copy(ds, us)
	go copy(us, ds)
}

//turnThirdServer this method allow run the third server if the two active are very overLoaded
func turnThirdServer() {

	//if all servers are up return
	if activeServers >= 3 {
		return
	}

	activeServers++

	//run sh command, that run bin file of Server3 compiled
	cmdStr := "./../../APIRest3/util/RunAPIRest.sh"
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	log.Println("Third Server Up!!!")
}

func main() {

	//Start HostServer Balancing two Servers and one sleep
	serversUP := func() {

		//Listen all entry points
		ln, err := net.Listen("tcp", hostBalancer)
		if err != nil {
			log.Fatalf("Failed to connect to HostBalancer: %s", err)
		}
		log.Printf("listening on %s, balancing %s and %s .%s like overload", hostBalancer, hostsServers[0], hostsServers[1], hostsServers[2])

		//Accept and Handle new connection
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("failed to accept: %s", err)
				continue
			}

			go handleConnection(conn, &hostsServers)
		}
	}

	//Every 3 seconds do ping to next server to measure latency, if it is more than 2ms then wake up third server
	serverResponseTime := func() {

		//Time 3 seconds
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C: //channel to handle 3 seconds

				//get time noew
				timestart := time.Now()

				//ping to next server
				resp, _ := http.Get("http://" + hostsServers[balanceCondition])

				//print the server to evaluate
				log.Print(balanceCondition)

				//if time is more than 2ms, wake up third server, else do nothing
				if time.Since(timestart) > time.Duration(2)*time.Millisecond {
					log.Print("Time more than 2ms: ")
					go turnThirdServer()
				} else {
					log.Print("Time less than 2ms: ")
				}

				//print ping time
				log.Println(time.Since(timestart))
				resp.Body.Close()
			}
		}
	}

	//parallelize all functions
	parallelize(serversUP, serverResponseTime)
}
