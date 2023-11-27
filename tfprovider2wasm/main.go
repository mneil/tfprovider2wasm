//go:build !wasm

package main

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/mneil/tfprovider2wasm/resource/ec2"
)

// var provider js.Value

func main() {
	// js.Global().Get("TerraformAwsProvider").Get("ec2").Set("ResourceInstanceCreate", ResourceInstanceCreate())
	// // fmt.Print("hello from wasm")
	// // js.FuncOf(Transform)
	// // TODO: How can we expose this without keeping it live the entire time?
	// // do we need to kill it or is this fine?
	// <-make(chan bool)
	mock :=request.GetMock()
	mock.FakeOptions(
		options.WithRandomMapAndSliceMaxSize(3),
		options.WithRandomMapAndSliceMinSize(1),
	)
	ec2.ResourceInstanceCreate("{}")
}

// func unmarshalMap(input string) (map[string]any, error) {
// 	var raw map[string]any
// 	err := json.Unmarshal([]byte(input), &raw)
// 	if err != nil {
// 		return map[string]any{}, err
// 	}
// 	return raw, nil
// }
// func NewClient(ctx context.Context) *conns.AWSClient {
// 	schemaProvider, err := provider.New(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	diags := schemaProvider.Configure(ctx, terraform.NewResourceConfigRaw(map[string]interface{}{
// 		"region":                				"us-east-1",
// 		"access_key": 									"a",
// 		"secret_key": 									"b",
// 		"skip_credentials_validation": 	true,
// 	}))
// 	if diags.HasError() {
// 		fmt.Printf("Warn %+v\n", diags)
// 	}
// 	client, ok := schemaProvider.Meta().(*conns.AWSClient)
// 	if client.Session, err = session.NewSession(); err != nil {
// 		panic(err)
// 	}
// 	if !ok {
// 		panic("unable to cast garbage tf meta interface to AWSClient because hashi doesn't understand interfaces or the adapter pattern")
// 	}
// 	return client
// }

// func ResourceInstanceCreate(args string) {

// 	content, err := unmarshalMap(args)
// 	if err != nil {
// 		fmt.Printf("unable to convert to json %s\n", err)
// 		return
// 	}
// 	// mock := request.GetMock()
// 	// mock.On("*", "*").Do(func (r *request.Request) {
// 	// 	fmt.Println("Call: ", r.ClientInfo.ServiceName, r.Operation.Name)
// 	// 	fmt.Println("PARAMS", r.Params)
// 	// 	fmt.Println("RESPONSE", r.Data)
// 	// 	fmt.Println("")
// 	// 	// call my jsfun with args
// 	// })
// 	ctx := context.Background()
// 	resource := ec2.ResourceInstance()
// 	data := schema.NewResourceDataFromResource(resource)
// 	for key, val := range content {
// 		// fmt.Printf("{'%s': %+v}\n", key, val)
// 		err := data.Set(key, val)
// 		if err != nil {
// 			fmt.Printf("unable to set %s\n", key)
// 			return
// 		}
// 	}
// 	client := NewClient(ctx)
// 	ec2.ResourceInstanceCreate(ctx, data, client)

// 	mock := request.GetMock()
// 	fmt.Println(mock.ExpectedCalls)
// 	// mock.On("*", "*").Do(func (r *request.Request) {
// 	// 	// call my jsfun with args
// 	// })
// }
