package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
)

var (
	mainslide = SlideSetup{
		addMain,
		8,
	}
)

func addMain(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("This Paper"))
	//
	// This paper has a lot of detailed background information
	// that goes into consensus algorithms, auctions, and task
	// assignment, but we should already have that background
	// info from the last papers we've gone through. I do have slides
	// for this info but I'll skip them unless someone wants a refresher.
	//
	ListSlide(sslides[i+1], "The New Algorithms",
		"- Consensus-Based Auction Algorithm (CBAA)",
		"    Combine Auctions and Consensus to guarantee convergence",
		"    and to perform efficiently.",
		"- Consensus-Based Bundle Algorithm (CBBA)",
		"    Expand CBAA to Multi-Assignment problems through generalization",
	)
	sslides[i+2].Append(show.Header("CBAA: Phase 1"))
	sslides[i+2].Append(show.ImageAt("cbaa1.PNG", .5, .5, mod.Scale(2.0, 2.0)))
	sslides[i+3].Append(show.Header("CBAA: Phase 2"))
	sslides[i+3].Append(show.ImageAt("cbaa2.PNG", .5, .5, mod.Scale(2.0, 2.0)))
	//
	sslides[i+4].Append(show.Header("CBBA: Phase 1"))
	sslides[i+4].Append(show.ImageAt("cbba1.PNG", .5, .5, mod.Scale(2.0, 2.0)))
	sslides[i+5].Append(show.Header("CBBA: Phase 2"))
	sslides[i+5].Append(show.ImageAt("cbbaTable.PNG", .4, .54, mod.Scale(1.2, 1.2)))
	sslides[i+5].Append(show.TxtSetFrom(Gnuolane44, .7, .3, 0, .07,
		"1) Update: y_ij = y_kj, z_ij = z_kj",
		"2) Reset: y_ij = 0, z_ij = {}",
		"3) Leave: y_ij = y_ij, z_ij = z_ij",
	)...)
	//
	sslides[i+6].Append(show.Header("Scoring: Diminishing Marginal Gain"))
	ListSlide(sslides[i+6], "Scoring: Diminishing Marginal Gain",
		"'The value of a task cannot increase as other elements",
		"  are added to the set before it.'",
	)
	// AKA: the value of being assigned to a task must be the same or must
	// decrease when new tasks are appended to the task list
	// (( draw something? ))
	ListSlide(sslides[i+7], "Scoring: Time-Discounted Reward",
		"'In a time-sensitive target assignment problem, the ",
		"  time-discounted reward for a target decreases as the",
		"  combat vehicle visits another target first.'",
	)
}
