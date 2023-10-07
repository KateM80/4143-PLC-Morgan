/*************************************************************************************************
*
*  Name:		Kate Morgan
*  P03:         Program 3 - Image Ascii Art
*  Course:      CMPS 4143
*
*  Description:
*        This program uses a package holding for modules that allow image manipulation. This is
*		 Done through the go get command to pull the package from another repo. Each module is used for
*		 Image manipulation in differnt ways
*
*  Usage:
*        - Run the the main file
*
*  Files:
*        main.go   			 :  Runs everything
*        go.mod	   			 :  Gives the path for the package and modules
*        go.sum	   			 :  Does something for go to be honest I dont know
*************************************************************************************************/

package main

import (
	"github.com/KateM80/img_mod/Colors"
	"github.com/KateM80/img_mod/Getpic"
	"github.com/KateM80/img_mod/Grayscale"
	"github.com/KateM80/img_mod/Text"
)

func main() {

	// Call function to get picture from URL
	Getpic.DownloadPicture()

	// Call function to process Pixel Colors
	Colors.PrintPixels()

	// Call function to grayscale image
	Grayscale.GrayScale()

	// Call function to print colored text to image
	Text.PrintText()
}
