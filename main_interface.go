package main

import (
	"github.com/Equanox/gotron"
)

func main() {
	// Create a new browser window instance
	window, err := gotron.New("interface/build")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 900
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "3kChat"
	window.WindowOptions.MinWidth = 600
	window.WindowOptions.MinHeight = 400
	// window.WindowOptions.TitleBarStyle =  "customButtonsOnHover"

	// Start the browser window.
	// This will establish a golang <=> nodejs bridge using websockets,
	// to control ElectronBrowserWindow with our window object.
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Open dev tools must be used after window.Start
	window.OpenDevTools()

	// Wait for the application to close
	<-done
}
