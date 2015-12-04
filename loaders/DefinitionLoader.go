package loaders

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func DefinitionLoader(definitionChannel chan types.Definition) {
	defer logger.Trace(logger.Enter())

	raw, err := ioutil.ReadFile(flags.FlagInputPath)
	if err != nil {
		panic(err)
	}

	var d types.Definition
	err = json.Unmarshal(raw, &d)
	if err != nil {
		panic(err)
	}
	logger.Log(d.ToString())

	definitionChannel <- d
}
