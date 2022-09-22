image:
	docker-compose build --no-cache

start:
	docker-compose run --rm app

release:
	docker build --file docker/release/Dockerfile -t sample-grpc-server-with-go:v1.0.0 .