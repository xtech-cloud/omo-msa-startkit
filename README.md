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

### Alpine
`For Alpine v3.11`

- 环境配置

    ```bash
    ~# apk add --no-cache autoconf automake libtool curl make g++ unzip alpine-sdk
    ```

- 安装Go

    ```bash
    ~# apk add go --no-cache --repository=https://mirrors.aliyun.com/alpine/v3.11/community/
    ```
    
    在/etc/profile中加入一行
    ```bash
    export GOPROXY=https://goproxy.io,direct
    ```

- 安装Protobuf

    ```bash
    ~# apk add --no-cache protoc --repository=http://mirrors.aliyun.com/alpine/v3.11/main/
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
- 安装grpc-java
    ```bash
    ~# apk add --no-cache grpc-java --repository=http://mirrors.aliyun.com/alpine/edge/testing
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

    添加以下代码到internal/handler/meta.go中
    ```go
    func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        //支持跨域
        helper.ServeCORS(w, r)
        if r.Method == "OPTIONS" {
            return
        }
        ...
    }
    ```

    修改internal/helper/helper.go中的以下代码
    ```go
    func ServeCORS(w http.ResponseWriter, r *http.Request) {                        
        set := func(w http.ResponseWriter, k, v string) {                           
            if v := w.Header().Get(k); len(v) > 0 {                                                                
                return                                                                                             
            }                                                                                                                                                                                                     
            w.Header().Set(k, v)                                                                                   
        }                                                                                                          
        
        if origin := r.Header.Get("Origin"); len(origin) > 0 {                      
            set(w, "Access-Control-Allow-Origin", origin)                           
        } else {                                                                                                   
            set(w, "Access-Control-Allow-Origin", "*")                              
        }                                                                                                          
        
        headers := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
       
        if reqHeaders := r.Header.Get("Access-Control-Request-Headers"); len(reqHeaders) > 0 {
            headers = headers + "," + reqHeaders                                                                   
        }                                                                                                          
        
        set(w, "Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
        set(w, "Access-Control-Allow-Headers", headers)                             
    }
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

### Windows

- 安装nuget
    这一步主要用于

    下载[NuGet](https://www.nuget.org/downloads)

    在命令行中运行以下命令

    ```bash
    > .\nuget.exe install Grpc.Tools -Version 2.27.0
    ```

    完成后将Grpc.Tools.2.27.0\tools\windows_x64\目录下的protoc.exe和grpc_csharp_plugin.exe拷贝到c:\_wsl目录下

- 安装protoc-gen-javalite
    下载https://repo1.maven.org/maven2/com/google/protobuf/protoc-gen-javalite/3.0.0/protoc-gen-javalite-3.0.0-windows-x86_64.exe，拷贝到c:\_wsl目录下，并改名为protoc-gen-javalite.exe
- 安装protoc-gen-web
    从https://github.com/grpc/grpc-web/releases下载protoc-gen-grpc-web-1.0.7-windows-x86_64.exe，拷贝到c:\_wsl目录下，并改名为protoc-gen-grpc-web.exe

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
    ~#  MICRO_REGISTRY=consul micro api --namespace=omo
    ```

    使用POST访问
    ```bash
    curl http://localhost:8080/msa/startkit/Echo/Call?name=Asim
    curl -H 'Content-Type: application/json' -d '{"name": "Asim"}' http://localhost:8080/startkit/Echo/Call
    ```

- WEB界面
    ```bash
    ~#  MICRO_REGISTRY=consul micro web --namespace=omo
    ```

- 构建Docker镜像

    ```
    make docker
    ```
