package starter

import (
	"errors"
	"fmt"
)

type Activities struct{}

func (a *Activities) StepOne(input string) (string, error) {
	fmt.Println("--> [Step One]  Input: " + input)
	return input, nil
}

func (a *Activities) StepTwo(input string) (string, error) {
	hasBug := true

	if hasBug {
		return "oops", errors.New("oops error. Hard to fix. Be warned.")
	} else {
		fmt.Println("--> [Step Two]  Input: " + input)
	}

	return input, nil
}

func (a *Activities) StepThree(input string) (string, error) {
	fmt.Println("--> [Step Three]  Input: " + input)
	return input, nil
}
