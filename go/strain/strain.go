package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T interface{}](input []T, predicate func(T) bool) []T {
	var newSlice []T

	for _, v := range input {
		if predicate(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func Discard[T interface{}](input []T, predicate func(T) bool) []T {
	var newSlice []T

	for _, v := range input {
		if !predicate(v) {
			newSlice = append(newSlice, v)
		}

	}
	return newSlice
}
