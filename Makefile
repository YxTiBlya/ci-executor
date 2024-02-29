build-executor:
	@cd ./executor && go build -o ../bin/executor ./cmd/main/main.go

run-executor: build-executor
	@./bin/executor --cfg config.yaml