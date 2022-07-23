generate:
	docker build --target gen -t proto-gen .
	docker run --name proto-gen proto-gen
	docker cp proto-gen:/proto/gen .
	docker rm proto-gen

lint:
	docker build --target lint -t proto-gen .