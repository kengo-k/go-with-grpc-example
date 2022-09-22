image-dev:
	docker-compose build dev --no-cache

start-dev:
	docker-compose run --rm dev

release:
	docker build --file docker/release/Dockerfile -t sample-grpc-server-with-go:v1.0.0 .

start:
	docker-compose run -p 3000:3000 --rm production