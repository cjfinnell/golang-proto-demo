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

func tryCommonFields() {
	// Make both
	source1 := &commonfields.Source1{
		First: "first",
		Common: &commonfields.CommonFields{
			Foo: "foo",
		},
	}
	raw1, err := proto.Marshal(source1)
	if err != nil {
		panic("Failed to marshal common fields 1")
	}
	source2 := &commonfields.Source2{
		Second: "second",
		Common: &commonfields.CommonFields{
			Foo: "foo",
		},
	}
	raw2, err := proto.Marshal(source2)
	if err != nil {
		panic("Failed to marshal common fields 2")
	}

	var rawMessages [][]byte
	rawMessages = append(rawMessages, raw1, raw2)

	// Read both
	for i, raw := range rawMessages {
		println("unmarshaling commonfields", i+1)
		readSource1 := &commonfields.Source1{}
		readSource2 := &commonfields.Source2{}

		if err := proto.Unmarshal(raw, readSource1); err == nil {
			println("found a source 1, with common:", readSource1.GetCommon().Foo)
			if i != 0 {
				panic("unmarshaled a 2 as a 1!!!")
			}
			continue
		}
		println("not a source 1, trying source 2")
		if err := proto.Unmarshal(raw, readSource2); err == nil {
			println("found a source 1, with common:", readSource2.GetCommon().Foo)
			if i != 1 {
				panic("unmarshaled a 1 as a 2!!!")
			}
			continue
		}
		panic("failed to unmarshal common fields impl")
	}
}

func main() {
	tryOneBig()
	tryCommonFields()
}
