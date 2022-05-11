LINUX_AMD64 = GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on

build:
	@$(LINUX_AMD64) go build -a -v -tags musl -o bff api.go

start:
	@docker-compose -f docker-compose.yml up -d

stop:
	@docker-compose --env-file .local.env -f docker-compose.yml down

server:
	go run framework/cmd/api/api.go

test:
	go test -cover ./...

.PHONY: server test