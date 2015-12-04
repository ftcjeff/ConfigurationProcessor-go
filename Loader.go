package main

import (
	"github.com/ftcjeff/ConfigurationProcessor/loaders"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func Loader() (types.Model, error) {
	defer logger.Trace(logger.Enter())

	var model types.Model

	definitionChannel := make(chan types.Definition)
	go loaders.DefinitionLoader(definitionChannel)
	definition := <-definitionChannel
	model.SetDefinition(definition)

	logger.Log(model.ToString())

	model = loadAllElements(model)

	return model, nil
}

func loadAllElements(model types.Model) types.Model {
	defer logger.Trace(logger.Enter())

	/*

		definition := model.Definition()
		elements := definition.Elements()

		elementChannel := make(chan types.ElementChannel, len(elements))

	*/

	return model
}
