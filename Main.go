package main

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func main() {
	defer logger.Trace(logger.Enter())

	var model types.ModelType
	var err error

	if model, err = Loader(); err != nil {
		logger.LogFatal(err.Error())
	}

	if model, err = Builder(model); err != nil {
		logger.LogFatal(err.Error())
	}

	if model, err = Generator(model); err != nil {
		logger.LogFatal(err.Error())
	}
}
