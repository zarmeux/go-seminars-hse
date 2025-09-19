package main

import (
	"slices"
	"strings"
)

func slicesEqual(a, b []int) bool {
	return slices.Equal(a, b)
}

func slicesEqualFunc(a, b []string) bool {
	return slices.EqualFunc(a, b, func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
}

func slicesContains(a []string, key string) bool {
	return slices.Contains(a, key)
}

func slicesIndex(a []string, key string) int {
	return slices.Index(a, key)
}

func slicesIndexFunc(a []string) int {
	return slices.IndexFunc(a, func(s string) bool {
		return len(s) > 3
	})
}
