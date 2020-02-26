# Micro Service Agent

This is the Micro Service Agent

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: xtc.omo.srv.omo
- Type: srv
- Alias: omo

## Dependencies


```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

You may need use goproxy 
```
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

```
make build
```


Run the service
```
./omo-msa-startkit
```

Build a docker image
```
make docker
```
