run: 
	go run .
build:
	go build 
build-image: 
	docker build . -t server-example
run-image: 
	docker run -p 8080:8080 --name  server-example -d server-example
stop-image: 
	docker stop server-example
remove-image:
	docker rm server-example
test: 
	go test -v ./...