GOPATH := $(HOME)/go

.PHONY: all
all: proto data

.PHONY: proto
proto: 
	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogo_out=.\
		proto/runeterra.proto

.PHONY: cards
cards:
	go run ./cmd/gen core third_party/core/en_us/data/globals-en_us.json
	go run ./cmd/gen set third_party/sets/set1/en_us/data/set1-en_us.json pkg/datadragon/set1.go 

.PHONY: unpack
unpack:
	unzip ./third_party/bundles/datadragon-core-en_us.zip -d ./third_party/core
	unzip ./third_party/bundles/datadragon-set1-lite-en_us.zip -d ./third_party/sets/set1