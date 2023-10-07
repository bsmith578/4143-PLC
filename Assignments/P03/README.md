## P03 - Image ASCII Art
### Brett Smith
### Description:

This program creates 4 image manipulation modules that allows the user to read RGB values of an image,
add colored text to an image, grayscale an image, and download an image from the internet.

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
|   1   | [img_mod](https://github.com/bsmith578/img_mod) | Go module containing packages for certain image manipulations |
|   2   | [main.go](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/main.go) | main driver function for the program |
|   3   | [go.mod](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/go.mod) | contains module's properties required to run the program |
|   4   | [go.sum](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/go.sum) | contains checksums for dependencies |
|   5   | [colors.jpg](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/colors.jpg) | the original image downloaded |
|   6   | [colors_labeled.jpg](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/colors_labeled.jpg) | example of colored text on an image |
|   7   | [colors_gray_scale.jpg](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/colors_gray_scale.jpg) | grayscale of downloaded image |
|   8   | [color_pixel_counts.txt](https://github.com/bsmith578/4143-PLC/blob/main/Assignments/P03/ColorTest/color_pixel_counts.txt) | RGB values of downloaded image's pixels |

### Usage

To use this package, run: 
    - `go get github.com/bsmith578/img_mod/ColorText`
    - `go get github.com/bsmith578/img_mod/Colors`
    - `go get github.com/bsmith578/img_mod/GetPic`
    - `go get github.com/bsmith578/img_mod/Grayscale`

### Example commands

GetPic.DownloadPic() - Asks for URL of an image, downloads the image, and saves it

Colors.GetRGB() - Takes the downloaded image and calculates the RGB for each pixel

Grayscale.Grayscale() - Takes the RGB of the image and turns it to grayscale

ColorText.ColorText() - Shows an example of colored text on an image