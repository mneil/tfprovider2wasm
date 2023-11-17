package main

import (
	// "context"
	// "encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/request"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	// "github.com/hashicorp/terraform-provider-aws/provider"

)

func main() {
	request.GetMock()
	fmt.Println("hello world")
}

// func unmarshalMap(input string) (map[string]any, error) {
// 	var raw map[string]any
// 	err := json.Unmarshal([]byte(input), &raw)
// 	if err != nil {
// 		return map[string]any{}, err
// 	}
// 	return raw, nil
// }

// func NewClient(ctx context.Context) interface{} {
// 	request.GetMock()
// 	schemaProvider, err := provider.New(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	diags := schemaProvider.Configure(ctx, terraform.NewResourceConfigRaw(map[string]interface{}{
// 		"region":                "us-east-1",
// 		"aws_access_key_id":     "a",
// 		"aws_secret_access_key": "b",
// 	}))
// 	if diags.HasError() {
// 		fmt.Printf("Errors %+v\n", diags)
// 		// panic("unable to configure provider")
// 	}
// 	// data := &schema.ResourceData{}
// 	// data.Set("access_key", "a")
// 	// schemaProvider.ConfigureContextFunc(ctx, data)
// 	client := schemaProvider.Meta()
// 	if client.Session, err = session.NewSession(); err != nil {
// 		panic(err)
// 	}
// 	// if !ok {
// 	// 	panic("unable to cast garbage tf meta interface to AWSClient because hashi doesn't understand interfaces or the adapter pattern")
// 	// }
// 	return client
// 	// client := &conns.AWSClient{
// 	// 	Partition: "aws",
// 	// 	Region:    "us-east-1",
// 	// 	// Service:   "ec2",
// 	// 	AccountID: "012345678910",
// 	// 	// Calls:     make([]*Call, 0),
// 	// 	Session: &session.Session{},
// 	// 	ServicePackages: ,
// 	// }
// 	// return client
// }

// // package main

// // import (
// // 	"context"
// // 	"encoding/json"
// // 	"fmt"
// // 	"syscall/js"

// // 	"github.com/aws/aws-sdk-go/aws/request"
// // 	"github.com/aws/aws-sdk-go/aws/session"
// // 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// // 	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

// // 	// "github.com/mneil/tfprovider2was/mod/hashicorp/terraform-provider-aws/conns"
// // 	"github.com/mneil/tfprovider2wasm/mod/hashicorp/terraform-provider-aws/ec2"
// // 	"github.com/mneil/tfprovider2wasm/mod/hashicorp/terraform-provider-aws/provider"
// // )

// // // var provider js.Value

// // func main() {
// // 	js.Global().Get("TerraformAwsProvider").Get("ec2").Set("ResourceInstanceCreate", ResourceInstanceCreate())
// // 	// fmt.Print("hello from wasm")
// // 	// js.FuncOf(Transform)
// // 	// TODO: How can we expose this without keeping it live the entire time?
// // 	// do we need to kill it or is this fine?
// // 	<-make(chan bool)
// // }

// // func unmarshalMap(input string) (map[string]any, error) {
// // 	var raw map[string]any
// // 	err := json.Unmarshal([]byte(input), &raw)
// // 	if err != nil {
// // 		return map[string]any{}, err
// // 	}
// // 	return raw, nil
// // }
// // func NewClient(ctx context.Context) interface{} {
// // 	request.GetMock()
// // 	schemaProvider, err := provider.New(ctx)
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	diags := schemaProvider.Configure(ctx, terraform.NewResourceConfigRaw(map[string]interface{}{
// // 		"region":                "us-east-1",
// // 		"aws_access_key_id":     "a",
// // 		"aws_secret_access_key": "b",
// // 	}))
// // 	if diags.HasError() {
// // 		fmt.Printf("Errors %+v\n", diags)
// // 		// panic("unable to configure provider")
// // 	}
// // 	// data := &schema.ResourceData{}
// // 	// data.Set("access_key", "a")
// // 	// schemaProvider.ConfigureContextFunc(ctx, data)
// // 	client, ok := schemaProvider.Meta()
// // 	if client.Session, err = session.NewSession(); err != nil {
// // 		panic(err)
// // 	}
// // 	if !ok {
// // 		panic("unable to cast garbage tf meta interface to AWSClient because hashi doesn't understand interfaces or the adapter pattern")
// // 	}
// // 	return client
// // 	// client := &conns.AWSClient{
// // 	// 	Partition: "aws",
// // 	// 	Region:    "us-east-1",
// // 	// 	// Service:   "ec2",
// // 	// 	AccountID: "012345678910",
// // 	// 	// Calls:     make([]*Call, 0),
// // 	// 	Session: &session.Session{},
// // 	// 	ServicePackages: ,
// // 	// }
// // 	// return client
// // }

// // func ResourceInstanceCreate() js.Func {
// // 	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
// // 		if len(args) != 1 {
// // 			return "Invalid no of arguments passed"
// // 		}
// // 		content, err := unmarshalMap(args[0].String())
// // 		if err != nil {
// // 			fmt.Printf("unable to convert to json %s\n", err)
// // 			return err.Error()
// // 		}
// // 		ctx := context.Background()
// // 		resource := ec2.ResourceInstance()
// // 		data := schema.NewResourceDataFromResource(resource)
// // 		for key, val := range content {
// // 			// fmt.Printf("{'%s': %+v}\n", key, val)
// // 			err := data.Set(key, val)
// // 			if err != nil {
// // 				fmt.Printf("unable to set %s\n", key)
// // 				return err.Error()
// // 			}
// // 		}
// // 		client := NewClient(ctx)
// // 		ec2.ResourceInstanceCreate(ctx, data, client)

// // 		return ""

// // 		// out, err := json.Marshal((*client).Calls)

// // 		// // fmt.Printf("diags %+v\n", diags)
// // 		// // fmt.Printf("out %s\n", out)
// // 		// // fmt.Printf("diags to js %+v\n", js.ValueOf(client.Calls))
// // 		// return string(out[:])
// // 		// // return content
// // 	})
// // 	return jsonFunc
// // }
