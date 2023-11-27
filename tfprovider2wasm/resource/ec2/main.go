package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/ec2"
	"github.com/mneil/tfprovider2wasm/resource"
)

func resourceInstanceCreate(args string) {
	content, err := resource.UnmarshalMap(args)
	if err != nil {
		fmt.Printf("unable to convert to json %s\n", err)
		return
	}
	// mock := request.GetMock()
	// mock.On("*", "*").Do(func (r *request.Request) {
	// 	fmt.Println("Call: ", r.ClientInfo.ServiceName, r.Operation.Name)
	// 	fmt.Println("PARAMS", r.Params)
	// 	fmt.Println("RESPONSE", r.Data)
	// 	fmt.Println("")
	// 	// call my jsfun with args
	// })
	ctx := context.Background()
	rsrc := ec2.ResourceInstance()
	data := schema.NewResourceDataFromResource(rsrc)
	for key, val := range content {
		// fmt.Printf("{'%s': %+v}\n", key, val)
		err := data.Set(key, val)
		if err != nil {
			fmt.Printf("unable to set %s\n", key)
			return
		}
	}
	client := resource.NewClient(ctx)
	ec2.ResourceInstanceCreate(ctx, data, client)

	mock := request.GetMock()
	fmt.Println(mock.ExpectedCalls)
	// mock.On("*", "*").Do(func (r *request.Request) {
	// 	// call my jsfun with args
	// })
}
