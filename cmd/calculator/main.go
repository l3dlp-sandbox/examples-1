// Package main launches the calculator example directly
package main

import "github.com/fyne-io/examples/calculator"
import "github.com/fyne-io/fyne/desktop"

func main() {
	app := desktop.NewApp()

	calculator.Show(app)
}
