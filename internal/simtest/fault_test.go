package simtest

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaultScheduler_FiresOnNominatedCall(t *testing.T) {
	tests := []struct {
		name        string
		targetCall  int
		totalCalls  int
		wantFaultAt []int
	}{
		{
			name:        "fault on third call",
			targetCall:  3,
			totalCalls:  5,
			wantFaultAt: []int{3},
		},
		{
			name:        "fault on first call",
			targetCall:  1,
			totalCalls:  3,
			wantFaultAt: []int{1},
		},
	}

	for _, tt := range tests {
		fn := tt
		t.Run(fn.name, func(t *testing.T) {
			sched := NewFaultScheduler(fn.targetCall, errors.New("injected fault"))
			var firedCalls []int

			for i := 1; i <= fn.totalCalls; i++ {
				err := sched.Check(i)
				if err != nil {
					firedCalls = append(firedCalls, i)
				}
			}

			assert.Equal(t, fn.wantFaultAt, firedCalls)
		})
	}
}

type FaultScheduler struct {
	target int
	err    error
}

func NewFaultScheduler(target int, err error) *FaultScheduler {
	return &FaultScheduler{target: target, err: err}
}

func (fs *FaultScheduler) Check(callIndex int) error {
	if callIndex == fs.target {
		return fs.err
	}
	return nil
}
