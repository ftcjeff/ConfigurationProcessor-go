package main

func Loader() (Model, error) {
	defer Trace(Enter())

	var model Model

	//	if model, err := ConfigLoader(); err != nil {
	//		LogFatal(err.Error)
	//	}

	//	Log(model)

	return model, nil
}
