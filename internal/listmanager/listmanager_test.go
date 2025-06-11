package listmanager_test

import (
	"testing"

	"github.com/harshau007/listmanager/internal/listmanager"
	"github.com/stretchr/testify/assert"
)

func TestAddAndList(t *testing.T) {
	tests := []struct {
		inputs   []float64
		expected [][]float64
	}{
		{
			inputs:   []float64{5, 10, -6},
			expected: [][]float64{{5}, {5, 10}, {9}},
		},
		{
			inputs:   []float64{1, 2, 3, 4},
			expected: [][]float64{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		lm := listmanager.New()
		for i, in := range tt.inputs {
			res := lm.Add(in)
			assert.Equal(t, tt.expected[i], res)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	lm := listmanager.New()
	assert.Empty(t, lm.List())
	assert.Equal(t, []float64{0}, listmanager.New().Add(0))
	assert.Equal(t, []float64{0, 0}, func() []float64 { lm := listmanager.New(); lm.Add(0); return lm.Add(0) }())
	assert.Equal(t, []float64{1}, func() []float64 { lm := listmanager.New(); lm.Add(1000000); return lm.Add(-999999) }())
	assert.Empty(t, func() []float64 { lm := listmanager.New(); lm.Add(5); lm.Add(3); return lm.Add(-8) }())
	assert.Empty(t, func() []float64 { lm := listmanager.New(); lm.Add(5); return lm.Add(-10) }())
}

func TestSignsMatch(t *testing.T) {
	tests := []struct {
		a, b float64
		exp  bool
	}{
		{1, 2, true}, {-1, -2, true}, {0, 5, true}, {0, -5, true}, {5, 0, true}, {-5, 0, true}, {0, 0, true}, {1, -1, false}, {-1, 1, false},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.exp, listmanager.New().SignsMatch(tt.a, tt.b))
	}
}
