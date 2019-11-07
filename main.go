package main

import (
	"fmt"

	"github.com/bmon/commonmod"
	"github.com/bmon/hellomod/v2"
)

func main() {
	fmt.Println("I'm the client!", hellomod.Hello())
	fmt.Println("Here's what common has to say:", commonmod.SaySomething())
}
