LINUX_AMD64 = GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on

start:
	@docker-compose -f docker-compose.yml up -d --build

stop:
	@docker-compose -f docker-compose.yml down

test:
	go test -cover ./...

.PHONY: server test