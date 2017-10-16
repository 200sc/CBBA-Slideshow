package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
)

var (
	proofs = SlideSetup{
		addProofs,
		7,
	}
)

func addProofs(i int, sslides []*static.Slide) {
	sslides[i].Append(show.Title("Proofs"))
	// The paper goes into a few proofs, all of which culminate
	// in concluding convergence and performance characteristics
	// of CBAA and CBBA.
	ListSlide(sslides[i+1], "Convergence",
		"- No Agent has more tasks assigned than allowed per agent.",
		"- No Task is assigned to more than one agent.",
		"- The total number of assignments is the lesser of the number",
		"    of tasks and the maximum number of possible assignments.",
		"- Every agent is either assigned to or not assigned to each task.",
	)
	// Convergence here has a very formal definition
	// described here in words instead of notation.
	sslides[i+2].Append(show.Header("Sequential Greedy Algorithm"))
	sslides[i+2].Append(show.ImageAt("sqgreed.PNG", .5, .5, mod.Scale(2.0, 2.0)))
	// The Sequential Greedy algorithm or SGA is introduced and proven
	// to perform equvialently to the CBAA in order to make the proofs
	// about the CBAA less complex.
	sslides[i+3].Append(show.Header("Static Network"))
	sslides[i+4].Append(show.Header("Dynamic Network and Asynchronous Conflict Resolution"))
	sslides[i+5].Append(show.Header("Inconsistent Information"))
	ListSlide(sslides[i+6], "Performance",
		"- SOPT <= 2 * CBAA",
		"- MOPT <= 2 * CBBA",
	)
	// To analyze the performance of the new algorithms, the paper
	// is ultimately mostly concerned with how optimal the solutions
	// produced are. For this it defines SOPT and MOPT, or the value
	// of the optimal assignment for a single or multiple assignment
	// problem.
	// sslides[i+6].Append(show.Header("SOPT vs CBAA"))
	// The premise behind the proof that CBAA is at most twice as bad as SOPT
	// is that, if we take the worst assignment that CBAA could possibly produce,
	// and perform swaps on that assignment until it is the optimal assignment,
	// how much utility do we gain? Complexity is avoided by using the SGA
	// instead of the CBAA for this task.
	// sslides[i+7].Append(show.Header("MOPT vs CBBA"))
	// The short "proof" of CBBA's efficiency is that you can treat it just
	// like CBAA vs single assignment.
}
