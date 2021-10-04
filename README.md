# Unit testing for Liatio's Automated-testing bootcamp section

## Usage
- First build the docker container
	- `docker build -t automated-testing:0.1.0 ./`
- Then create a volume to keep the data from the jenkins container
    - `docker create volume jenkins-data`
- Then you'll run the docker container on port 8080.
	- `docker run -p 8080:8080 -v jenkins-data:/var/jenkins_home automated-testing:0.1.0`
- You'll need to grab the jenkins key from the continaer
    - `docker exec ${container name} cat /home/myuser/.jenkins/secrets/initialAdminPassword`

- Finally you'll configure the jenkins job to point to this git repo and run the following commands:
	- TEST stage: `go test -v`
	- BUILD stage: `go build`
