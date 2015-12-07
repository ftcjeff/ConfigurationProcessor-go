package builders

import (
	"os"
	"sort"
	"sync"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelBBuilder(model types.ModelType, wg sync.WaitGroup) {
	defer logger.Trace(logger.Enter())

	modelB := model.Models.ModelB

	fileName := flags.FlagOutputPath + "/" + "ModelB.txt"
	fp, _ := os.Create(fileName)
	defer fp.Close()

	sort.Strings(modelB.Components)

	for _, b := range modelB.Components {
		fp.WriteString(b + "\n")
	}

	logger.Log("Created file: " + fileName)
}
