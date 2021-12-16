package main

import (
	"github.com/cjfinnell/golang-proto-demo/commonfields"
	"github.com/cjfinnell/golang-proto-demo/onebig"
	"github.com/golang/protobuf/proto"
)

func tryOneBig() {
	// Make it
	topWrapMsg := &onebig.TopWrapper{
		Foo: "foo",
		Bar: 12,
		Baz: true,
		Nested1: &onebig.Nested1{
			First: "inside nested 1",
		},
		// No Nested2 !
	}

	raw, err := proto.Marshal(topWrapMsg)
	if err != nil {
		panic("Failed to marshal top wrapper")
	}

	// Read it
	readTopWrapMsg := &onebig.TopWrapper{}
	if err := proto.Unmarshal(raw, readTopWrapMsg); err != nil {
		panic("Failed to unmarshal top wrapper")
	}
	if readTopWrapMsg.Nested1 != nil {
		println("top wrap has nested 1 (expected), value:", readTopWrapMsg.Nested1.First)
	} else {
		println("top wrap has no nested 1 (bad!)")
	}

	if readTopWrapMsg.Nested2 != nil {
		println("top wrap has nested 2 (bad!), value:", readTopWrapMsg.Nested1.First)
	} else {
		println("top wrap no nested 2 (expected)")
	}
}

func main() {
	tryOneBig()

	// TODO: common fields
	_ = commonfields.Source1{}
	_ = commonfields.Source2{}
}
