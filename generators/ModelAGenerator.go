package generators

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelAGenerator(model types.ModelType, channel chan types.ModelAType) {
	defer logger.Trace(logger.Enter())

	var modelA types.ModelAType

	definition := model.Definition
	for _, element := range definition.Elements {
		modelA.FileNames = append(modelA.FileNames, element.File)
	}

	channel <- modelA
}
