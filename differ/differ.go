package differ

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Diff(origin, other string) string {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(origin, other, true)

	return dmp.DiffPrettyText(diffs)
}
