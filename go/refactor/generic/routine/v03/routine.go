package v03

import (
	v01 "github.com/yangsf5/ddu3mad/refactor/generic/routine/v01"
)

type DataInterface[Tk comparable, Tv any] interface {
	Len() int
	Range(func(Tk, Tv))
}

type SliceData[T any] struct {
	data []T
}

func NewSliceData[T any](data []T) *SliceData[T] {
	return &SliceData[T]{data: data}
}

func (m *SliceData[T]) Len() int {
	return len(m.data)
}

func (m *SliceData[T]) Range(funcDo func(int, T)) {
	for k, v := range m.data {
		funcDo(k, v)
	}
}

type MapData[Tk comparable, Tv any] struct {
	data map[Tk]Tv
}

func NewMapData[Tk comparable, Tv any](data map[Tk]Tv) *MapData[Tk, Tv] {
	return &MapData[Tk, Tv]{data: data}
}

func (m *MapData[Tk, Tv]) Len() int {
	return len(m.data)
}

func (m *MapData[Tk, Tv]) Range(funcDo func(Tk, Tv)) {
	for k, v := range m.data {
		funcDo(k, v)
	}
}

func rMap[Ti DataInterface[Tk, Tv], Tk comparable, Tv any, To any](input Ti, funcDo func(Tk, Tv) To) []To {
	chanRet := make(chan To, input.Len())
	pool := v01.NewPool(5, input.Len())

	input.Range(func(k Tk, v Tv) {
		pool.AddOne()

		go func() {
			defer pool.DelOne()
			chanRet <- funcDo(k, v)
		}()
	})

	pool.Waiting()
	close(chanRet)

	rets := make([]To, 0, input.Len())
	for ret := range chanRet {
		rets = append(rets, ret)
	}

	return rets
}

func Slice[Ti any, To any](input []Ti, funcDo func(int, Ti) To) []To {
	return rMap(NewSliceData(input), funcDo)
}

func Map[Tk comparable, Tv any, To any](input map[Tk]Tv, funcDo func(k Tk, v Tv) To) []To {
	return rMap(NewMapData(input), funcDo)
}
