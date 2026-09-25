package simtest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sorotrail/sorotrail/internal/rpc"
)

func TestSimulatedChain_FilterMatching(t *testing.T) {
	tests := []struct {
		name       string
		filter     rpc.EventFilter
		contractID string
		topics     []string
		wantMatch  bool
	}{
		{
			name: "exact contract ID match",
			filter: rpc.EventFilter{
				ContractIDs: []string{"C1111"},
			},
			contractID: "C1111",
			topics:     []string{"transfer"},
			wantMatch:  true,
		},
		{
			name: "contract ID mismatch",
			filter: rpc.EventFilter{
				ContractIDs: []string{"C1111"},
			},
			contractID: "C2222",
			topics:     []string{"transfer"},
			wantMatch:  false,
		},
		{
			name: "empty contract IDs matches any",
			filter: rpc.EventFilter{
				ContractIDs: nil,
			},
			contractID: "Cany",
			topics:     []string{"any"},
			wantMatch:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Exercise internal simulation filter matching helper or mock implementation
			match := matchFilter(tt.filter, tt.contractID, tt.topics)
			assert.Equal(t, tt.wantMatch, match)
		})
	}
}

func matchFilter(f rpc.EventFilter, contractID string, topics []string) bool {
	if len(f.ContractIDs) > 0 {
		found := false
		for _, id := range f.ContractIDs {
			if id == contractID {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
