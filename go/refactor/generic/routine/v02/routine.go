package v02

import v01 "github.com/yangsf5/ddu3mad/refactor/generic/routine/v01"

func Slice[Ti any, To any](input []Ti, funcDo func(int, Ti) To) []To {
	chanRet := make(chan To, len(input))
	pool := v01.NewPool(5, len(input))

	for k, v := range input {
		pool.AddOne()

		go func(key int, value Ti) {
			defer pool.DelOne()

			chanRet <- funcDo(key, value)
		}(k, v)
	}

	pool.Waiting()
	close(chanRet)

	rets := make([]To, 0, len(input))
	for ret := range chanRet {
		rets = append(rets, ret)
	}

	return rets
}

func Map[Tk comparable, Tv any, To any](input map[Tk]Tv, funcDo func(k Tk, v Tv) To) []To {
	chanRet := make(chan To, len(input))
	pool := v01.NewPool(5, len(input))

	for k, v := range input {
		pool.AddOne()

		go func(key Tk, value Tv) {
			defer pool.DelOne()

			chanRet <- funcDo(key, value)
		}(k, v)
	}

	pool.Waiting()
	close(chanRet)

	rets := make([]To, 0, len(input))
	for ret := range chanRet {
		rets = append(rets, ret)
	}

	return rets
}
