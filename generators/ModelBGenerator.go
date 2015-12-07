package generators

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelBGenerator(model types.ModelType, channel chan types.ModelBType) {
	defer logger.Trace(logger.Enter())

	var modelB types.ModelBType

	definition := model.Definition

	for _, element := range definition.Elements {
		data := element.Data

		for _, tier := range data.Tiers {
			modelB.Components = append(modelB.Components, tier.Components...)
		}
	}

	channel <- modelB
}
