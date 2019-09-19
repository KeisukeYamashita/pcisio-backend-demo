.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/server \
		github.com/KeisukeYamashita/pcisio-server/server

.PHONY: docker
docker:
	docker build . -t payment-server