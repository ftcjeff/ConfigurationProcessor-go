package main

import (
	"fmt"

	"github.com/ftcjeff/ConfigurationProcessor/loaders"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func Loader() (types.ModelType, error) {
	defer logger.Trace(logger.Enter())

	var model types.ModelType

	definition := loaders.DefinitionLoader()
	model.Definition = definition

	logger.Log(fmt.Sprintf("%+v", model))

	model = loadAllElements(model)

	return model, nil
}

func loadAllElements(model types.ModelType) types.ModelType {
	defer logger.Trace(logger.Enter())

	/*

		definition := model.Definition()
		elements := definition.Elements()

		elementChannel := make(chan types.ElementChannel, len(elements))

	*/

	return model
}
