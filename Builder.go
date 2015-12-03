package main

func Builder(model Model) (Model, error) {
	defer Trace(Enter())

	return model, nil
}
