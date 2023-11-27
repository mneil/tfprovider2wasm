//go:build wasm

package main

import (
	"syscall/js"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/mneil/tfprovider2wasm/resource/ec2"
)

func main() {
	// // fmt.Print("hello from wasm")
	// // js.FuncOf(Transform)
	// // TODO: How can we expose this without keeping it live the entire time?
	// // do we need to kill it or is this fine?
	//
	mock :=request.GetMock()
	mock.FakeOptions(
		options.WithRandomMapAndSliceMaxSize(3),
		options.WithRandomMapAndSliceMinSize(1),
	)
	js.Global().Get("TerraformAwsProvider").Get("ec2").Set("ResourceInstanceCreate", ec2.ResourceInstanceCreate())

	<-make(chan bool)
}
