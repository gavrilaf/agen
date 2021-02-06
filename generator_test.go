package agen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gavrilaf/agen"
)

func TestUniqueNumber(t *testing.T) {
	g := agen.NewGenerator()

	attempts := 1000

	collisions := 0
	mm := map[uint64]struct{}{}
	for i := 0; i < attempts; i++ {
		n := g.UniqueNumber()

		if _, ok := mm[n]; ok {
			collisions += 1
		}
		mm[n] = struct{}{}
		//t.Logf("%d: %d", i, n)
	}

	assert.Equal(t, 0, collisions)
}

func TestCodes(t *testing.T) {
	g := agen.NewGenerator()

	codes := g.Codes(1000)

	collisions := 0
	mm := map[string]struct{}{}
	for _, c := range codes {
		if _, ok := mm[c]; ok {
			collisions += 1
		}
		mm[c] = struct{}{}
		//t.Log(c)
	}

	assert.Equal(t, 0, collisions)
}

func BenchmarkCodes(b *testing.B) {
	g := agen.NewGenerator()

	b.Run("generate codes", func(b *testing.B) {
		g.Codes(b.N)
	})
}