# Application commands
server:
	go run main.go

build:
	go build -o server main.go

test:
	go test ./...

# Docker commands
docker-build:
	docker build -t vietnam-zipcode-api .

docker-run:
	docker run -p 8080:8080 vietnam-zipcode-api

docker-stop:
	docker stop $$(docker ps -q --filter ancestor=vietnam-zipcode-api)

# Clean up
clean:
	rm -f main

.PHONY: server build test docker-build docker-run docker-stop clean