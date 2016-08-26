package main

import (
	"fmt"
	"github.com/irlndts/go-rpn"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: parameter expression not informed")
		return
	}

	input := os.Args[1]

	if ok,_ := regexp.MatchString("[0-9/*-+.]+$", input); !ok {
		fmt.Println("Error: invalid expression")
		return
	}

	result,err := rpn.Calc(input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
