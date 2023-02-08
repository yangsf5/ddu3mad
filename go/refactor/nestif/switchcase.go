package nestif

func BadManyBranch(a, b int, s []string) int {
	sum := b * b

	if a == 1 {
		sum += a * 2
	} else if a == 2 {
		sum += a * 3
	} else if a == 3 {
		sum += a * 5
	} else if a == 4 {
		sum += a * 9
	} else if a == 5 {
		sum += a * 10
	} else {
		sum += a
	}

	return sum
}

func RefactorManyBranch(a, b int, s []string) int {
	sum := b * b

	switch a {
	case 1:
		sum += a * 2
	case 2:
		sum += a * 3
	case 3:
		sum += a * 5
	case 4:
		sum += a * 9
	case 5:
		sum += a * 10
	default:
		sum += a
	}

	return sum
}

func BadManyCondition(a, b int, s []string) int {
	sum := b * b

	if a == 1 || a == 3 || a == 4 || a == 7 || a == 8 {
		sum += a * 2
	} else if a == 9 || a == 13 || a == 14 {
		sum += a * 10
	} else {
		sum += a
	}

	return sum
}

func RefactorManyCondition(a, b int, s []string) int {
	sum := b * b

	switch a {
	case 1, 3, 4, 7, 8:
		sum += a * 2
	case 9, 13, 14:
		sum += a * 10
	default:
		sum += a
	}

	return sum
}
