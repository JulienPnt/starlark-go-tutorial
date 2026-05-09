package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
	"os"
)

func main() {

	thread := &starlark.Thread{Name: "greet"}
	opts := syntax.LegacyFileOptions()
	globals, err := starlark.ExecFileOptions(opts, thread, "greet.star", nil, nil)
	if err != nil {
		os.Exit(1)
	}
	greet := globals["greet"]
	out, err := starlark.Call(thread, greet, starlark.Tuple{starlark.String("Julien")}, nil)
	if err != nil {
		os.Exit(2)
	}
	fmt.Printf("%s\n", out)
	os.Exit(0)
}
