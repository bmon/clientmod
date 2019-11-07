package hellomod

import "fmt"

const version = "2.1.1"

func Hello() string {
	return fmt.Sprintf("Hello, modules, version: %s", version)
}
