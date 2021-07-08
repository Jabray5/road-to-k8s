package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const FILE_LOCATION = "stories.json"

type StoryPost struct {
	Story string `json:"message"`
	UID   string `json:"uid"`
}

type StoryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UID     string `json:"uid"`
}

func storyEndpoint(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var storyPost StoryPost
	storyPost.UID = makeUID()

	err := json.Unmarshal(reqBody, &storyPost)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Println("story:", storyPost.Story, "-", "uid:", storyPost.UID)
		var storyResponse = StoryResponse{Success: true, Message: storyPost.Story, UID: storyPost.UID}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(storyResponse)

		// add to local text file
		fmt.Println("Writing story to file...")
		writeStoryToFile(storyPost)
	}
}

func writeStoryToFile(s StoryPost) {
	bytes, err := json.Marshal(s)

	if err != nil {
		fmt.Println(err)
	}

	// Currently not writing proper json
	// Will probably need to unmarshal the entire thing, append new json, then write back to the file
	jsonFile, err := os.OpenFile(FILE_LOCATION, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(err)
	}
	defer jsonFile.Close()

	if _, err := jsonFile.Write(bytes); err != nil {
		println(err)
	}

	fmt.Println("Write complete!")
}

func makeUID() string {
	// Should generate a random ID
	return "PLACEHOLDER UID"
}
