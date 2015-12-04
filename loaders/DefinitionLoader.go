package loaders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func DefinitionLoader() types.DefinitionType {
	defer logger.Trace(logger.Enter())

	raw, err := ioutil.ReadFile(flags.FlagInputPath)
	if err != nil {
		panic(err)
	}

	var d types.DefinitionType
	err = json.Unmarshal(raw, &d)
	if err != nil {
		panic(err)
	}
	logger.Log(fmt.Sprintf("%+v", d))

	return d
}
