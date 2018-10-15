package main

import (
	"log"

	"bitbucket.org/kit-systems/dari/pkg/parser"
)

func main() {
	p, err := parser.New(parser.Options{})
	if err != nil {
		log.Fatal(err)
	}
	if err = p.Run(); err != nil {
		log.Fatal(err)
	}
}
