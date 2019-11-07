package commonmod

import (
	"fmt"

	"github.com/bmon/hellomod/v2"
)

func SaySomething() string {
	return fmt.Sprintln("commonmod", hellomod.Hello())
}
