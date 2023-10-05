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
