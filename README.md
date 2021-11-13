# Server Example

A simple server for prototyping made using golang.  
This is not production ready code, it is just for proof of concept, which means I don't write all the required code (ex. data validation), I just write enough to get the most of what I am trying to implement/learn.

---

## Dependences

- `go 1.16+`

## Commands

- Run Server: &nbsp;`make run`
- Run Test: &nbsp;`make test`

---

## Running the project inside a docker container

### Prerequisites

- Make sure you have docker installed for your operating system. [Download Link](https://docs.docker.com/get-docker/)
- make sure that you are not running the server on your local machine (aka. have run the `go run .`)

### Steps to run

1. Open Terminal

2. Navigate to the project directory

3. Build the docker image (with tag `server-example`)

   ```console
   docker build . -t server-example
   ```

4. Check that the image was craeted  
   To check that the docker image was created run the command
   ```console
    docker images
   ```
   This should output something like this
   ```console
    REPOSITORY       TAG       IMAGE ID       CREATED         SIZE
    server-example   latest      ...          1 minute ago   323MB
   ```
   If you do not see the seccond line go back to step 3.
   Make sure that you type the command correclty, or even better copy it and paste it in your terminal.
5. Run the docker image
   We need to run the image while making sure to bound a port so that we can access our services
   ```console
    docker run -p 8080:8080 --name  server-example -d server-example
   ```
6. Make sure your image is running
   To check that the docker image is running run the command
   ```console
     docker ps
   ```
   This should output something like this
   ```console
    CONTAINER ID   IMAGE            COMMAND     CREATED          STATUS          PORTS                    NAMES
        ...        server-example   "/server"   26 seconds ago   Up 25 seconds   0.0.0.0:8080->8080/tcp   server-example
   ```
   If you do not see the seccond line go back to step 5.
   Make sure that you type the command correclty, or even better copy it and paste it in your terminal.
7. Test your app.  
   To test the app you can use the routes, as described in the `Routes`.

## Routes

You can find the available routes inside the `request-examples` folder.  
If you are using VS Code as your editor you can install the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extention and click on the send request button above each Request to execute the request.
If you are not using VS Code you can use [Postman](https://www.postman.com/)

## Progress

- Create a server using Golang, with a simple endpoint that logs in the server console  
  &nbsp; Here I wanted:
  1. to learn how to create a simple Server in golang without using any external libraries
  2. to experiment with different file structures
  3. to learn how to use make files
- Create a image manipulation enpoint (I chose gausian blur of an image)  
  &nbsp; Here I wanted:
  1. to learn how to handle file uploading via form data post request
  2. do simple image processing using go
  3. use goroutines and channels to run intensive code on separate thread
- Create a simple sum endpoint
  &nbsp; Here I wanted:
  1. to learn how to write unit tests (test utils.Sum)
  2. to learn how to write integration tests (test the whole endpoint)
  3. to experiment with recurrsion in golang
- Create a CI/CD environment
  &nbsp; Here I wanted:
  1. to learn how linting works in go
  2. to learn how to write a worklow file that checks linting, ability to build and that the tests pass

## Want to learn

- gRPC (Microservices)
  &nbsp; I want to look more into that
