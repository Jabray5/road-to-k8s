package api

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
	Date string `json:"date"`
	Zone string `json:"zone"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage endpoint hit")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	timeResponse := TimeResponse{Time: t.Format("15:04:05"), Date: t.Format("2006-01-02"), Zone: t.Format("MST")}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeResponse)
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/story", storyEndpoint)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func StartListening() {
	handleRequests()
}
