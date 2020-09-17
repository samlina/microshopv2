# User Service

This is the User service

Generated with

```
micro new --namespace=sam.gxc --alias=user --type=service --plugin=regisry=etcd:broker=nat user-service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: sam.gxc.service.user
- Type: service
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-service
```

Build a docker image
```
make docker
```