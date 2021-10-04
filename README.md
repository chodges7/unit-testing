# Unit testing for Liatio's Automated-testing bootcamp section

## Usage
- To spin up the two containers use docker-compose with the following command:
    - `docker-compose up -d --build`
- To see the logs use this command:
    - `docker-compose logs -f`
- Finally, it might be beneficial to also have this command open:
    - `watch docker-compose ps`
- Jenkins should be spun up in `http://localhost:8080`
- Artifactory should be spun up in `http://localhost:8082`

- Finally you'll configure the jenkins job to point to this git repo and run the following commands:
	- TEST stage: `go test -v`
	- BUILD stage: `go build`