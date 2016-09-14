package img

import (
	"github.com/disintegration/imaging"
	"image"
)

func IncreaseImageBrightness(image *image.Image) *image.NRGBA {
	dstImage := imaging.AdjustBrightness(*image, 10) // increase image brightness by 10%
	return dstImage
}
