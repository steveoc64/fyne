package gl

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg" // avoid users having to import when using image widget
	_ "image/png"  // avoid the same for PNG images
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/goki/freetype"
	"github.com/goki/freetype/truetype"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var textures = make(map[fyne.CanvasObject]uint32)
var refreshQueue = make(chan fyne.CanvasObject, 1024)

const vectorPad = 10

func getTexture(object fyne.CanvasObject, creator func(canvasObject fyne.CanvasObject) uint32) uint32 {

	img, skipCache := object.(*canvas.Image)
	if skipCache && img.PixelColor == nil {
		skipCache = false
	}

	if skipCache {
		return creator(object)
	}
	texture := textures[object]

	texture = creator(object)
	textures[object] = texture
	return texture
}

func newTexture() uint32 {
	var texture uint32

	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return texture
}

func (c *glCanvas) newGlCircleTexture(obj fyne.CanvasObject) uint32 {
	circle := obj.(*canvas.Circle)
	texture := newTexture()
	radius := fyne.Min(circle.Size().Width, circle.Size().Height) / 2

	width := scaleInt(c, circle.Size().Width+vectorPad*2)
	height := scaleInt(c, circle.Size().Height+vectorPad*2)
	stroke := circle.StrokeWidth * c.scale

	raw := image.NewRGBA(image.Rect(0, 0, width, height))
	scanner := rasterx.NewScannerGV(circle.Size().Width, circle.Size().Height, raw, raw.Bounds())

	if circle.FillColor != nil {
		filler := rasterx.NewFiller(width, height, scanner)
		filler.SetColor(circle.FillColor)
		rasterx.AddCircle(float64(width/2), float64(height/2), float64(scaleInt(c, radius)), filler)
		filler.Draw()
	}

	dasher := rasterx.NewDasher(width, height, scanner)
	dasher.SetColor(circle.StrokeColor)
	dasher.SetStroke(fixed.Int26_6(float64(stroke)*64), 0, nil, nil, nil, 0, nil, 0)
	rasterx.AddCircle(float64(width/2), float64(height/2), float64(scaleInt(c, radius)), dasher)
	dasher.Draw()

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(raw.Rect.Size().X), int32(raw.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(raw.Pix))

	return texture
}

func (c *glCanvas) newGlLineTexture(obj fyne.CanvasObject) uint32 {
	line := obj.(*canvas.Line)
	texture := newTexture()

	col := line.StrokeColor
	width := scaleInt(c, line.Size().Width+vectorPad*2)
	height := scaleInt(c, line.Size().Height+vectorPad*2)
	stroke := line.StrokeWidth * c.scale

	raw := image.NewRGBA(image.Rect(0, 0, width, height))
	scanner := rasterx.NewScannerGV(line.Size().Width, line.Size().Height, raw, raw.Bounds())
	dasher := rasterx.NewDasher(width, height, scanner)
	dasher.SetColor(col)
	dasher.SetStroke(fixed.Int26_6(float64(stroke)*64), 0, nil, nil, nil, 0, nil, 0)
	p1x, p1y := scaleInt(c, line.Position1.X-line.Position().X+vectorPad), scaleInt(c, line.Position1.Y-line.Position().Y+vectorPad)
	p2x, p2y := scaleInt(c, line.Position2.X-line.Position().X+vectorPad), scaleInt(c, line.Position2.Y-line.Position().Y+vectorPad)

	dasher.Start(rasterx.ToFixedP(float64(p1x), float64(p1y)))
	dasher.Line(rasterx.ToFixedP(float64(p2x), float64(p2y)))
	dasher.Stop(true)
	dasher.Draw()

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(raw.Rect.Size().X), int32(raw.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(raw.Pix))

	return texture
}

func (c *glCanvas) newGlRectTexture(rect fyne.CanvasObject) uint32 {
	texture := newTexture()

	col := theme.BackgroundColor()
	if wid, ok := rect.(fyne.Widget); ok {
		widCol := widget.Renderer(wid).BackgroundColor()
		if widCol != nil {
			col = widCol
		}
	} else if rect, ok := rect.(*canvas.Rectangle); ok {
		if rect.FillColor != nil {
			col = rect.FillColor
		}
	}

	r, g, b, a := col.RGBA()
	r8, g8, b8, a8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
	data := []uint8{r8, g8, b8, a8}
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, 1, 1, 0, gl.RGBA,
		gl.UNSIGNED_BYTE, gl.Ptr(data))

	return texture
}

func (c *glCanvas) newGlTextTexture(obj fyne.CanvasObject) uint32 {
	text := obj.(*canvas.Text)
	texture := newTexture()

	textScale := 1
	if runtime.GOOS == "darwin" {
		textScale = 2
	}
	bounds := text.MinSize()
	width := scaleInt(c, bounds.Width*textScale)
	height := scaleInt(c, bounds.Height*textScale)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var opts truetype.Options
	fontSize := float64(text.TextSize) * float64(c.Scale())
	opts.Size = fontSize
	opts.DPI = float64(textDPI * textScale)
	face := cachedFontFace(text.TextStyle, &opts)

	d := font.Drawer{}
	d.Dst = img
	d.Src = &image.Uniform{text.Color}
	d.Face = face
	d.Dot = freetype.Pt(0, height-face.Metrics().Descent.Ceil())
	d.DrawString(text.Text)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Rect.Size().X), int32(img.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))

	return texture
}

func renderGlImagePortion(point image.Point, width, height int,
	raw draw.Image, pixels image.Image, wg *sync.WaitGroup) {

	defer wg.Done()
	bounds := image.Rect(point.X, point.Y, point.X+width, point.Y+height)

	draw.Draw(raw, bounds, pixels, point, draw.Src)
}

func (c *glCanvas) newGlImageTexture(obj fyne.CanvasObject) uint32 {
	var raw *image.RGBA
	img := obj.(*canvas.Image)
	texture := newTexture()

	width := scaleInt(c, img.Size().Width)
	height := scaleInt(c, img.Size().Height)
	if width <= 0 || height <= 0 {
		return 0
	}

	if img.File != "" || img.Resource != nil {
		var file io.Reader
		var name string
		if img.Resource != nil {
			name = img.Resource.Name()
			file = bytes.NewReader(img.Resource.Content())
		} else {
			name = img.File
			handle, _ := os.Open(img.File)
			defer handle.Close()
			file = handle
		}

		if strings.ToLower(filepath.Ext(name)) == ".svg" {
			rasterMutex.Lock()
			defer rasterMutex.Unlock()
			info := rasters[img.Resource]
			if info == nil || info.w != width || info.h != height || info.alpha != img.Alpha() {
				icon, err := oksvg.ReadIconStream(file)
				if err != nil {
					log.Println("SVG Load error:", err)

					return 0
				}
				icon.SetTarget(0, 0, float64(width), float64(height))

				w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
				// this is used by our render code, so let's set it to the file aspect
				img.PixelAspect = float32(w) / float32(h)
				// if the image specifies it should be original size we need at least that many pixels on screen
				if img.FillMode == canvas.ImageFillOriginal {
					pixSize := fyne.NewSize(unscaleInt(c, w), unscaleInt(c, h))
					img.SetMinSize(pixSize)
				}

				raw = image.NewRGBA(image.Rect(0, 0, width, height))
				scanner := rasterx.NewScannerGV(w, h, raw, raw.Bounds())
				raster := rasterx.NewDasher(width, height, scanner)
				icon.Draw(raster, img.Alpha())

				if cacheDuration > 0 {
					rasters[img.Resource] = &rasterInfo{
						pix:     raw,
						w:       width,
						h:       height,
						alpha:   img.Alpha(),
						expires: time.Now().Add(cacheDuration),
					}
				}
			} else {
				raw = info.pix
				info.expires = time.Now().Add(cacheDuration)
			}
		} else {
			pixels, _, err := image.Decode(file)

			if err != nil {
				log.Println("image err", err)

				return 0
			}
			origSize := pixels.Bounds().Size()
			// this is used by our render code, so let's set it to the file aspect
			img.PixelAspect = float32(origSize.X) / float32(origSize.Y)
			// if the image specifies it should be original size we need at least that many pixels on screen
			if img.FillMode == canvas.ImageFillOriginal {
				pixSize := fyne.NewSize(unscaleInt(c, origSize.X), unscaleInt(c, origSize.Y))
				img.SetMinSize(pixSize)
			}

			raw = image.NewRGBA(pixels.Bounds())
			draw.Draw(raw, pixels.Bounds(), pixels, image.ZP, draw.Src)
		}
	} else if img.PixelColor != nil {
		raw = image.NewRGBA(image.Rect(0, 0, width, height))
		pixels := newPixelImage(img, c.Scale())

		halfWidth := raw.Bounds().Size().X / 2
		halfHeight := raw.Bounds().Size().Y / 2

		// use a WaitGroup so we don't return our image until it's complete
		var wg sync.WaitGroup
		wg.Add(4)

		go renderGlImagePortion(image.ZP, halfWidth, halfHeight, raw, pixels, &wg)
		go renderGlImagePortion(image.Pt(0, halfHeight), halfWidth, height-halfHeight, raw, pixels, &wg)
		go renderGlImagePortion(image.Pt(halfWidth, 0), width-halfWidth, halfHeight, raw, pixels, &wg)
		go renderGlImagePortion(image.Pt(halfWidth, halfHeight), width-halfWidth, height-halfHeight, raw, pixels, &wg)

		wg.Wait()
	} else {
		raw = image.NewRGBA(image.Rect(0, 0, 1, 1))
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(raw.Rect.Size().X), int32(raw.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(raw.Pix))

	return texture
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		info := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(info))

		return 0, fmt.Errorf("failed to compile %v: %v", source, info)
	}

	return shader, nil
}

const (
	vertexShaderSource = `
    #version 150
    in vec3 vert;
    in vec2 vertTexCoord;
    out vec2 fragTexCoord;

    void main() {
        fragTexCoord = vertTexCoord;

        gl_Position = vec4(vert, 1);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 150
    uniform sampler2D tex;

    in vec2 fragTexCoord;
    out vec4 frag_colour;
    
    void main() {
        vec4 color = texture(tex, fragTexCoord);
        if(color.a < 0.001)
            discard;

        frag_colour = color;
    }
` + "\x00"
)

func (c *glCanvas) initOpenGL() {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	c.program = prog
}
