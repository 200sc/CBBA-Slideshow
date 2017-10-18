package internal

import (
	"github.com/oakmound/oak/examples/slide/show"
	"github.com/oakmound/oak/examples/slide/show/static"
	"github.com/oakmound/oak/render/mod"
)

var (
	proofs = SlideSetup{
		addProofs,
		6,
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
	// to perform equvialently to the CBBA in order to make the proofs
	// about the CBBA less complex.
	ListSlide(sslides[i+3], "SGA = CBAA",
		"Given CBBA resolves conflicts synchronously with DMG:",
		"1) Entries in agent bundles match to entries in SGA bundles",
		"     at the same iteration.",
		"2) No agent will outbid an agent assigned a task that SGA would assign",
		"     in following rounds.",
		"3) Assignments that SGA would assign won't be unassigned.",
		"4) After information has propagated to all agents, all agents",
		"     will agree on assignments SGA would assign.",
	)

	// A number of things can be proven via modeling a fleet's communication
	// as an undirected graph

	ListSlide(sslides[i+4], "Dynamic Network and Asynchronous Conflict Resolution",
		"'[I]f it is ensured that an agent eventually communicates with",
		"    its neighbor, then the CBBA process converges in finite time,",
		"    even in the case when asynchronous conflict resolution is allowed.'",
	)

	// But the static networks proof revolved around sequential conflict resolution
	// to make sure each agent was in lock-step with one another. The paper presents
	// a short argument that you can treat asynchronous conflict resolution as
	// synchronous in a static network where, occasionally, you remove links
	// from the network to represent a busy or blocked agent. Then so long as
	// all links get added back in in finite time, everything eventually converges
	// all the same.

	ListSlide(sslides[i+5], "Performance",
		"- SOPT: Optimal Single Assignment",
		"- SOPT <= 2 * CBAA",
		"",
		"- MOPT: Optimal Multiple Assignment",
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
	//
	// (write on board:)
	// c_ii >= c_jj if i < j
	//
	// CBAA / SGA = sum over all c_ii
	//
	// c_ii >= c_ij for all i,j s.t. j > i
	// c_ii >= c_ji for all i,j s.t. j > i
	//
	// score after a swap becomes c_ij + c_ji
	// score was c_ii + c_jj
	//
	// because of (c_ii >= c_ij), c_ij + c_ji <= c_ii + c_ii or 2c_ii
	// when c_ij = c_ji = c_ii
	//
	// sslides[i+7].Append(show.Header("MOPT vs CBBA"))
	// The short "proof" of CBBA's efficiency is that you can treat it just
	// like CBAA vs single assignment.
}
