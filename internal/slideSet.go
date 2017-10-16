package internal

import "github.com/oakmound/oak/examples/slide/show/static"

var (
	Setups = []SlideSetup{
		intro,
		authors,
		history,
		mainslide,
		proofs,
		results,
		sincePaper,
	}
)

type SlideSetup struct {
	Add func(int, []*static.Slide)
	Len int
}
