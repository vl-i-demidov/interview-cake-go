package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeetingsOverlap(t *testing.T) {

	meetings := []*meeting{
		newMeeting(1, 2),
		newMeeting(2, 4),
	}

	actual := mergeRanges(meetings)
	expected := []*meeting{
		newMeeting(1, 4)}

	assert.Equal(t, expected, actual)
}
