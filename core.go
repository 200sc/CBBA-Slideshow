package main

import (
	"fmt"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/render/mod"
	"github.com/oakmound/oak/shape"
	"golang.org/x/image/colornames"
)

const (
	width  = 1920
	height = 1080
)

var (
	Express28  = show.FontSize(28)(show.Express)
	Gnuolane28 = show.FontSize(28)(show.Gnuolane)
	Libel28    = show.FontSize(28)(show.Libel)

	RLibel28 = show.FontColor(colornames.Blue)(Libel28)

	Express44  = show.FontSize(44)(show.Express)
	Gnuolane44 = show.FontSize(44)(show.Gnuolane)
	Libel44    = show.FontSize(44)(show.Libel)

	Express72  = show.FontSize(72)(show.Express)
	Gnuolane72 = show.FontSize(72)(show.Gnuolane)
	Libel72    = show.FontSize(72)(show.Libel)
)

func main() {

	show.SetDims(width, height)
	show.SetTitleFont(Gnuolane72)

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

	setups := []slideSetup{
		intro,
		authors,
	}

	total := 0

	for _, setup := range setups {
		total += setup.len
	}

	fmt.Println("Total slides", total)

	oak.LoadingR = bkg

	sslides := static.NewSlideSet(total,
		static.Background(bkg),
	)

	nextStart := 0

	for _, setup := range setups {
		setup.add(nextStart, sslides)
		nextStart += setup.len
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
	show.AddNumberShortcuts(len(slides))
	show.Start(slides...)
}

type slideSetup struct {
	add func(int, []*static.Slide)
	len int
}

var (
	intro = slideSetup{
		addIntro,
		2,
	}
	authors = slideSetup{
		addAuthors,
		3,
	}
)

func addIntro(i int, sslides []*static.Slide) {

	sslides[i].Append(
		show.Title("Consensus-Based Decentralized Auctions for Robust Task Allocation"),
		show.TxtAt(Gnuolane44, "Han-Lim Choi, Luc Brunet, and Jonathan P. How", .5, .5),
		show.TxtAt(Gnuolane44, "Presentation by Patrick Stephen", .5, .8),
	)
	sslides[i+1].Append(show.Header("Overview"))
	sslides[i+1].Append(
		show.TxtSetFrom(Gnuolane44, .25, .35, 0, .07,
			"- Authors",
			"- Problem History",
			"- The Paper",
			"- Happenings Since 2009",
		)...,
	)
}

func addAuthors(i int, sslides []*static.Slide) {

	imgpos := .15

	sslides[i].Append(
		show.Header("Authors: Han-Lim Choi"),
		show.Image("hanlim.jpg", imgpos, .35).Modify(mod.Scale(.4, .4)))
	sslides[i].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Asso. Professor of Aerospace Engineering at KAIST",
			"- PhD from MIT in 2009",
			"- Director of the Laboratory for Information and Control Systems",
			"- Research work with sensors, targetting",
			"- This paper was his first work with auctions, consensus algorithms*",
			"- Shifted towards decentralized task assignment work after this paper",
		)...,
	)
	sslides[i+1].Append(
		show.Header("Authors: Luc Brunet"),
		show.Image("lbrunet.jpg", imgpos, .35),
	)
	sslides[i+1].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Master's Degree from MIT in Aerospace Engineering",
			"- Just worked on this paper",
			"- Since founded Provectus Robotics Solutions",
		)...,
	)
	sslides[i+2].Append(
		show.Header("Authors:  Jonathan P. How"),
		show.Image("s_jon2.gif", imgpos, .35),
	)
	sslides[i+2].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Professor of Aeronautics and Astronautics at MIT",
			"- Received the Automatica Applications Prize in 2011 with Han-Lim Choi",
			"- Continued work on CBBA algorithm through 2011 (at least)",
		)...,
	)
}
