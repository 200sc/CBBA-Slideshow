package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
)

var (
	results = SlideSetup{
		addResults,
		5,
	}
)

func addResults(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Results"))
	//
	// The paper then gets into its experimental results.
	// Nothing the paper did uses real robots, it was just simulations.
	//
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
	// used to simulate network configurations, and random positional error
	// was given to each robot for detecting its task locations.
	//
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
	//
	// ICBAA - iterative CBAA, each agent has the same number of task
	sslides[i+4].Append(show.Header("Previous CBBA Results"))
	sslides[i+4].Append(show.ImageAt("cbbaResult5.PNG", .33, .39, mod.Scale(.8, .8)))
	sslides[i+4].Append(show.ImageAt("cbbaResult6.PNG", .66, .71, mod.Scale(.8, .8)))
	//
	// It's perhaps worth mentioning that this paper is almost identical to another
	// paper that was written for a conference a year before this was released,
	// except that the conference paper had completely different results and compared
	// against different algorithms. This gives the competing algorithms a much
	// more even-looking comparison with CBBA, and might more closely indicate
	// CBBA's performance. This is hypothetical, however, as many things could
	// have changed in between the two papers' releases.
	//
	// ETSP - Euclidean Travelling Salesperson problem
	//        referred to as a decentralized auction algorithm by this paper
	// GBAA - Greedy Auction Algorithm
}
