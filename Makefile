.PHONY: run protos

run:
	@go run main.go

protos:
	protoc --go_out=onebig one-big.proto
	protoc --go_out=commonfields common-fields.proto
