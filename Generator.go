package main

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"

	"github.com/ftcjeff/ConfigurationProcessor/generators"
)

func Generator(model types.ModelType) (types.ModelType, error) {
	defer logger.Trace(logger.Enter())

	modelAChan := make(chan types.ModelAType)
	modelBChan := make(chan types.ModelBType)
	modelCChan := make(chan types.ModelCType)
	serverChan := make(chan []types.ServerType)

	go generators.ModelAGenerator(model, modelAChan)
	go generators.ModelBGenerator(model, modelBChan)
	go generators.ModelCGenerator(model, modelCChan)
	go generators.ServerGenerator(model, serverChan)

	model.Models.ModelA = <-modelAChan
	model.Models.ModelB = <-modelBChan
	model.Models.ModelC = <-modelCChan
	model.Models.Servers = <-serverChan

	return model, nil
}
