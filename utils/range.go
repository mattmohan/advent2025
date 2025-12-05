package utils

import (
	"strconv"
)

type Range struct {
	Start int64
	End   int64
}

func (r Range) InRange(x int64) bool {
	return x >= r.Start && x <= r.End
}
func (r Range) Overlaps(other Range) bool {
	return r.Start <= other.End && r.End >= other.Start
}
func (r Range) Merge(other Range) Range {
	return Range{
		Start: min(r.Start, other.Start),
		End:   max(r.End, other.End),
	}
}
func (r Range) String() string {
	return strconv.FormatInt(r.Start, 10) + "-" + strconv.FormatInt(r.End, 10)
}
