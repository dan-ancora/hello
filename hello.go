package main

import (
	"fmt"
	"github.com/dan-ancora/stringutil"
)

func main() {

	const firstMessage = "First Hello World examples on GIT 2"

	fmt.Println(firstMessage)
	fmt.Println(stringutil.Reverse(firstMessage))
}
