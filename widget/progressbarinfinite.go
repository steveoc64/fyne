package widget

import (
	"image/color"
	"sync/atomic"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/internal/cache"
	"fyne.io/fyne/theme"
)

const (
	infiniteRefreshRate              = 50 * time.Millisecond
	maxProgressBarInfiniteWidthRatio = 1.0 / 5
	minProgressBarInfiniteWidthRatio = 1.0 / 20
	progressBarInfiniteStepSizeRatio = 1.0 / 50
)

type infProgressRenderer struct {
	objects []fyne.CanvasObject
	bar     *canvas.Rectangle
	ticker  *time.Ticker
	running atomic.Value

	progress *ProgressBarInfinite
}

// MinSize calculates the minimum size of a progress bar.
func (p *infProgressRenderer) MinSize() fyne.Size {
	// this is to create the same size infinite progress bar as regular progress bar
	text := textMinSize("100%", theme.TextSize(), fyne.TextStyle{})

	return fyne.NewSize(text.Width+theme.Padding()*4, text.Height+theme.Padding()*2)
}

func (p *infProgressRenderer) updateBar() {
	progressSize := p.progress.Size()
	barWidth := p.bar.Size().Width
	barPos := p.bar.Position()

	maxBarWidth := int(float64(progressSize.Width) * maxProgressBarInfiniteWidthRatio)
	minBarWidth := int(float64(progressSize.Width) * minProgressBarInfiniteWidthRatio)
	stepSize := int(float64(progressSize.Width) * progressBarInfiniteStepSizeRatio)

	// check to make sure inner bar is sized correctly
	// if bar is on the first half of the progress bar, grow it up to maxProgressBarInfiniteWidthPercent
	// if on the second half of the progress bar width, shrink it down until it gets to minProgressBarInfiniteWidthPercent
	if barWidth < maxBarWidth && barPos.X+barWidth < progressSize.Width/2 {
		// slightly increase width
		newBoxSize := fyne.Size{Width: barWidth + stepSize, Height: progressSize.Height}
		p.bar.Resize(newBoxSize)
	} else {
		barPos.X += stepSize

		if barWidth <= minBarWidth {
			// loop around to start when bar goes to end
			barPos.X = 0
			stepSize = 0
			newBoxSize := fyne.Size{Width: minBarWidth, Height: progressSize.Height}
			p.bar.Resize(newBoxSize)
		} else if barPos.X+barWidth > progressSize.Width {
			// crop to the end of the bar
			barWidth = progressSize.Width - barPos.X
			newBoxSize := fyne.Size{Width: barWidth, Height: progressSize.Height}
			p.bar.Resize(newBoxSize)
		}
	}

	p.bar.Move(barPos)
}

// Layout the components of the infinite progress bar
func (p *infProgressRenderer) Layout(size fyne.Size) {
	p.updateBar()
}

// applyTheme updates the infinite progress bar to match the current theme
func (p *infProgressRenderer) applyTheme() {
}

func (p *infProgressRenderer) BackgroundColor() color.Color {
	return theme.ButtonColor()
}

// Refresh updates the size and position of the horizontal scrolling infinite progress bar
func (p *infProgressRenderer) Refresh() {
	p.bar.FillColor = theme.PrimaryColor()

	p.updateBar()
	canvas.Refresh(p.progress)
}

func (p *infProgressRenderer) Objects() []fyne.CanvasObject {
	return p.objects
}

// Start the infinite progress bar background thread to update it continuously
func (p *infProgressRenderer) start() {
	if !p.running.Load().(bool) {
		p.ticker = time.NewTicker(infiniteRefreshRate)
		p.running.Store(true)

		go p.infiniteProgressLoop()
	}
}

// Stop the infinite progress goroutine and sets value to the Max
func (p *infProgressRenderer) stop() {
	p.running.Store(false)
}

// infiniteProgressLoop should be called as a goroutine to update the inner infinite progress bar
// the function can be exited by calling Stop()
func (p *infProgressRenderer) infiniteProgressLoop() {
	for p.running.Load().(bool) {
		select {
		case <-p.ticker.C:
			p.Refresh()
			break
		}
	}
	if p.ticker != nil {
		p.ticker.Stop()
	}
}

func (p *infProgressRenderer) Destroy() {
	p.stop()
}

// ProgressBarInfinite widget creates a horizontal panel that indicates waiting indefinitely
// An infinite progress bar loops 0% -> 100% repeatedly until Stop() is called
type ProgressBarInfinite struct {
	BaseWidget
}

// Show this widget, if it was previously hidden
func (p *ProgressBarInfinite) Show() {
	p.Start()
	p.BaseWidget.Show()
}

// Hide this widget, if it was previously visible
func (p *ProgressBarInfinite) Hide() {
	p.Stop()
	p.BaseWidget.Hide()
}

// Start the infinite progress bar background thread to update it continuously
func (p *ProgressBarInfinite) Start() {
	cache.Renderer(p).(*infProgressRenderer).start()
}

// Stop the infinite progress goroutine and sets value to the Max
func (p *ProgressBarInfinite) Stop() {
	cache.Renderer(p).(*infProgressRenderer).stop()
}

// Running returns the current state of the infinite progress animation
func (p *ProgressBarInfinite) Running() bool {
	if !cache.IsRendered(p) {
		return false
	}

	return cache.Renderer(p).(*infProgressRenderer).running.Load().(bool)
}

// MinSize returns the size that this widget should not shrink below
func (p *ProgressBarInfinite) MinSize() fyne.Size {
	p.ExtendBaseWidget(p)
	return p.BaseWidget.MinSize()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (p *ProgressBarInfinite) CreateRenderer() fyne.WidgetRenderer {
	p.ExtendBaseWidget(p)
	bar := canvas.NewRectangle(theme.PrimaryColor())
	render := &infProgressRenderer{
		objects:  []fyne.CanvasObject{bar},
		bar:      bar,
		progress: p,
	}
	render.running.Store(false)
	render.start()
	return render
}

// NewProgressBarInfinite creates a new progress bar widget that loops indefinitely from 0% -> 100%
// SetValue() is not defined for infinite progress bar
// To stop the looping progress and set the progress bar to 100%, call ProgressBarInfinite.Stop()
func NewProgressBarInfinite() *ProgressBarInfinite {
	p := &ProgressBarInfinite{}
	Renderer(p).Layout(p.MinSize())
	return p
}

// TODO - Bind the ProgressBarInfinite to a float ?
