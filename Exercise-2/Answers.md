# Question 1 - How do you convert a Go string to  starlark.String ?

- starlark.String("golang string")

# Question 2 - What is the type of the result returned by  Call() ?

Call return a result types as starlark.Value which is an interface made to represent a Starlark object, namely any kinds of variables or functions or callback etc...

# Question 3 - How do you extract the Go value from a  starlark.Value ?

To extract Go values from starlark.Value:
- Type assertion: value.(starlark.String) then .GoString()
- For starlark.Int: use AsInt() or AsInt64()
- For starlark.Bool: use Bool()
- For collections (List, Dict): iterate with .Iterate() or access methods
