package image

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"math"
)

type pixel struct {
	x int
	y int
}

const blurKernelSize = 25

func BlurImage(imageToBlur image.Image, standardDeviation float64, out chan []byte) {
	defer close(out)

	dimentions := imageToBlur.Bounds()
	imageResult := image.NewRGBA(dimentions)

	blurMask := getBlurMask(standardDeviation)

	for x := 1; x < imageResult.Bounds().Max.X; x += 1 {
		for y := 1; y < imageResult.Bounds().Max.Y; y += 1 {
			bluredPixel := getBluredPixel(imageToBlur, pixel{x: x, y: y}, blurMask)

			imageResult.Set(x, y, bluredPixel)
		}
	}

	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, imageResult, nil)
	out <- buffer.Bytes()
}

func getBlurMask(standardDeviation float64) [blurKernelSize][blurKernelSize]float64 {
	positionOffset := int(math.Abs(blurKernelSize / 2))
	gausianBlurKernel := [blurKernelSize][blurKernelSize]float64{}
	for x := 0; x < blurKernelSize; x += 1 {
		for y := 0; y < blurKernelSize; y += 1 {
			distanceFromCenter := pixel{x: x - positionOffset, y: y - positionOffset}
			sSquared := math.Pow(standardDeviation, 2)
			// https://en.wikipedia.org/wiki/Gaussian_blur
			gausianBlurKernel[x][y] = math.Exp(
				(-math.Pow(float64(distanceFromCenter.x), 2)-
					math.Pow(float64(distanceFromCenter.y), 2))/
					(2*sSquared),
			) /
				(2 * math.Pi * sSquared)
		}
	}
	return gausianBlurKernel
}

func getBluredPixel(image image.Image, coordinates pixel, blurMask [blurKernelSize][blurKernelSize]float64) color.RGBA64 {
	var bluredPixelR, bluredPixelG, bluredPixelB float64

	imagePositioningOffset := int(math.Floor(float64(len(blurMask)) / 2))
	for i := 0; i < len(blurMask); i += 1 {
		for j := 0; j < len(blurMask[i]); j += 1 {
			xOffset := i - imagePositioningOffset
			yOffset := j - imagePositioningOffset

			targetPixelX := coordinates.x + xOffset
			targetPixelY := coordinates.y + yOffset
			outOfBbounts := targetPixelX < 0 || targetPixelX > image.Bounds().Max.X ||
				targetPixelY < 0 || targetPixelY > image.Bounds().Max.Y
			if outOfBbounts {
				continue // ignore when out of bounds
			}

			r, g, b, _ := image.At(targetPixelX, targetPixelY).RGBA()
			bluredPixelR += float64(r) * blurMask[i][j]
			bluredPixelG += float64(g) * blurMask[i][j]
			bluredPixelB += float64(b) * blurMask[i][j]
		}
	}

	_, _, _, a := image.At(coordinates.x, coordinates.y).RGBA()
	return color.RGBA64{
		getUint16Value(bluredPixelR),
		getUint16Value(bluredPixelG),
		getUint16Value(bluredPixelB),
		uint16(a),
	}
}

func getUint16Value(pixelValue float64) uint16 {
	return uint16(math.Abs(pixelValue))
}
