package greetings

import (
	"fmt"
)

func PrintName(name string) (string) {
	fmt.Println(fmt.Sprintf("Hi, %v. Welcome from greeting!", name))

	return name
}
