package starter

import (
	"errors"
	"fmt"
	"log"
)

var (
	BUG_EXISTS = true
)

func main() {
	result, err := nonBuggy("Nonbuggy", "First Time")
	if err != nil {
		log.Fatal("FATAL ERROR OH NO!!")
	}
	fmt.Println("[main] Result from first step:", result)

	result, err = buggy()
	if err != nil {
		log.Fatal("FATAL ERROR OH NO!!")
	}
	fmt.Println("[main] Result from second step:", result)

	result, err = nonBuggy("Nonbuggy", "Second Time")
	if err != nil {
		log.Fatal("FATAL ERROR OH NO!!")
	}
	fmt.Println("[main] Result from third step:", result)
}

func buggy() (string, error) {
	if BUG_EXISTS {
		return "oops", errors.New("oops, found a bug")
	} else {
		return "Hello, World!", nil
	}
}

func nonBuggy(one string, two string) (string, error) {
	fmt.Println("--> Running 'NormalActivity' with params, '" + one + "' and '" + two + "'")
	return one + ", " + two + "!", nil
}
