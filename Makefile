
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/msa/msa.proto

.PHONY: build
build: proto

	go build 

.PHONY: test
test:
	go test -v 

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
