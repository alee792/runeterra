GOPATH := $(HOME)/go

.PHONY: all
all: proto data

.PHONY: proto
proto: 
	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogo_out=.\
		proto/runeterra.proto

.PHONY: data
data:
	# go run ./cmd/gen core third_party/core/en_us/data/globals-en_us.json
	go run ./cmd/gen set third_party/set1/en_us/data/set1-en_us.json pkg/datadragon/set1.go 