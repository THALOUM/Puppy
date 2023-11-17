package puppy

import (
	"github.com/THALOUM/dog"
)

func Bark() string {
	return "woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}
