
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/msa/msa.proto

.PHONY: build
build: proto

	go build 

.PHONY: test
test:
	micro call omo.msa.startkit StartKit.Call '{"name":"John"}'

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
