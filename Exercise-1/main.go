package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
	"os"
)

const (
	starlark_script = "greet.star"
	object_name     = "greet"
	my_name         = "Julien"
)

func main() {

	thread := &starlark.Thread{Name: "greet"}
	opts := syntax.LegacyFileOptions()
	globals, err := starlark.ExecFileOptions(opts, thread, starlark_script, nil, nil)
	if err != nil {
		fmt.Printf("Error: on ExecFileOptions: %s", err.Error())
		os.Exit(1)
	}
	greet, ok := globals[object_name]
	if greet == nil || !ok {
		fmt.Printf("Error: %s object not defined: %s", object_name, starlark_script)
		os.Exit(2)
	}
	out, err := starlark.Call(thread, greet, starlark.Tuple{starlark.String(my_name)}, nil)
	if err != nil {
		fmt.Printf("Error: on ExecCall: %s", err.Error())
		os.Exit(3)
	}
	fmt.Printf("Result: %s\n", out)

	// OUT OF STATEMENT
	// TO TEST starlark.StringDict
	//message := globals["MESSAGE"]
	//fmt.Printf("%s: Julien\n", message)
	os.Exit(0)
}
