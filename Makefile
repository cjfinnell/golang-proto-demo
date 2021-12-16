.PHONY: run protos

run:
	@go run main.go

protos:
	protoc --go_out=toplevel top-level-wrapper.proto
	protoc --go_out=commonfields common-fields.proto
