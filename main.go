package main

import (
	"flag"
	"fmt"

	"go.uber.org/zap"

	"bitbucket.org/kit-systems/dari/pkg/parser"
)

var (
	debug = flag.Bool("debug", false, "Debug mode")
)

func main() {
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Print(err)
	}
	p, err := parser.New(parser.Options{
		Logger: logger,
		Debug:  *debug,
	})
	if err != nil {
		logger.Fatal("new parser", zap.Error(err))
	}
	err = p.Run()
	if err != nil {
		logger.Fatal("run parser", zap.Error(err))
	}
}
