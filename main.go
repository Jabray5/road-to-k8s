package main

import (
	"fmt"
	"time"

	"./api"
)

const SLEEPTIME int = 10
const LOGSTRING string = "A random string!"

var enableLogging bool = true

func main() {

	// Loop indefinitely
	for {
		if enableLogging {
			logTime()
			wait()
		}
		api.StartListening()

		// if newMessage {
		// 	// log the new message
		// }

	}
}

func wait() {
	time.Sleep(time.Duration(SLEEPTIME) * time.Second)
}

func logTime() {
	t := time.Now()
	// There's probably a tidier way to do this!
	zone, _ := t.Zone()
	fmt.Printf("%v-%v-%v %v:%v:%v %v - %v\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), zone, LOGSTRING)
}
