// imageManipulator/main.go

package main

import "example.com/go/pkg/mod/github.com/bsmith578/4143-!p!l!c@v0.0.0-20230918220341-105bb03207ec/Assignments/P02/imagemod/imageManipulator"

func main() {
	// Create an ImageManipulator instance
	im := imageManipulator.NewImageManipulator(800, 600)

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the image to a file
	im.SaveToFile("mustangs.png")
}
