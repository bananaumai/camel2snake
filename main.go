package main

import (
	"fmt"
	"flag"
	"github.com/bananaumai/camel2snake/converter"
)

func main() {
	flag.Parse()
	//c := converter.NewConverter()
	for _, arg := range flag.Args() {
		//fmt.Println(c.Convert(arg))
		//fmt.Println(converter.Convert(arg))
		fmt.Println(converter.Convert(arg))
	}
}
