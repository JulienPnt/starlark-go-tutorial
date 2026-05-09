package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
	"os"
)

const (
	starScript  = "./greet.star"
	greetFunc   = "greet"
	myFirstName = "Julien"
	myLastName  = "Pnt"
)

func main() {
	thread := starlark.Thread{Name: greetFunc}
	opts := syntax.LegacyFileOptions()
	objs, err := starlark.ExecFileOptions(opts, &thread, starScript, nil, nil)
	if err != nil || objs == nil {
		fmt.Printf("Error on ExecFileOptions: %s\n", err.Error())
		os.Exit(1)
	}
	greet, ok := objs[greetFunc]
	if greet == nil || !ok {
		fmt.Printf("Error on %s into %s script\n", greetFunc, starScript)
		os.Exit(2)
	}
	value, err := starlark.Call(&thread, greet, starlark.Tuple{starlark.String(myFirstName), starlark.String(myLastName)}, nil)
	if err != nil {
		fmt.Printf("Error on Call: %s\n", err.Error())
		os.Exit(3)
	}
	fmt.Printf("Result: %s\n", value)
}
