package main

// N is an alias for an unallocated struct
func N(size int) []struct{} {
	return make([]struct{}, size)
}