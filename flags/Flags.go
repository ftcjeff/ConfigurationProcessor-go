package flags

import (
	"flag"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
)

var FlagInputPath string
var FlagOutputPath string

func init() {
	defer logger.Trace(logger.Enter())

	flag.StringVar(&FlagInputPath, "input", "config/definition.json", "The path to the primary configuration file")
	flag.StringVar(&FlagOutputPath, "output", "output", "The path to the generated files")

	flag.Parse()

	logger.Log("FlagInputPath: " + FlagInputPath)
	logger.Log("FlagOutputPath: " + FlagOutputPath)
}
