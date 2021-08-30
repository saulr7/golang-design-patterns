package statregy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

type PrintStrategy interface {
	Print() error
}

type ConsoleSquare struct{}

type ImageSquare struct {
	DestinationFilePath string
}

func (c *ConsoleSquare) Print() error {
	println("Square")
	return nil
}

func (t *ImageSquare) Print() error {

	width := 800
	height := 600

	origin := image.Point{0, 0}

	bgImg := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	// bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})

	squareImg := image.NewRGBA(square)

	draw.Draw(bgImg, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)

	if err != nil {
		return fmt.Errorf("Error opening image")
	}

	defer w.Close()

	if err = jpeg.Encode(w, bgImg, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	return nil

}
