.PHONY: dc run test lint

dc:
	docker-compose up  --remove-orphans --build -v

run:
	go build -o app cmd/main.go && HTTP_ADDR=:8080 ./app

test:
	go test -race ./...

lint:
	golangci-lint run