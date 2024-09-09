build:
	go build -o kong-plugin-gen
docker-build:
	docker build -t kong-plugin-gen -f environments/Dockerfile .
docker-run:
	docker run -it --rm -p8000:8000 -p8001:8001 kong-plugin-gen