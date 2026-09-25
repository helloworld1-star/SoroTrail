package simtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOracle_DetectsGapsAndDuplicates(t *testing.T) {
	tests := []struct {
		name      string
		ledgers   []int64
		wantError bool
		errorType string
	}{
		{
			name:      "continuous sequential ledgers",
			ledgers:   []int64{100, 101, 102, 103},
			wantError: false,
		},
		{
			name:      "detected gap in sequence",
			ledgers:   []int64{100, 101, 103},
			wantError: true,
			errorType: "gap",
		},
		{
			name:      "detected duplicate ledger",
			ledgers:   []int64{100, 101, 101, 102},
			wantError: true,
			errorType: "duplicate",
		},
	}

	for _, tt := range tests {
		fn := tt
		t.Run(fn.name, func(t *testing.T) {
			err := checkLedgerSequence(fn.ledgers)
			if fn.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func checkLedgerSequence(ledgers []int64) error {
	seen := make(map[int64]bool)
	for i, l := range ledgers {
		if seen[l] {
			return errorsNew("duplicate ledger")
		}
		seen[l] = true
		if i > 0 && l != ledgers[i-1]+1 {
			return errorsNew("gap in ledger sequence")
		}
	}
	return nil
}

func errorsNew(s string) error {
	return &simError{msg: s}
}

type simError struct {
	msg string
}

func (e *simError) Error() string {
	return e.msg
}
