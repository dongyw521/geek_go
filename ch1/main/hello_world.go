package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World! " + os.Args[1])
	}
	// fmt.Println("Hello World")
	// fmt.Println(os.Args)
	os.Exit(-1)
}
