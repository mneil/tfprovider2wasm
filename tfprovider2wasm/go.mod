module github.com/mneil/tfprovider2wasm

go 1.20

// replace github.com/hashicorp/go-cty => github.com/zclconf/go-cty v1.14.1

replace github.com/hashicorp/terraform-plugin-sdk/v2 => ../terraform-plugin-sdk

replace github.com/terraform-providers/terraform-provider-aws => ../terraform-provider-aws

replace github.com/hashicorp/terraform-provider-aws => ../terraform-provider-aws

replace github.com/aws/aws-sdk-go => ../aws-sdk-go-mock

// replace github.com/aws/aws-sdk-go-v2 => ./mod/aws/aws-sdk-go-v2

require (
	github.com/aws/aws-sdk-go v1.46.4
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.29.0
	github.com/hashicorp/terraform-provider-aws v0.0.0-00010101000000-000000000000
)
