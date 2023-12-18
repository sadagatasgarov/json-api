build:
	go build -o bin/gobank

run: build
	./bin/gobank
test:
	go test -v ./..

docker:
	@docker build -t sadagatasgarov/my-json-app .
	@docker push sadagatasgarov/my-json-app:latest
	@docker compose up -d
