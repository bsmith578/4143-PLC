//imageManipulator/main.go

package main

import (
	"Assignments/P02/imagemod/imageManipulator"
)

func main() {
	// Create an ImageManipulator instance
	im := imageManipulator.NewImageManipulator(800, 600)

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the image to a file
	im.SaveToFile("mustangs.png")
}
