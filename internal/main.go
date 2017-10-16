package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	mainslide = SlideSetup{
		addMain,
		7,
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
	sslides[i+1].Append(show.Header("CBAA: Phase 1"))
	sslides[i+2].Append(show.Header("CBAA: Phase 2"))
	//
	sslides[i+3].Append(show.Header("CBBA: Phase 1"))
	sslides[i+4].Append(show.Header("CBBA: Phase 2"))
	//
	sslides[i+5].Append(show.Header("Scoring: Diminishing Marginal Gain"))
	sslides[i+6].Append(show.Header("Scoring: Time-Discounted Reward"))
	//
}