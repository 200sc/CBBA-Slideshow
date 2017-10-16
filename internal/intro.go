package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	intro = SlideSetup{
		addIntro,
		2,
	}
	// https://www.youtube.com/watch?v=FhmRaRd8G24
	// https://www.youtube.com/watch?v=M_TfyDl1Uik
)

func addIntro(i int, sslides []*static.Slide) {

	sslides[i].Append(
		show.Title("Consensus-Based Decentralized Auctions for Robust Task Allocation"),
		show.TxtAt(Gnuolane44, "Han-Lim Choi, Luc Brunet, and Jonathan P. How", .5, .5),
		show.TxtAt(Gnuolane44, "Presentation by Patrick Stephen", .5, .8),
	)
	ListSlide(sslides[i+1], "Overview",
		"- Authors",
		"- Problem History",
		"- This Paper's Methods",
		"- Happenings Since 2009",
	)

}
