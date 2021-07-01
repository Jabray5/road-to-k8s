package main

import (
	"fmt"
	"time"
)

const SLEEPTIME int = 3
const LOGSTRING string = "A random string!"

func main() {

	// FOR CREATING A LOG.TXT FILE
	// Get the log file, or create one if it doesn't exist
	// file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)

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
