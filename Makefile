.PHONY: run protos

run:
	@go run main.go

protos:
	protoc --go_out=toplevel top-level-wrapper.proto
