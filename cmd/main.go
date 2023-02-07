// Package main stores the main application logic and is responsilbe for dependency
// injecting the application functions
package main

// main will create and initialize an instance of application,
// it will panic if it fails
// a destructor is deferred
// and the main loop goes on until a quit signal is recieved or escape is pressed
func main() {
	app := newApp()
	err := app.initialize()
	if err != nil {
		panic(err)
	}

	defer app.destructor()

	for app.IsRunning() {
		app.processInput()
		app.update()
		app.render()
	}
}
