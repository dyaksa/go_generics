package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var str = Length[string]("Dyaksa")
	assert.True(t, true)
	assert.Equal(t, "Dyaksa", str)

	var num = Length[int](10)
	assert.Equal(t, 10, num)
}
