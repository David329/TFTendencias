package main

import "time"
import "log"

func main() {
	timestart := time.Now()
	log.Println(time.Since(timestart) * time.Millisecond)
	log.Println(130 / 4)
}
