# Question - 1 What type is returned by  ExecFileOptions() ?

ExecFileOptions returns a starlark.StringDict namely a map bounding a String as index to starlark.Value.
A starlark.Value is an interface describing a starlark objects namely every functions or variables into your starlark scripts.

# Question - 2 How do you handle execution errors ?

Execution errors are handled by checking the error return value from 
ExecFileOptions() and starlark.Call().

# Question - 3 What happens if the file doesn't exist ?

Non-existing file produces an empty starlark.StringDict and an error at the ExecFileOptions call.

