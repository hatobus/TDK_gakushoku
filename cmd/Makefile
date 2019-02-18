build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o=./bin/TDKserver main.go

docker-build:
	make build
	docker build ./ --tag hatobus/tdkserver:latest

deploy:
	make docker-build
	docker push hatobus/tdkserver

up-gcp:
	sudo git pull
	sudo docker pull hatobus/tdkserver
	sudo docker-compose down 
	sudo docker-compose up
