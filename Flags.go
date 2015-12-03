package main

import (
	"flag"
)

var FlagInputPath string
var FlagOutputPath string

func init() {
	defer Trace(Enter())

	flag.StringVar(&FlagInputPath, "input", "config", "The path to the configuration files")
	flag.StringVar(&FlagOutputPath, "output", "output", "The path to the generated files")

	flag.Parse()

	Log("FlagInputPath: " + FlagInputPath)
	Log("FlagOutputPath: " + FlagOutputPath)
}
