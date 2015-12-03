package main

func main() {
	defer Trace(Enter())

	var model Model
	var err error

	if model, err = Builder(); err != nil {
		LogFatal(err.Error())
	}

	if err := Generator(model); err != nil {
		LogFatal(err.Error())
	}
}
