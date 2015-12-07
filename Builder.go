package main

import (
	"sync"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"

	"github.com/ftcjeff/ConfigurationProcessor/builders"
)

func Builder(model types.ModelType) (types.ModelType, error) {
	defer logger.Trace(logger.Enter())

	var wg sync.WaitGroup

	builders.ModelABuilder(model, wg)
	builders.ModelBBuilder(model, wg)
	builders.ModelCBuilder(model, wg)

	wg.Wait()

	return model, nil
}
