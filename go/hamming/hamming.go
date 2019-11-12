package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("invalid lengths")
	}

	length := len(a)
	counter := 0

	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}
