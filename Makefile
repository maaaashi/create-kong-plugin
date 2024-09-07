build:
	go build -o kong-plugin-gen
docker-build:
	docker build -t kong-plugin-gen -f environments/Dockerfile .