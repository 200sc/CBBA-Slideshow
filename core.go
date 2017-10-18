package main

import (
	"fmt"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/shape"
	"golang.org/x/image/colornames"

	"./internal"
)

const (
	width  = 1920
	height = 1080
)

func main() {

	show.SetDims(width, height)
	show.SetTitleFont(internal.Gnuolane72)

	bz1, _ := shape.BezierCurve(
		width/15, height/5,
		width/15, height/15,
		width/5, height/15)

	bz2, _ := shape.BezierCurve(
		width-(width/15), height/5,
		width-(width/15), height/15,
		width-(width/5), height/15)

	bz3, _ := shape.BezierCurve(
		width/15, height-(height/5),
		width/15, height-(height/15),
		width/5, height-(height/15))

	bz4, _ := shape.BezierCurve(
		width-(width/15), height-(height/5),
		width-(width/15), height-(height/15),
		width-(width/5), height-(height/15))

	bkg := render.NewComposite(
		render.NewColorBox(width, height, colornames.Slateblue),
		render.BezierThickLine(bz1, colornames.White, 1),
		render.BezierThickLine(bz2, colornames.White, 1),
		render.BezierThickLine(bz3, colornames.White, 1),
		render.BezierThickLine(bz4, colornames.White, 1),
	)

	setups := internal.Setups

	total := 0

	for _, setup := range setups {
		total += setup.Len
	}

	fmt.Println("Total slides", total)

	oak.LoadingR = bkg

	sslides := static.NewSlideSet(total,
		static.Background(bkg),
	)

	nextStart := 0

	for _, setup := range setups {
		setup.Add(nextStart, sslides)
		nextStart += setup.Len
	}

	oak.SetupConfig.Screen = oak.Screen{
		Width:  width,
		Height: height,
	}
	oak.SetupConfig.FrameRate = 30
	oak.SetupConfig.DrawFrameRate = 30

	slides := make([]show.Slide, len(sslides))
	for i, s := range sslides {
		slides[i] = s
	}
	shotIndex := 0
	oak.AddCommand("shot", func([]string) {
		rgba := oak.ScreenShot()
		f, err := os.Create("shot" + strconv.Itoa(shotIndex) + ".png")
		if err != nil {
			dlog.Error(err)
			return
		}
		png.Encode(f, rgba)
		dlog.ErrorCheck(f.Close())
		shotIndex++
	})
	oak.AddCommand("allShots", func([]string) {
		for i := 0; i <= len(slides); i++ {
			rgba := oak.ScreenShot()
			f, err := os.Create("shot" + strconv.Itoa(shotIndex) + ".png")
			if err != nil {
				dlog.Error(err)
				return
			}
			png.Encode(f, rgba)
			dlog.ErrorCheck(f.Close())
			shotIndex++
			time.Sleep(100 * time.Millisecond)
			event.Trigger("KeyUpRightArrow", nil)
		}
	})
	show.AddNumberShortcuts(len(slides))
	show.Start(slides...)
}
