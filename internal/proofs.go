package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
)

var (
	proofs = SlideSetup{
		addProofs,
		9,
	}
)

func addProofs(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Proofs"))
	sslides[i+1].Append(show.Title("Convergence"))
	sslides[i+2].Append(show.Header("Sequential Greedy Algorithm"))
	sslides[i+3].Append(show.Header("Static Network"))
	sslides[i+4].Append(show.Header("Dynamic Network and Asynchronous Conflict Resolution"))
	sslides[i+5].Append(show.Header("Inconsistent Information"))
	sslides[i+6].Append(show.Title("Performance"))
	sslides[i+7].Append(show.Header("SOPT vs CBAA"))
	sslides[i+8].Append(show.Header("MOPT vs CBBA"))
}
