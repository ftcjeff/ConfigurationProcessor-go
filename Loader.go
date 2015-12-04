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

	model = loadAllElements(model)

	logger.Log(fmt.Sprintf("%+v", model))

	return model, nil
}

func loadAllElements(model types.ModelType) types.ModelType {
	defer logger.Trace(logger.Enter())

	definition := model.Definition
	elements := definition.Elements

	elementChannel := make(chan types.ElementChannel, len(elements))

	for k, v := range elements {
		go loaders.ElementLoader(k, v.File, elementChannel)
	}

	for i := 0; i < len(elements); i++ {
		elem := <-elementChannel

		idx := elem.Id
		elements[idx].Data = elem.Element
	}

	return model
}
