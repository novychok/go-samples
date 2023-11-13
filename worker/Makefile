run:
	sudo docker compose up --remove-orphans --build

build-master:
	@go build -o ./bin/master/master ./cmd/master/*.go

build-worker:
	@go build -o ./bin/worker/worker ./cmd/worker/*.go

rmaster: build-master
	@./bin/master/master

rworker: build-worker
	@./bin/worker/worker