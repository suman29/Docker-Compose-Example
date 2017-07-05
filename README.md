# Docker Compose Example

Set up for running two containers (1 App and 1 Nginx image) individually and linking them using docker compose
- App is written in golang, which has a get request handle
- nginx image which is pulled from docker hub and modified to redirect to my app

Code structure:
- app folder which contains Dockerfile for app and golang linux executable for app
- nginx folder which contains Dockerfile for nginx and configuration files to configure route for app

### Installation

> prerequisite: golang 1.6, glide and docker should be installed on your system.

Note: Executable is already checked in the code, you might not need to build it again.

Run glide install and build linux executable.

```
  $ env GOOS=linux GOARCH=amd64 go build -o app/docker-trial-linux
```

Bringing up containers:
```
  $ docker-compose build
  $ docker-compose up -d
```

Test the containers are working:
```
curl http://localhost/my-app/find
```


