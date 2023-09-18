package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	steps:= 0
	if n < 1 {
		return steps, fmt.Errorf("n must be positive")
	}
	for n > 1 {
		if (n%2 == 0) {
			n = n/2
		} else {
			n = n * 3 + 1
		}
		steps++
		}

	return steps, nil
}
