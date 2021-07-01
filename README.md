# road-to-k8s
- [x] 1. Create your own Git repository for the project. Use a name of your choosing

- [x] 2. Write an application in the language of your choice that meets the following requirements:
* - [x] The app starts up and runs indefinitely until terminated
* - [x] The app is writing periodically a log message for every 10 seconds
* - [x] The message should contain the current time and a random string of your choice

- [ ] 3. Build the app into an Docker image and let it run in a container

- [ ] 4.Define a `deployment.yaml` with your app to run it in k8s
* - [ ] Play with it in K8S

- [ ] 5. Add an API that meets the following requirements:
* - [ ] Disable the log statements from task 1 with a boolean environment variable
* - [ ] Add the following endpoints:
  * - [ ] Request 1:
    * - [ ] “Get me the current time and date”
    * - [ ] GET request to /time
    * - [ ] response: {"time": "11:54", "date": "30-06-2021"}
  * - [ ] Request 2:
    * - [ ] “Send a text and write it in the logs”
    * - [ ] POST request to /story
    * - [ ] response: {"success": true, "message": "Han shot first"}
  * - [ ] All request and response content-types should be `application/json`
  * - [ ] Test it all out with a local docker container
  * - [ ] Hint: use “Postman” (or curl!) to define and send the requests 

- [ ] 6. Rebuild the application and deploy it in k8s. 
* - [ ] Reach your application running inside the cluster with your new API
* - [ ] Examine the response and the logs of your application
* - [ ] Scale your deployment to 5
* - [ ] Send some requests
* - [ ] Examine the logs again


Hints:
* Make use of existing frameworks
* Build, deploy and test often
* Name variables and components appropriately
* Write a Dockerfile
* Try Nodeport (`kubectl expose deployment [..]`)
* Try port-forward (`kubectl port-forward [..]`)
* Get help when you’re stuck :) 
