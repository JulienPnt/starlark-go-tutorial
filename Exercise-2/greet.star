#!/usr/bin/env starlark


MESSAGE="Hello from Starlark"

def greet(first_name, last_name):
  return MESSAGE + " " + first_name + " " + last_name + " !"

greet("Julien", "Pnt")
