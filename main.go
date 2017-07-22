package main

import (
	"fmt"
	"flag"
	"github.com/bananaumai/camel2snake/converter"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Println(converter.Convert(arg))
	}
}
