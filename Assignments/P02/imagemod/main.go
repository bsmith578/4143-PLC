// imageManipulator/main.go

package main

import "github.com/bsmith578/4143-PLC/tree/main/Assignments/P02/imageManipulator"

func main() {
	// Create an ImageManipulator instance
	im := imageManipulator.NewImageManipulator(800, 600)

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the image to a file
	im.SaveToFile("mustangs.png")
}
