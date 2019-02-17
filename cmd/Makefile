build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o=./bin/TDKserver main.go

docker-build:
	make build
	docker build ./ --tag hatobus/tdkserver:latest
