package builders

import (
	"os"
	"sort"
	"sync"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelCBuilder(model types.ModelType, wg sync.WaitGroup) {
	defer logger.Trace(logger.Enter())

	modelC := model.Models.ModelC

	fileName := flags.FlagOutputPath + "/" + "ModelC.txt"
	if fp, err := os.Create(fileName); err != nil {
		panic(err)
	} else {
		defer fp.Close()

		sort.Strings(modelC.NodeNames)
		lastName := ""

		for _, c := range modelC.NodeNames {
			if c != lastName {
				fp.WriteString(c + "\n")
				lastName = c
			}
		}

		logger.Log("Created file: " + fileName)
	}
}
