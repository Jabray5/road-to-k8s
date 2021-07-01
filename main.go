package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const SLEEPTIME int = 10
const LOGSTRING string = "A random string of your choice."

func main() {

	// Get the log file, or create one if it doesn't exist
	file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// Loop indefinitely
	for {
		wait()
		log.Println(LOGSTRING)
		fmt.Println("Log added.")
	}
}

func wait() {
	time.Sleep(time.Duration(SLEEPTIME) * time.Second)
}
