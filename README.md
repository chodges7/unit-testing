# Unit testing Go in Jenkins
---
## Description
This repo uses a Jenkins container instance to run a Jenkins pipeline. This pipeline then 
runs the unit tests associated with this go module. Then, the pipeline will release the
built contents to Github Releases.

## Usage
To use this repo, you will need to setup a Jenkins pipeline project pointing to this repo.

- First, spin up the container using docker compose with the following command:
	   - `docker compose up -d --build`
- To see the logs use this command: (this is helpful to see the Jenkins key)
    - `docker compose logs -f`
- It might be beneficial to also have this command open:
    - `watch docker compose ps`
- Jenkins should be spun up in `http://localhost:8080`
- You will also need to setup a github PAT (Personal Access Token) setup in the Jenkins instance.
The name of the PAT is described as `github-token` in the Jenkinsfile, but the name of the
credential can be changed to whatever you like.
- Finally, make a pipeline job that points to this repo.
- The Jenkinsfile in the repo will run the following stages:

```go
go build 						// build stage
go test  						// test stage
goreleaser release --rm-dist	// deploy to github releases
```
- Note: You will need to add a tag to the commit just before running the build in Jenkins or 
else it will thow an error.

	
