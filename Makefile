build:
	@go build -o ./bin/executor ./cmd/main/main.go

run: build
	@./bin/executor --cfg config.yaml

docker-build:
	@docker build -t ci-executor .
