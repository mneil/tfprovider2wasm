//go:build !wasm

package ec2

func ResourceInstanceCreate(args string) {
	resourceInstanceCreate(args)
}
