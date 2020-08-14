package balancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	output, err := Solution("40 80 10 10 10")
	assert.Equal(t, nil, err)
	assert.Equal(t, 10, output)

	output, err = Solution("2 2 4 4 6 6 8 8 2 2")
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, output)

	output, err = Solution("2 a 4 4 6 6 8 8 2 2")
	assert.NotEqual(t, nil, err)
	assert.Equal(t, -1, output)
}

func TestBalancer(t *testing.T) {
	assert.Equal(t, 10, Balancer([]int{40, 80, 10, 10, 10}))
	assert.Equal(t, 0, Balancer([]int{2, 2, 4, 4, 6, 6, 8, 8, 2, 2}))

	assert.Equal(t, 2, Balancer([]int{1, 3}))
	assert.Equal(t, 0, Balancer([]int{10, 3, 10, 3, 4, 6, 5, 5}))
	assert.Equal(t, 3, Balancer([]int{10, 3, 10, 3, 7, 4, 6, 5, 5}))
}
