package tasks

func invertMap[K comparable, V comparable](_ map[K]V) map[V]K {
	return make(map[V]K)
}
