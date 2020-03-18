.PHONY: build
build: proto
	go build -o ./bin/

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/msa/msa.proto

.PHONY: run
run:
	./bin/omo-msa-startkit

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.startkit StartKit.Call '{"name":"John"}'

.PHONY: tcall
tcall:
	go build -o ./bin/ ./client
	./bin/client

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
