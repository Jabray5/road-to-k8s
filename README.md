# road-to-k8s
- [x] 1. Create your own Git repository for the project. Use a name of your choosing

- [x] 2. Write an application in the language of your choice that meets the following requirements:
* - [x] The app starts up and runs indefinitely until terminated
* - [x] The app is writing periodically a log message for every 10 seconds
* - [x] The message should contain the current time and a random string of your choice

- [x] 3. Build the app into an Docker image and let it run in a container

- [x] 4.Define a `deployment.yaml` with your app to run it in k8s
* - [x] Play with it in K8S

- [x] 5. Add an API that meets the following requirements:
* - [x] Disable the log statements from task 1 with a boolean environment variable
* - [x] Add the following endpoints:
  * - [x] Request 1:
    * - [x] “Get me the current time and date”
    * - [x] GET request to /time
    * - [x] response: {"time": "11:54", "date": "30-06-2021"}
  * - [x] Request 2:
    * - [x] “Send a text and write it in the logs”
    * - [x] POST request to /story
    * - [x] response: {"success": true, "message": "Han shot first"}
  * - [x] All request and response content-types should be `application/json`
  * - [x] Test it all out with a local docker container
  * - [x] Hint: use “Postman” (or curl!) to define and send the requests 

- [x] 6. Rebuild the application and deploy it in k8s. 
* - [x] Reach your application running inside the cluster with your new API
* - [x] Examine the response and the logs of your application
* - [x] Scale your deployment to 5
* - [x] Send some requests
* - [x] Examine the logs again

## Extended Features
- [ ] Return the time in a specified timezone
- [ ] Re-enable automatic logs every 10s using threads
- [ ] Implement some kind of tests

- [ ] 7. Expand the /story endpoint with the following:
* - [ ] The response for POST /story should now return a unique identifier in the JSON as well
* - [ ] Add to the /story endpoint the functionality to retrieve past stories by the unique identifier. The response should include the time and date when the story has been created as well.
* - [ ] The stories should be persisted / add a storage solution of your choice
* - [ ] If one or all pods crash/restart the stories should still be retrievable
