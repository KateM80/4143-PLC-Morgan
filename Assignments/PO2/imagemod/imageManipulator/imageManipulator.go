// imagemod/imageManipulator/imageManipulator.go

package imageManipulator

import (
	"/github.com/KateM80/2143-OOP-morgan/blob/main/NewPF.jpg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image *NewPF.Context
}

// NewImageManipulator creates a new ImageManipulator instance.
func NewImageManipulator(width, height int) *ImageManipulator {
	img := NewPF.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

// SaveToFile saves the manipulated image to a file.
func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

// DrawRectangle draws a rectangle on the image.
func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
