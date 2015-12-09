package loaders

import (
	"encoding/json"
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

	networkFileName := d.Config.Network
	raw, err = ioutil.ReadFile(networkFileName)
	if err != nil {
		panic(err)
	}

	var nets types.NetworkType
	err = json.Unmarshal(raw, &nets)
	if err != nil {
		panic(err)
	}

	return d
}
