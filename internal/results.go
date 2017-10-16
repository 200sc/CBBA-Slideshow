package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
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
	sslides[i+1].Append(show.ImageAt("cbbaResult1.PNG", .33, .5, mod.Scale(1.5, 1.5)))
	sslides[i+1].Append(show.ImageAt("cbbaResult2.PNG", .66, .5, mod.Scale(1.5, 1.5)))
	//
	// As a sort of aside, the paper notes that, because agents don't actually
	// share information with each other about the environment, some conclusions
	// they reach about the value of their taking on tasks can be innaccurate.
	// This is seen in the form of SA Error, but this error doesn't stop the
	// agents from converging or performing well.
	//
	// These figures show results for things entirely outside of CBBA which
	// revolve around how networks deal with inconsistent information across
	// the network. A random spanning tree with additional random links was
	// used to simulate network configurations.
	//
	sslides[i+2].Append(show.Header("Aside: Prim Allocation"))
	ListSlide(sslides[i+2], "Aside: Prim Allocation",
		"- Introduced in a 2004 paper with some credit to Koenig (monday's paper)",
		"- Simple, approximate multi-assignment algorithm",
		"- Makes no promises in terms of iterations",
		"- Claims to be computationally easy",
	)
	// (read points)
	// Just to emphasize here-- the 2004 papre on Prim Allocation doesn't say that
	// it only takes n iterations to do its job, rather the optimizations its concerned
	// with are the total number of bids made.
	// http://idm-lab.org/bib/abstracts/papers/iros04a.pdf
	sslides[i+3].Append(show.Header("CBBA vs Prim Allocation"))
	sslides[i+3].Append(show.ImageAt("cbbaResult3.PNG", .33, .5, mod.Scale(1.5, 1.5)))
	sslides[i+3].Append(show.ImageAt("cbbaResult4.PNG", .66, .5, mod.Scale(1.5, 1.5)))
	// So I want to point out in the time step graph here, and the paper is up
	// front with this, that Prim is being given a rough deal here- Prim will only
	// assign one task per iteration, and each iteration is designed to be very
	// efficient. This graph, however, makes it look like Prim is outrageously
	// slower than the methods proposed in this paper when really the way they
	// are being judged explicitly benefits these methods without judging a
	// useful characteristic, but with the implication that these new algorithms
	// are faster.
	// (consider going back to inconsistent slide and pointing out they're also
	//  using this very useless metric)
}
