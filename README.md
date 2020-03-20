# Micro Service Agent

This is the Micro Service Agent

## 开始使用

- [配置](#配置)
- [依赖](#依赖)
- [用法](#用法)

## 配置

- MSA_REGISTRY_PLUGIN
    服务注册的插件，默认值为`consul`

- MSA_REGISTRY_ADDRESS
    服务注册的地址,默认值为`127.0.0.1:8500`

- MSA_CONFIG_DEFINE
    文件的配置
    ```json
    {	
        "source": "file",
        "prefix": "./runpath/",
        "key": "default.yaml"
    }	
    ```

    consul的配置
    ```json
    {	
        "source": "consul",
        "prefix": "/omo/msa/config",
        "key": "default.yaml",
        "address": [
            "127.0.0.1:8500"
        ]
    }	
    ```

## 依赖

`For Alpine v3.11`

- 环境配置

    ```bash
    ~# apk add --no-cache autoconf automake libtool curl make g++ unzip alpine-sdk
    ```

- 安装Go

    ```bash
    ~# apk add go --no-cache --repository=https://mirrors.aliyun.com/alpine/v3.11/community/
    ```

- 安装Protobuf

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

- 安装protoc-gen-go

    ```bash
    ~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
    ~# go get -u github.com/golang/protobuf/protoc-gen-go
    ~# cp /root/go/bin/protoc-gen-go /usr/local/bin/
    ```

- 安装protoc-gen-micro

    ```bash
    ~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
    ~# go get github.com/micro/protoc-gen-micro
    ~# cp /root/go/bin/protoc-gen-micro /usr/local/bin/
    ```

- 安装Micro

    ```bash
    ~# export GOPROXY=https://mirrors.aliyun.com/goproxy/
    ~# git clone --branch=v2.1.1 --depth=1 https://github.com/micro/micro
    ```

    将以下两行代码加入到main.go中
    ```go
    _ "github.com/micro/go-plugins/registry/consul/v2"
    _ "github.com/micro/go-plugins/registry/etcdv3/v2"
    ```

    ```
    ~# cd micro
    ~# go install
    ~# cp /root/go/bin/micro /usr/local/bin/
    ```

- 安装Consul

    ```bash
    ~# apk add --no-cache consul --repository=http://mirrors.aliyun.com/alpine/edge/testing/
    ```

- 安装etcd 

    ```bash
    ~# cd ~
    ~# git clone --branch=v3.4.4 --depth=1 https://github.com/etcd-io/etcd
    ~# cd etcd
    ~# make build
    ~# cp ./bin/etcd /usr/local/bin/
    ~# cp ./bin/etcdctl /usr/local/bin/
    ```

## 用法

A Makefile is included for convenience

- 构建二进制文件

    You may need use goproxy 
    ```
    export GOPROXY=https://mirrors.aliyun.com/goproxy/
    ```

    ```
    ~# make 
    ```

- 运行服务

    run consul
    ```bash
    ~# consul agent -dev
    ```

    run msa
    ```
    ~# make run
    ```

- 测试


    单次调用服务
    ```bash
    ~# make call
    ```

    循环调用服务
    ```bash
    ~# make tester
    ~# ./bin/tester
    ```

- HTTP 调用

    需要先开启API网关
    ```bash
    ~#  MICRO_REGISTRY=consul micro api --namespace=omo.msa
    ```

    使用POST访问
    ```bash
    curl http://localhost:8080/startkit/Echo/Call?name=Asim
    curl -H 'Content-Type: application/json' -d '{"name": "Asim"}' http://localhost:8080/startkit/Echo/Call
    ```

- 构建Docker镜像

    ```
    make docker
    ```
