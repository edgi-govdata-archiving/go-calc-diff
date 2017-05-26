package main

import (
	"github.com/qri-io/go-diff/diffmatchpatch"
	"net/http"
)

// NewDifferFromRequest can later be expanded to support
// configuring the differ with request params
func NewDifferFromRequest(r *http.Request) *diffmatchpatch.DiffMatchPatch {
	dmp := diffmatchpatch.New()

	// Number of seconds to map a diff before giving up (0 for infinity).
	// dmp.DiffTimeout = time.Second

	// Cost of an empty edit operation in terms of edit characters.
	// default: 4
	if dec, err := reqParamInt("diffEditCost", r); dec != 0 && err == nil {
		dmp.DiffEditCost = dec
	}

	// At what point is no match declared (0.0 = perfection, 1.0 = very loose).
	// default: 0.5
	if mt, err := reqParamFloat("matchThreshold", r); mt != 0 && err == nil {
		dmp.MatchThreshold = mt
	}

	// How far to search for a match (0 = exact location, 1000+ = broad match).
	// A match this many characters away from the expected location will add
	// 1.0 to the score (0.0 is a perfect match).
	// default: 1000
	if md, err := reqParamInt("matchDistance", r); md != 0 && err == nil {
		dmp.MatchDistance = md
	}

	// When deleting a large block of text (over ~64 characters), how close do
	// the contents have to be to match the expected contents. (0.0 = perfection,
	// 1.0 = very loose).  Note that MatchThreshold controls how closely the
	// end points of a delete need to match.
	// default: 0.5
	if pdt, err := reqParamFloat("patchDeleteThreshold", r); pdt != 0 && err == nil {
		dmp.PatchDeleteThreshold = pdt
	}

	// Chunk size for context length.
	if pm, err := reqParamInt("patchMargin", r); pm != 0 && err == nil {
		dmp.PatchMargin = pm
	}

	// The number of bits in an int.
	// dmp.MatchMaxBits=         32
	return dmp
}

// CalcDiff is just a wrapper around dmp.DiffMain
func CalcDiff(dmp *diffmatchpatch.DiffMatchPatch, text1, text2 string) []diffmatchpatch.Diff {
	return dmp.DiffMain(text1, text2, false)
}
