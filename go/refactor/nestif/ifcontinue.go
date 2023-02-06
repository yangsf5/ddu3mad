package nestif

func BadLoopManyLevel(a, b, c, d int, s []string) int {
	sum := 0

	for range s {
		if a != 1 {
			// do something: a
			if b != 2 {
				// do something: b
				if c != 3 {
					// do something: c
					if d != 4 {
						// do something: d
						// many lines ...
						// many lines ...
						// many lines ...
						sum += 1
					}
				}
			}
		}
	}

	return sum
}

func RefactorLoopManyLevel(a, b, c, d int, s []string) int {
	sum := 0

	for range s {
		if a == 1 {
			continue
		}

		// do something: a

		if b == 2 {
			continue
		}

		// do something: b

		if c == 3 {
			continue
		}

		// do something: c

		if d == 4 {
			continue
		}

		// do something: d
		// many lines ...
		// many lines ...
		// many lines ...
		sum += 1
	}

	return sum
}
