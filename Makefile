.PHONY: clean build_docker deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -o main.go -o build/quotes_api

build_docker:
	export GO111MODULE=on
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main.go -o build/quotes_api
	docker build -t quotes_api .

clean:
	rm -rf ./build

deploy: clean build_docker
	docker push quotes_api

start_docker: clean build_docker
	docker run --env-file .env -p 7000:7000 quotes_api

start_docker_compose: clean build_docker
	docker-compose up --build

start_local: clean build
	build/quotes_api
