/*****************************************************************************
*  Author:           Brett Smith
*  Email:            brettsmith578@gmail.com
*  Label:            P03
*  Title:            Program 3 - Image ASCII Art
*  Course:           CMPS 4143 - PLC
*  Semester:         Fall 2023
*
*  Description:
*        This program uses a package, img_mod, with functions that allow the
*		 user to download an image, acquire the RGB values of the image, grayscale
*		 the image, and add colored text to an image.
*
*  Usage:
*        ColorText, Colors, GetPic, and Grayscale files are imported from the img_mod
*		 package. The files contain the functions DownloadPic(), GetRGB(), Grayscale(),
*		 and ColorText() respectively. DownloadPic() downloads an image from the internet,
*		 GetRGB() calculates the RGB values of the image, Grayscale() uses the RGB and
*		 transforms the image to grayscale, and ColorText() allows for colored text to be
*		 added to an image.
*
*  Files:
*        main.go 	  			: main driver file for the program
*		 go.mod  	  			: contains module's properties required to run the program
*		 go.sum  	 			: contains checksums for dependencies
*		 colors.jpg	 			: the original image downloaded
*		 colors_labeled.jpg	    : example of colored text on an image
*		 colors_gray_scale.jpg  : grayscale of downloaded image
*		 color_pixel_counts.txt : RGB values of downloaded image's pixels
*****************************************************************************/

package main

import (
	"fmt"

	"github.com/bsmith578/img_mod/ColorText"
	"github.com/bsmith578/img_mod/Colors"
	"github.com/bsmith578/img_mod/GetPic"
	"github.com/bsmith578/img_mod/Grayscale"
)

func main() {
	fmt.Println("Hello")
	GetPic.DownloadPic()
	Colors.GetRGB()
	Grayscale.Grayscale()
	ColorText.ColorText()
}
