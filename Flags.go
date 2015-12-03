package main

import (
	"flag"
)

var FlagInputPath string
var FlagOutputPath string

func init() {
	defer Trace(Enter())

	flag.StringVar(&FlagInputPath, "input", "config/definition.json", "The path to the primary configuration file")
	flag.StringVar(&FlagOutputPath, "output", "output", "The path to the generated files")

	flag.Parse()

	Log("FlagInputPath: " + FlagInputPath)
	Log("FlagOutputPath: " + FlagOutputPath)
}
