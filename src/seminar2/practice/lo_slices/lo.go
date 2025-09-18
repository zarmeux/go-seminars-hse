package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func loMap() {
	names := []string{"Alice", "Bob", "Charlie"}

	upperNative := make([]string, len(names))
	for i, name := range names {
		upperNative[i] = strings.ToUpper(name)
	}

	upper := lo.Map(names, func(name string, _ int) string {
		return strings.ToUpper(name)
	})

	fmt.Println(upperNative, upper)
}

func loUniq() (unique []int) {
	nums := []int{1, 2, 2, 3, 1, 4}
	unique = lo.Uniq(nums)
	fmt.Println(unique)
	return
}

func loChunk() {
	arr := []int{1, 2, 3, 4, 5}
	chunks := lo.Chunk(arr, 2)
	fmt.Println(chunks)
}
