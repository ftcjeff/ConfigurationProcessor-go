package builders

import (
	"os"
	"sync"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ModelABuilder(model types.ModelType, wg sync.WaitGroup) {
	defer logger.Trace(logger.Enter())

	modelA := model.Models.ModelA

	fileName := flags.FlagOutputPath + "/" + "ModelA.txt"
	fp, _ := os.Create(fileName)
	defer fp.Close()

	for _, a := range modelA.FileNames {
		fp.WriteString(a + "\n")
	}

	logger.Log("Created file: " + fileName)
}
