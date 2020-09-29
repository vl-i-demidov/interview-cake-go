package arrays

import (
	"sort"
)

type meeting struct {
	Start int
	End   int
}

func newMeeting(start, end int) *meeting {
	m := new(meeting)
	m.Start = start
	m.End = end
	return m
}

func mergeRanges(a []*meeting) []*meeting {
	// todo make a copy of array
	sort.Slice(a[:], func(i, j int) bool {
		return a[i].Start < a[j].Start
	})
	var result []*meeting
	start := a[0].Start
	end := a[0].End

	for i := 0; i < len(a); i++ {
		m := a[i]
		if m.Start > end {
			merged := newMeeting(start, end)
			result = append(result, merged)
			start = m.Start
			end = m.End
		} else if m.End > end {
			end = m.End
		}
	}

	merged := newMeeting(start, end)
	result = append(result, merged)
	return result
}
