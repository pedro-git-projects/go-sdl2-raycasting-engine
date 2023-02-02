package main

func main() {
	app := newApp()
	err := app.initialize()
	if err != nil {
		panic(err)
	}

	defer app.destructor()

	for app.Running() {
		app.processInput()
		app.update()
		app.render()
	}
}
