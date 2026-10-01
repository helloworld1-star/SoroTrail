package simtest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVirtualClock_DeterministicAdvancement(t *testing.T) {
	base := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	clock := NewSimVirtualClock(base)

	assert.Equal(t, base, clock.Now())

	clock.Advance(5 * time.Second)
	assert.Equal(t, base.Add(5*time.Second), clock.Now())

	clock.Advance(10 * time.Minute)
	assert.Equal(t, base.Add(5*time.Second).Add(10*time.Minute), clock.Now())
}

type SimVirtualClock struct {
	current time.Time
}

func NewSimVirtualClock(start time.Time) *SimVirtualClock {
	return &SimVirtualClock{current: start}
}

func (vc *SimVirtualClock) Now() time.Time {
	return vc.current
}

func (vc *SimVirtualClock) Advance(d time.Duration) {
	vc.current = vc.current.Add(d)
}
