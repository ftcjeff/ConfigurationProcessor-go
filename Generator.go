package main

func Generator(model Model) (Model, error) {
	defer Trace(Enter())

	return model, nil
}
