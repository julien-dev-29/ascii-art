package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a string")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too much arguments")
		return
	}
	
}
