package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const FILE_LOCATION = "shared/stories.json"

type StoryPost struct {
	Story string `json:"message"`
	UID   string `json:"uid"`
}

type StoryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UID     string `json:"uid"`
}

type StoryList struct {
	Stories map[string]StoryPost `json:"stories"`
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
	// initialise struct
	var storyList StoryList

	// read in json data
	jsonData, err := ioutil.ReadFile(FILE_LOCATION)
	if os.IsNotExist(err) {
		os.Create(FILE_LOCATION)
	} else if err != nil {
		fmt.Println(err)
	}

	// populate struct with json data
	json.Unmarshal(jsonData, &storyList)
	fmt.Println("Read json file")

	// add new story to json
	storyList.Stories[s.UID] = s

	// marshal back into byte array
	bytes, err := json.Marshal(storyList)

	if err != nil {
		fmt.Println(err)
	}

	// write back to file
	ioutil.WriteFile(FILE_LOCATION, bytes, 0644)
}

func makeUID() string {
	// Should generate a random ID
	return "PLACEHOLDER UID"
}
