package puppy

import (
	"fmt"

	"github.com/THALOUM/dog"
)

func Bark() string {
	return "woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func From1() {
	fmt.Println("I'm from version v1.0.0")
}
