package starter

import (
	"errors"
	"fmt"
	//"go.temporal.io/sdk/temporal"
)

type Activities struct{}

var (
	HAS_BUG = false
)

func (a *Activities) BuggyActivity() (string, error) {

	if HAS_BUG {
		return "oops", errors.New("oops error. Hard to fix. Be warned.")
	} else {
		return "Hello, World!", nil
	}
}

func (a *Activities) NormalActivity(greeting string, name string) (string, error) {
	fmt.Println("--> Running 'NormalActivity' with params, '" + greeting + "' and '" + name + "'")
	return greeting + " " + name + "!", nil
}
