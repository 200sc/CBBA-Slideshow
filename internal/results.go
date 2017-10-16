package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	results = SlideSetup{
		addResults,
		4,
	}
)

func addResults(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Results"))
	sslides[i+1].Append(show.Header("Results With Inconsistent Information"))
	sslides[i+2].Append(show.Header("Aside: Prim Allocation"))
	sslides[i+3].Append(show.Header("CBBA vs Prim Allocation"))
}
