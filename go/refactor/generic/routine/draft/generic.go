package routine

import v01 "github.com/yangsf5/ddu3mad/refactor/generic/routine/v01"

type MapDataInterface[Tk comparable, Tv any] interface {
	Len() int
	Range(func(Tk, Tv))
}

type SliceRoutine[T any] struct {
	data []T
}

func NewSliceRoutine[T any](data []T) *SliceRoutine[T] {
	return &SliceRoutine[T]{data: data}
}

func (m *SliceRoutine[T]) Len() int {
	return len(m.data)
}

func (m *SliceRoutine[T]) Range(funcDo func(int, T)) {
	for k, v := range m.data {
		funcDo(k, v)
	}
}

type MapRoutine[Tk comparable, Tv any] struct {
	data map[Tk]Tv
}

func NewMapRoutine[Tk comparable, Tv any](data map[Tk]Tv) *MapRoutine[Tk, Tv] {
	return &MapRoutine[Tk, Tv]{data: data}
}

func (m *MapRoutine[Tk, Tv]) Len() int {
	return len(m.data)
}

func (m *MapRoutine[Tk, Tv]) Range(funcDo func(Tk, Tv)) {
	for k, v := range m.data {
		funcDo(k, v)
	}
}

type Routine[Ti MapDataInterface[Tik, Tiv], Tik comparable, Tiv any, To any] struct {
	input Ti
}

func NewRoutine[Ti MapDataInterface[Tik, Tiv], Tik comparable, Tiv any, To any](
	input Ti,
) *Routine[Ti, Tik, Tiv, To] {
	return &Routine[Ti, Tik, Tiv, To]{input: input}
}

func (m *Routine[Ti, Tik, Tiv, To]) Map(funcDo func(Tik, Tiv) To) []To {
	chanRet := make(chan To, m.input.Len())
	pool := v01.NewPool(5, m.input.Len())

	m.input.Range(func(k Tik, v Tiv) {
		pool.AddOne()

		go func() {
			defer pool.DelOne()

			chanRet <- funcDo(k, v)
		}()
	})

	pool.Waiting()
	close(chanRet)

	rets := make([]To, 0, m.input.Len())
	for ret := range chanRet {
		rets = append(rets, ret)
	}

	return rets
}

func (m *Routine[Ti, Tik, Tiv, To]) MapNoRet(funcDo func(Tik, Tiv)) {
	pool := v01.NewPool(5, m.input.Len())

	m.input.Range(func(k Tik, v Tiv) {
		pool.AddOne()

		go func() {
			defer pool.DelOne()

			funcDo(k, v)
		}()
	})

	pool.Waiting()
}

func FlatSlice[T any](in [][]T) []T {
	rets := []T{}
	for _, v := range in {
		rets = append(rets, v...)
	}

	return rets
}
