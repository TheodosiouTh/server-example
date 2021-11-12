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

## Routes

You can find the available routes inside the `request-examples` folder.  
If you are using VS Code as your editor you can install the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extention and click on the send request button above each Request to execute the request.

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
