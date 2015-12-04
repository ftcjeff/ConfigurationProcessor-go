package main

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func Builder(model types.Model) (types.Model, error) {
	defer logger.Trace(logger.Enter())

	return model, nil
}
