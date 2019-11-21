# Shippydemo Service

This is the Shippydemo service

Generated with

```
micro new github.com/alactic/shippydemo --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.shippydemo
- Type: srv
- Alias: shippydemo

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

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
./shippydemo-srv
```

Build a docker image
```
make docker
```