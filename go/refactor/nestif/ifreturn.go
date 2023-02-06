package ifreturn

func BadFunc(a, b, c, d int, s []string) int {
	if a != 1 {
		if b != 2 {
			if c != 3 {
				if d != 4 {
					for range s {
						// do something
						// ...
					}

					return 1
				}
			}
		}
	}

	return 0
}

func Refactor(a, b, c, d int, s []string) int {
	if a == 1 {
		return 0
	}

	if b == 2 {
		return 0
	}

	if c == 3 {
		return 0
	}

	if d == 4 {
		return 0
	}

	for range s {
		// do something
		// ...
	}

	return 1
}
