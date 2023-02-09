package nestif

func BadRepeat(a, b int) (int, int) {
	var count int

	sum := b * b

	if a == 1 {
		sum += a * 2
		count = a // repeat
	} else if a == 2 {
		sum += a * 3
		count = a // repeat
	} else if a == 3 {
		sum += a * 5
		count = a // repeat
	} else if a == 4 {
		sum += a * 9
		count = a // repeat
	} else if a == 5 {
		sum += a * 10
		count = a // repeat
	} else {
		sum += a
		count = a // repeat
	}

	return sum, count
}

func RefactorRepeat(a, b int) (int, int) {
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

	count := a // 所有分支都是这样，那就提取为开头或者最后公共地方；本函数还可以直接放到 return 里

	return sum, count
}

func BadSimpleElse(a, b int, name string) (int, int) {
	var sum, c int
	if a == 999 && name == "3Mad" {
		// do something
		// many lines ...
		sum += a * 2
		c = b * a
	} else {
		c = b * (a + 1000) // else 里非常简单
	}

	return sum, c
}

func RefactorSimpleElse(a, b int, name string) (int, int) {
	var sum int

	c := b * (a + 1000)

	if a == 999 && name == "3Mad" {
		// do something
		// many lines ...
		sum += a * 2
		c = b * a
	}

	return sum, c
}
