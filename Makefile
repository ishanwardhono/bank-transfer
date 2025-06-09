.PHONY: install clean

install:
	docker-compose up -d

deploy:
	docker-compose up -d --no-deps --build app

clean:
	docker-compose down -v

run-test:
	go generate ./...
	go mod tidy
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out