package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[string]int{"1": 1, "2": 2}
	m2 := map[string]int{"1": 1, "2": 2}

	fmt.Println(MapsEqual(m1, m2))
	fmt.Println(MapsEqualReflect(m1, m2))
}

// Iterating through each mao and comparing the values of each key.
func MapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// Using reflect.DeepEqaul
func MapsEqualReflect(m1, m2 map[string]int) bool {
	return reflect.DeepEqual(m1, m2)
}
