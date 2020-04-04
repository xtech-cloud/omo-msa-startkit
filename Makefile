APP_NAME := omo-msa-startkit
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: build
build: proto
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/startkit/echo.proto
	/mnt/c/_wsl/protoc.exe --proto_path=./ --csharp_out=./proto/startkit --grpc_out=./proto/startkit --plugin=protoc-gen-grpc=c:/_wsl/grpc_csharp_plugin.exe proto/startkit/echo.proto
	protoc --proto_path=./ --java_out=./proto/startkit --plugin=protoc-gen-grpc-java=/usr/bin/protoc-gen-grpc-java proto/startkit/echo.proto
	protoc --proto_path=./ --grpc-java_out=nano:./proto/startkit --plugin=protoc-gen-grpc-java=/usr/bin/protoc-gen-grpc-java proto/startkit/echo.proto
	mv proto/startkit/omo/msa/startkit/* proto/startkit/
	rm -rf proto/startkit/omo

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.startkit Echo.Call '{"name":"John"}'

.PHONY: tester
tester:
	go build -o ./bin/ ./tester

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build . -t omo-msa-startkit:latest
