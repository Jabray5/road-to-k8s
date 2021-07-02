package main

import (
	"fmt"
	"time"
)

const SLEEPTIME int = 10
const LOGSTRING string = "A random string!"

func main() {

	// Loop indefinitely
	for {
		wait()
		t := time.Now()

		// There's probably a tidier way to do this!
		zone, _ := t.Zone()
		fmt.Printf("%v-%v-%v %v:%v:%v %v - %v\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), zone, LOGSTRING)
	}
}

func wait() {
	time.Sleep(time.Duration(SLEEPTIME) * time.Second)
}
