package handlers

import "iter"

type Vec2 struct {
	X float32
	Y float32
}

func Collect[T any](seq iter.Seq2[int, T], n uint64) []T {
	if seq == nil {
		return nil
	}

	out := make([]T, 0, n)
	for _, value := range seq {
		out = append(out, value)
	}
	return out
}
