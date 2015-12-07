package generators

import (
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelCGenerator(model types.ModelType, channel chan types.ModelCType) {
	defer logger.Trace(logger.Enter())

	var modelC types.ModelCType

	definition := model.Definition

	for _, element := range definition.Elements {
		data := element.Data

		for _, tier := range data.Tiers {
			node := tier.Node

			modelC.NodeNames = append(modelC.NodeNames, node.Name)
		}
	}

	channel <- modelC
}
