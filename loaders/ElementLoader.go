package loaders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ElementLoader(id int, fileName string, channel chan types.ElementChannel) {
	defer logger.Trace(logger.Enter())

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var e types.ElementDataType
	err = json.Unmarshal(raw, &e)
	if err != nil {
		panic(err)
	}
	logger.Log(fmt.Sprintf("%+v", e))

	var ec types.ElementChannel
	ec.Id = id
	ec.Element = e

	channel <- ec
}
