build_master:
	@go build -o ./bin/master ./cmd/master/*.go

build_worker:
	@go build -o ./bin/worker ./cmd/worker/*.go

master: build_master
	@./bin/master

worker: build_worker
	@./bin/worker

test:
	@go test -v ./...