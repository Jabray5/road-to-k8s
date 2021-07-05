package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
	Date string `json:"date"`
}

type StoryPost struct {
	Story string `json:"message"`
}

type StoryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage endpoint hit")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	timeResponse := TimeResponse{Time: t.Format("15:04:05"), Date: t.Format("2006-01-02")}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timeResponse)
}

func story(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var storyPost StoryPost

	err := json.Unmarshal(reqBody, &storyPost)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Println(storyPost)
		var storyResponse = StoryResponse{Success: true, Message: storyPost.Story}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(storyResponse)
	}
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
