package main

import (
	"fmt"
	"time"

	"github.com/Jabray5/road-to-k8s/api"
)

const SLEEPTIME int = 10
const LOGSTRING string = "A random string!"

var enableLogging bool = true

func main() {

	api.StartListening()

	// Loop indefinitely
	for {
		if enableLogging {
			logTime()
			wait()
		}

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
