package nestif

func BadManyLevel(a, b, c, d int, s []string) int {
	if a == 1 {
		// do something: a
		if b == 2 {
			// do something: b
			if c == 3 {
				// do something: c
				if d == 4 {
					// do something: d
					for range s {
						// do something
						// many lines ...
						// many lines ...
						// many lines ...
					}

					return 1
				}
			}
		}
	}

	return 0
}

func RefactorManyLevel(a, b, c, d int, s []string) int {
	if a != 1 {
		return 0
	}

	// do something: a

	if b != 2 {
		return 0
	}

	// do something: b

	if c != 3 {
		return 0
	}

	// do something: c

	if d != 4 {
		return 0
	}

	// do something: d

	for range s {
		// do something
		// many lines ...
		// many lines ...
		// many lines ...
	}

	return 1
}
