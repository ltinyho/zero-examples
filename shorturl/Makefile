.PHONY: shorturl
shorturl:
	rm -rf api/shorturl
	GOOS=linux go build -o ./api/shorturl ./api/
	docker-compose build shorturl

.PHONY: transform
transform:
	rm -rf rpc/transform/transform
	GOOS=linux go build -o ./rpc/transform ./rpc/transform
	docker-compose build transform

.PHONY: dev
dev:
	docker-compose up
