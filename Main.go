package main

func main() {
	defer Trace(Enter())

	var model Model
	var err error

	if model, err = Loader(); err != nil {
		LogFatal(err.Error())
	}

	if model, err = Builder(model); err != nil {
		LogFatal(err.Error())
	}

	if model, err = Generator(model); err != nil {
		LogFatal(err.Error())
	}
}
