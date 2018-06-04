.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/NiluPlatform/go-nilu/cmd/swarm
	go build -o bin/geth     github.com/NiluPlatform/go-nilu/cmd/geth
	go build -o bin/bootnode github.com/NiluPlatform/go-nilu/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/NiluPlatform/go-nilu/swarm/...
