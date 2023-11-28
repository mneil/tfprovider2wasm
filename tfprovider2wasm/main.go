//go:build !wasm

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/mneil/tfprovider2wasm/resource/ec2"
)

func main() {
	mock := request.GetMock()
	mock.On("*", "*").Do(func (r *request.Request) {
		fmt.Println("Call: ", r.ClientInfo.ServiceName, r.Operation.Name)
		fmt.Println("PARAMS", r.Params)
		fmt.Println("RESPONSE", r.Data)
		fmt.Println("")
		// call my jsfun with args
	})
	ec2.ResourceInstanceCreate("{\"ami\":\"ami-2023\",\"instance_type\":\"t3.micro\",\"disable_api_termination\":true}")
}
