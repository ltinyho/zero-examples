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

.PHONY: build
push-shorturl: shorturl
	docker-compose push shorturl

.PHONY: build
push-transform: transform
	docker-compose push transform

.PHONY: build
build: shorturl transform

.PHONY: push
push: build
	docker-compose push

.PHONY: dev
dev:
	docker-compose up
