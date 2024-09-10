build:
	go build -o create-kong-plugin
docker-build:
	docker build -t create-kong-plugin -f environments/Dockerfile .
docker-run:
	docker run -it --rm -p8000:8000 -p8001:8001 create-kong-plugin