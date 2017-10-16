package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	history = SlideSetup{
		addHistory,
		3,
	}
)

func addHistory(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("History"))
	ListSlide(sslides[i+1], "History of Fleet Cooperation: in the paper",
		"- 2002-4: Centralized Planning",
		// This has troubles because it requires that the robots can all
		// connect to some single point.
		"- 2003-5: Multi-Instance Centralized Planning",
		// A solution to the connection problem gave each robot the capability
		// to act as a centralized planner, so the robots just needed to connect
		// to said central robot. It also hypothetically allows other robots to
		// take over if that said central robot fails.
		// These were also troubled due to unrealistic expectations of effective
		// connectivity and bandwidth.
		"- 2004-7: Consensus Algorithms",
		// These concerns led to the use of consensus algorithms, which we've gone
		// over with the fleets of vibrating shape-forming robots, but which
		// don't require any centralized decider and thus can work with weaker
		// network connectivity.
		"- 2002-7: Auction Algorithms",
		// Auction algorithms, which we've gone over a lot recently,
		// were being applied to multi-robot systems at the same time, to good
		// success, but they used centralized planning of some kind, or avoided
		// global planning all together.
		"- 2005-7: Sequential Auction Algorithms",
		// That brings us to the sequential auction algorithms that we discussed
		// on Monday. Attempts to combine this with bundling required enumeration
		// over all possible bundles, however.
	)
	ListSlide(sslides[i+2], "History of Fleet Cooperation: not in the paper",
		"- 1997-8: MARTHA",
		// MARTHA focused on letting a fleet of robots operate individually and
		// avoid coordinate to avoid collisions in their pathing.
		// There's related research around this time where things like establishing
		// traffic rules for the robots were considered. MARTHA is a rather
		// general process that targets collisions as trajectory-based or
		// node-based.
		// http://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=619307
		// http://ieeexplore.ieee.org/stamp/stamp.jsp?tp=&arnumber=667325
	)
}
