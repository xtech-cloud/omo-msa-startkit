# Micro Service Agent

This is the Micro Service Agent

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

## Dependencies

`For Alpine v3.11`

Setup Envinorment

```bash
~# apk add --no-cache autoconf automake libtool curl make g++ unzip alpine-sdk
```

Install Go
```bash
~# apk add go --no-cache --repository=https://mirrors.aliyun.com/alpine/v3.11/community/
```

Install Protobuf
```bash
~# cd ~
~# git clone --branch v3.11.4 --depth=1 https://github.com/protocolbuffers/protobuf
~# cd protobuf
~# git submodule update --init --recursive
~# ./autogen.sh
~# ./configure
~# make
~# make install
~# ldconfig
```

Install protoc-gen-go
```bash
~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
~# go get -u github.com/golang/protobuf/protoc-gen-go
~# cp /root/go/bin/protoc-gen-go /usr/local/bin/
```

Install protoc-gen-micro
```bash
~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
~# go get github.com/micro/protoc-gen-micro
~# cp /root/go/bin/protoc-gen-micro /usr/local/bin/
```

Install Micro
```bash
~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
~# git clone --branch=v2.1.1 --depth=1 https://github.com/micro/micro
```

append follow code in main.go
```go
_ "github.com/micro/go-plugins/registry/consul/v2"
_ "github.com/micro/go-plugins/registry/kubernetes/v2"
```

```
~# cd micro
~# go install
~# cp /root/go/bin/micro /usr/local/bin/
```

Install Consul
```bash
~# apk add --no-cache consul --repository=http://mirrors.aliyun.com/alpine/edge/testing/
```

## Usage

A Makefile is included for convenience

- Build the binary

You may need use goproxy 
```
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

```
~# make build
```

- Run the service

run consul
```bash
~# consul agent -dev
```

run msa
```
~# ./omo-msa-startkit
```

test
```
~# export MICRO_REGISTRY=consul
~# make test
```


- Build a docker image

```
make docker
```
