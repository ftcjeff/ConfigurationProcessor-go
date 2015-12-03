package main

import (
	"encoding/json"
	"io/ioutil"
)

func ConfigLoader(model Model) (model, error) {
	defer Trace(Enter())

	raw, err := ioutil.ReadFile(FlagInputPath)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(raw, &model)

	return model, nil
}
