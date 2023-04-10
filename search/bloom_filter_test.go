package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BloomFilter(t *testing.T) {
	bf := NewBloomFilter(100, 3)
	bf.Add("hello")
	bf.Add("world")
	bf.Add("golang")

	assert.Equal(t, true, bf.Contains("hello"))
	assert.Equal(t, true, bf.Contains("world"))
	assert.Equal(t, true, bf.Contains("golang"))
	assert.Equal(t, false, bf.Contains("java"))
}
