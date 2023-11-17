package scheduling_test

import (
	"testing"

	"github.com/pijusn/files-rename/lib/scheduling"
	"github.com/stretchr/testify/require"
)

func TestOptimize(t *testing.T) {
	testCases := []struct {
		Description string
		Input       []scheduling.Task
		Output      []scheduling.Task
	}{
		{
			Description: "original order is OK",
			Input: []scheduling.Task{
				{
					NameSource: "A",
					NameTarget: "_1",
				},
				{
					NameSource: "B",
					NameTarget: "_2",
				},
				{
					NameSource: "C",
					NameTarget: "_3",
				},
			},
			Output: []scheduling.Task{
				{
					NameSource: "A",
					NameTarget: "_1",
				},
				{
					NameSource: "B",
					NameTarget: "_2",
				},
				{
					NameSource: "C",
					NameTarget: "_3",
				},
			},
		},
		{
			Description: "reorder to resolve collisions",
			Input: []scheduling.Task{
				{
					NameSource: "A",
					NameTarget: "_1",
				},
				{
					NameSource: "_1",
					NameTarget: "_2",
				},
				{
					NameSource: "_2",
					NameTarget: "_3",
				},
			},
			Output: []scheduling.Task{
				{
					NameSource: "_2",
					NameTarget: "_3",
				},
				{
					NameSource: "_1",
					NameTarget: "_2",
				},
				{
					NameSource: "A",
					NameTarget: "_1",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			items := make([]scheduling.Task, len(tc.Input))
			copy(items, tc.Input)
			scheduling.Optimize(items)
			require.Equal(t, tc.Output, items)
		})
	}
}
