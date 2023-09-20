/*****************************************************************************
*  Author:           Brett Smith
*  Email:            brettsmith578@gmail.com
*  Label:            P02
*  Title:            Program 2 - Baby Steps
*  Course:           CMPS 4143 - PLC
*  Semester:         Fall 2023
*
*  Description:
*        This program uses a package named imageManipulator to draw a rectangle
*		 around an image and save it as a new image.
*
*  Usage:
*        NewImageManipulatorWithImage(), DrawRectangle(), and SaveToFile() are
*		 imported into a main.go file from imageManipulator.go in the imageManipulator
*		 package. This allows us to create an instance of an imageManipulator and pass
*		 in an existing image. We then draw a rectangle around the image and save it as
*		 a new image.
*
*  Files:
*        main.go 	  : main driver file for the program
*		 go.mod  	  : contains module's properties required to run the program
*		 go.sum  	  : contains checksums for dependencies
*		 mustangs.jpg : image file before manipulation
*		 mustangs.png : image file after manipulation
*****************************************************************************/

// imageManipulator/main.go

package main

import (
	"fmt"
	"myimageapp/imageManipulator"
)

func main() {

	// Create an ImageManipulator instance with an existing image
	im, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the image to a file
	im.SaveToFile("mustangs.png")
}
