package go_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MultipleParam[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)
	return param1, param2
}

func TestMultipleParam(t *testing.T) {
	name1, _ := MultipleParam[string, string]("Dyaksa", "Jauharuddin")
	_, wife := MultipleParam[string, string]("Dyaksa", "Kartika")

	assert.Equal(t, "Dyaksa", name1)
	assert.Equal(t, "Kartika", wife)
}
