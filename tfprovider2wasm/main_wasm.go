//go:build wasm

package main

import (
	"syscall/js"

	"github.com/mneil/tfprovider2wasm/resource/ec2"
)

func main() {
	// fmt.Print("hello from wasm")
	// js.FuncOf(Transform)
	// TODO: How can we expose this without keeping it live the entire time?
	// do we need to kill it or is this fine?
	js.Global().Get("TerraformAwsProvider").Get("ec2").Set("ResourceInstanceCreate", ec2.ResourceInstanceCreate())
	<-make(chan bool)
}
