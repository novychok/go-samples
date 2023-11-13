build:
	@go build -o ./bin/master/master ./cmd/master/*.go
	@go build -o ./bin/worker/worker ./cmd/worker/*.go

run: build
	@./bin/master/master
	@./bin/worker/worker