package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
	Date string `json:"date"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage endpoint hit")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	timeResponse := TimeResponse{Time: t.Format("15:04:05"), Date: t.Format("2006-01-02")}
	byteArray, err := json.Marshal(timeResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(byteArray))
}

func story(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "story endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/story", story)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func StartListening() {
	handleRequests()
}
