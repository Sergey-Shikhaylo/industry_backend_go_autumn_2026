package main

import "strings"

func greet(name string) string {
	trim_name := strings.TrimSpace(name)
	if trim_name == "" {
		trim_name = "World"
	}
	return "Hello, " + trim_name + "!"
}
