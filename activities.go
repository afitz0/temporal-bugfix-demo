package starter

import (
	"fmt"
	"go.temporal.io/sdk/temporal"
)

type Activities struct{}

var (
	HAS_BUG = false
)

func (a *Activities) BuggyActivity() (string, error) {

	if HAS_BUG {
		return "oops", temporal.NewApplicationError("oops, found a bug", "Error")
	} else {
		return "Hello, World!", nil
	}
}

func (a *Activities) NormalActivity(greeting string, name string) (string, error) {
	fmt.Println("Running 'NormalActivity' with params, '" + greeting + "' and '" + name + "'")
	return greeting + " " + name + "!", nil
}
