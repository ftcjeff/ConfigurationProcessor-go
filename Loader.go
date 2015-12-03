package main

func Loader() (Model, error) {
	defer Trace(Enter())

	if model, err := ConfigLoader(); err != nil {
		LogFatal(err.Error)
	}

	Log(model)

	return model, nil
}
