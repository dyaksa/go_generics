package go_generics

import (
	"fmt"
	"testing"
)

func IsWife[T comparable](value1, value2 T) bool {
	if value1 != value2 {
		return true
	} else {
		return false
	}
}

func TestIsWife(t *testing.T) {
	fmt.Println(IsWife[string]("Dyaksa", "Kartika"))
}
