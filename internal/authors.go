package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
)

var (
	authors = SlideSetup{
		addAuthors,
		4,
	}
)

func addAuthors(i int, sslides []*static.Slide) {

	imgpos := .15

	sslides[i].Append(show.Title("Authors"))

	sslides[i+1].Append(
		show.Header("Authors: Han-Lim Choi"),
		show.Image("hanlim.jpg", imgpos, .35).Modify(mod.Scale(.4, .4)))
	sslides[i+1].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Associate Professor of Aerospace Engineering at KAIST",
			"- PhD from MIT in 2009",
			"- Was Director of the Laboratory for Information and Control Systems", // MIT
			"- Research work with sensors, targetting",
			"- This paper was his first work with auctions, consensus algorithms",
			"- Shifted towards decentralized task assignment work after this paper",
		)...,
	)
	// Korea Advanced Institute of Science and Technology
	sslides[i+2].Append(
		show.Header("Authors: Luc Brunet"),
		show.Image("lbrunet.jpg", imgpos, .35),
	)
	sslides[i+2].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Master's Degree from MIT in Aerospace Engineering",
			"- More or less only worked on this paper",
			"- Since founded Provectus Robotics Solutions",
		)...,
	)
	sslides[i+3].Append(
		show.Header("Authors:  Jonathan P. How"),
		show.Image("s_jon2.gif", imgpos, .35),
	)
	sslides[i+3].Append(
		show.TxtSetFrom(Gnuolane44, imgpos+.2, .35, 0, .07,
			"- Professor of Aeronautics and Astronautics at MIT",
			"- Received the Automatica Applications Prize in 2011 with Han-Lim Choi",
			"- Continued work on CBBA algorithm through 2011 (at least)",
		)...,
	)
	// Prize had nothing to do with assignment problems
	// was for http://www.sciencedirect.com/science/article/pii/S0005109810002104
	// Continuous trajectory planning of mobile sensors for informative forecasting
}
