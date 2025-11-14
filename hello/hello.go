package main

import (
	"example.com/greetings"
	"fmt"
	"rsc.io/quote"
)

func say_hello() {
	fmt.Println(quote.Go())

	message := greetings.PrintName("Mike")

	fmt.Println(fmt.Sprintf("Hi, %v. Welcome from Hello module!", message))
}
