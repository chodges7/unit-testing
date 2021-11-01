# Unit testing for Liatio's Automated-testing bootcamp section

## Usage
- This repo uses Jenkins and Artifactory containers. 
	- To spin up the two containers use docker-compose with the following command:
	   - `docker compose up -d --build`
- To see the logs use this command: (this is helpful to see the Jenkins key)
    - `docker compose logs -f`
- Finally, it might be beneficial to also have this command open:
    - `watch docker compose ps`
- Jenkins should be spun up in `http://localhost:8080`
- Artifactory should be spun up in `http://localhost:8081`

- Finally you'll configure the jenkins job to point to this git repo and run the following commands:
	- TEST stage: `go test -v`
	- BUILD stage: `go build`

---
# November 1st
- As of the above date, continuous delivery functionality has been added.
- The Jenkinsfile in the repo will run a `go build`, `go test`, and then finally a `goreleaser release --rm-dist`
	- Note: You will need to add a tag to the commit just before running the build in Jenkins.	
