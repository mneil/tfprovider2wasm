package resource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/conns"
	"github.com/hashicorp/terraform-provider-aws/provider"
)

func UnmarshalMap(input string) (map[string]any, error) {
	var raw map[string]any
	err := json.Unmarshal([]byte(input), &raw)
	if err != nil {
		return map[string]any{}, err
	}
	return raw, nil
}

func NewClient(ctx context.Context) *conns.AWSClient {
	schemaProvider, err := provider.New(ctx)
	if err != nil {
		panic(err)
	}

	diags := schemaProvider.Configure(ctx, terraform.NewResourceConfigRaw(map[string]interface{}{
		"region":                				"us-east-1",
		"access_key": 									"a",
		"secret_key": 									"b",
		"skip_credentials_validation": 	true,
	}))
	if diags.HasError() {
		fmt.Printf("Warn %+v\n", diags)
	}
	client, ok := schemaProvider.Meta().(*conns.AWSClient)
	if client.Session, err = session.NewSession(); err != nil {
		panic(err)
	}
	if !ok {
		panic("unable to cast garbage tf meta interface to AWSClient because hashi doesn't understand interfaces or the adapter pattern")
	}
	return client
}
