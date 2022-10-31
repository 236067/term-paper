# BikeWeb Service

This is the BikeWeb service

Generated with

```
micro new sss/BikeWeb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.BikeWeb
- Type: web
- Alias: BikeWeb

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./BikeWeb-web
```

Build a docker image
```
make docker
```
