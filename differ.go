package main

import (
	"github.com/qri-io/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

func CalcDiff(text1, text2 string) []diffmatchpatch.Diff {
	return dmp.DiffMain(text1, text2, false)
}
