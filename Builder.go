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

	builders.ModelBuilder(model, wg)
	builders.ModelABuilder(model, wg)
	builders.ModelBBuilder(model, wg)
	builders.ModelCBuilder(model, wg)
	builders.HostsBuilder(model, wg)

	wg.Wait()

	return model, nil
}
