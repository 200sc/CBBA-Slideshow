package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	sincePaper = SlideSetup{
		addSincePaper,
		2,
	}
)

func addSincePaper(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Since This Paper"))
	ListSlide(sslides[i+1], "CBBA Variants",
		"- Asynchronous CBBA",
		"- Coupled-Constraint CBBA",
		"- CBBA with Relays",
		"- CBBA with Information-rich Rapidly-exploring Random Trees (IRRT)",
	)
	// http://acl.mit.edu/projects/cbba.html
	// https://www.youtube.com/watch?v=-PBJMtMiFRA
}
