#!/usr/bin/env starlark


MESSAGE="Hello from Starlark"

def greet(name):
  return MESSAGE + " " + name + " !"

greet("Julien")
