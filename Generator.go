package main

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func Generator(model types.ModelType) (types.ModelType, error) {
	defer logger.Trace(logger.Enter())

	return model, nil
}
