package canvas

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"runtime/debug"
	"time"

	"github.com/steveoc64/memdebug"
)

// Raster describes a raster image area that can render in a Fyne canvas
type Raster struct {
	baseObject

	Generator func(w, h int, r *Raster) image.Image // Render the raster image from code

	Translucency float64 // Set a translucency value > 0.0 to fade the raster

	img draw.Image
}

// Alpha is a convenience function that returns the alpha value for a raster
// based on it's Translucency value. The result is 1.0 - Translucency.
func (r *Raster) Alpha() float64 {
	return 1.0 - r.Translucency
}

// NewRaster returns a new Image instance that is rendered dynamically using
// the specified generate function.
// Images returned from this method should draw dynamically to fill the width
// and height parameters passed to pixelColor.
func NewRaster(generate func(w, h int, r *Raster) image.Image) *Raster {
	return &Raster{Generator: generate}
}

// NewRasterWithPixels returns a new Image instance that is rendered dynamically
// by iterating over the specified pixelColor function for each x, y pixel.
// Images returned from this method should draw dynamically to fill the width
// and height parameters passed to pixelColor.
func NewRasterWithPixels(pixelColor func(x, y, w, h int) color.Color) *Raster {
	return &Raster{
		Generator: func(w, h int, r *Raster) image.Image {
			memdebug.Print(time.Now(), "in here with", r)
			if r == nil {
				memdebug.Print(time.Now(), "r is nil")
				debug.PrintStack()
				os.Exit(0)
			}
			if r.img == nil || r.img.Bounds() != r.img.Bounds() {
				// raster first pixel, figure out color type
				var dst draw.Image
				rect := image.Rect(0, 0, w, h)
				switch pixelColor(0, 0, w, h).(type) {
				case color.Alpha:
					dst = image.NewAlpha(rect)
				case color.Alpha16:
					dst = image.NewAlpha16(rect)
				case color.CMYK:
					dst = image.NewCMYK(rect)
				case color.Gray:
					dst = image.NewGray(rect)
				case color.Gray16:
					dst = image.NewGray16(rect)
				case color.NRGBA:
					dst = image.NewNRGBA(rect)
				case color.NRGBA64:
					dst = image.NewNRGBA64(rect)
				case color.RGBA:
					dst = image.NewRGBA(rect)
				case color.RGBA64:
					dst = image.NewRGBA64(rect)
				default:
					dst = image.NewRGBA(rect)
				}
				r.img = dst
				memdebug.Print(time.Now(), "set to a new img")
			}

			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					r.img.Set(x, y, pixelColor(x, y, w, h))
				}
			}

			return r.img
		},
	}
}

type subImg interface {
	SubImage(r image.Rectangle) image.Image
}

// NewRasterFromImage returns a new Raster instance that is rendered from the Go
// image.Image passed in.
// Rasters returned from this method will map pixel for pixel to the screen
// starting img.Bounds().Min pixels from the top left of the canvas object.
// Truncates rather than scales the image.
// If smaller than the target space, the image will be padded with zero-pixels to the target size.
func NewRasterFromImage(img image.Image) *Raster {
	return &Raster{
		Generator: func(w int, h int, r *Raster) image.Image {
			bounds := img.Bounds()

			rect := image.Rect(0, 0, w, h)

			switch {
			case w == bounds.Max.X && h == bounds.Max.Y:
				return img
			case w >= bounds.Max.X && h >= bounds.Max.Y:
				// try quickly truncating
				if sub, ok := img.(subImg); ok {
					return sub.SubImage(image.Rectangle{
						Min: bounds.Min,
						Max: image.Point{
							X: bounds.Min.X + w,
							Y: bounds.Min.Y + h,
						},
					})
				}
			default:
				if !rect.Overlaps(bounds) {
					return image.NewUniform(color.RGBA{})
				}
				bounds = bounds.Intersect(rect)
			}

			// respect the user's pixel format (if possible)
			if r.img == nil || r.img.Bounds() != img.Bounds() {
				var dst draw.Image
				switch i := img.(type) {
				case (*image.Alpha):
					dst = image.NewAlpha(rect)
				case (*image.Alpha16):
					dst = image.NewAlpha16(rect)
				case (*image.CMYK):
					dst = image.NewCMYK(rect)
				case (*image.Gray):
					dst = image.NewGray(rect)
				case (*image.Gray16):
					dst = image.NewGray16(rect)
				case (*image.NRGBA):
					dst = image.NewNRGBA(rect)
				case (*image.NRGBA64):
					dst = image.NewNRGBA64(rect)
				case (*image.Paletted):
					dst = image.NewPaletted(rect, i.Palette)
				case (*image.RGBA):
					dst = image.NewRGBA(rect)
				case (*image.RGBA64):
					dst = image.NewRGBA64(rect)
				default:
					dst = image.NewRGBA(rect)
				}
				r.img = dst
				memdebug.Print(time.Now(), "set r.img", r.img)
			} else {
				memdebug.Print(time.Now(), "re-use existing img", r.img)
			}

			draw.Draw(r.img, bounds, img, bounds.Min, draw.Over)
			return r.img
		},
	}
}
