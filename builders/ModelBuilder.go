package builders

import (
	"os"
	"sync"

	"encoding/json"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelBuilder(model types.ModelType, wg sync.WaitGroup) {
	defer logger.Trace(logger.Enter())

	wg.Add(1)

	fileName := flags.FlagOutputPath + "/" + "Model.json"
	if fp, err := os.Create(fileName); err != nil {
		panic(err)
	} else {
		defer fp.Close()

		if jsonVal, err := json.MarshalIndent(model, "", "  "); err != nil {
			panic(err)
		} else {
			fp.WriteString(string(jsonVal))

			logger.Log("Created file: " + fileName)
		}
	}
}
