package simtest

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeededGeneration_ReproducesIdenticalScenarios(t *testing.T) {
	seed := int64(42)

	gen1 := generateScenario(seed, 10)
	gen2 := generateScenario(seed, 10)
	genDiff := generateScenario(123, 10)

	assert.Equal(t, gen1, gen2, "identical seeds must produce identical scenarios")
	assert.NotEqual(t, gen1, genDiff, "different seeds must produce distinct scenarios")
}

func generateScenario(seed int64, count int) []int {
	rng := rand.New(rand.NewSource(seed))
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = rng.Intn(1000)
	}
	return result
}
