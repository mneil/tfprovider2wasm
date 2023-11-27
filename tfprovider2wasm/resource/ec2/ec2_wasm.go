//go:build wasm

package ec2

import (
	"syscall/js"
)

func ResourceInstanceCreate() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		resourceInstanceCreate(args[0].String())

		return ""
		// out, err := json.Marshal((*client).Calls)

		// // fmt.Printf("diags %+v\n", diags)
		// // fmt.Printf("out %s\n", out)
		// // fmt.Printf("diags to js %+v\n", js.ValueOf(client.Calls))
		// return string(out[:])
		// // return content
	})
	return jsonFunc
}
