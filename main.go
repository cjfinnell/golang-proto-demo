package main

import (
	"github.com/cjfinnell/golang-proto-demo/commonfields"
	"github.com/cjfinnell/golang-proto-demo/toplevel"
)

func main() {
	_ = toplevel.TopWrapper{}
	_ = commonfields.Source1{}
	_ = commonfields.Source2{}
}
