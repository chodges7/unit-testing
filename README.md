# Unit testing for Liatio's Automated-testing bootcamp section

## Usage
* First build the docker container
	* `docker build -t automated-testing:0.1.0 ./`
* Then you'll run the docker container on port 8080.
	* `docker run -p 8080:8080 automated-testing:0.1.0`

* Finally you'll configure the jenkins job to point to this git repo and run the following commands:
	* TEST stage: `go test -v`
	* BUILD stage: `go build`
