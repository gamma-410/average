package main

import (
	"flag"
	"fmt"
)

var nationalLang int = 63
var socialStudies int = 50
var science int = 65

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	switch arg {
	case "total":
		total()
	case "average":
		average()
	default:
		fmt.Println("Please specify an argument.")
	}
}

func total() {
	total := nationalLang + socialStudies + science
	fmt.Println(total)
}

func average() {
	total := nationalLang + socialStudies + science
	average := total / 3
	fmt.Println(average)
}
