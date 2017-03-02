# Authentication Tutorial

I went through this tutorial to learn about JWTs and the Golang net/http package. The original tutorial can be found [here](https://auth0.com/blog/authentication-in-golang/).

## Building and Running the Project

This project includes a Dockerfile which can be used to build and run the code. You need to have [Docker](https://www.docker.com/) and [Golaing](https://golang.org/) installed on your system for the setup to work. First, retrive the project using:

```
go get https://github.com/cmvandrevala/authentication-golang-tutorial.git
```

This will add the project under ```$GOPATH/src/github.com/cmvandrevala/authentication-golang-tutorial```. Next, navigate to the project directory and run:

```
docker build -t="<name of image>" .
```

where ```<name of image>``` is the name that you want to give the Docker image for this app. To run the app, use the following command:

```
docker run -d -p 3000:3000 <name of image>
```

You can then navigate to ```localhost:3000``` in your browser to view the landing page of this app.

In order to stop the Docker container, run ```docker ps``` to get a list of the current running containers. Find the container with an IMAGE of ```<name of image>```, retrieve its Container ID, and then run:

```
docker stop <CONTAINER ID>
```
