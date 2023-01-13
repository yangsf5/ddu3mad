package v05

import (
	"errors"

	v01 "github.com/yangsf5/ddu3mad/refactor/generic/routine/v01"
)

var ErrSentry = errors.New("err sentry for routine doFunc")

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

func rMap[Ti DataInterface[Tk, Tv], Tk comparable, Tv any, To any](input Ti, funcDo func(Tk, Tv) (To, error)) []To {
	chanRet := make(chan To, input.Len())
	pool := v01.NewPool(5, input.Len())

	input.Range(func(k Tk, v Tv) {
		pool.AddOne()

		go func() {
			defer pool.DelOne()
			ret, err := funcDo(k, v)
			if err == nil {
				chanRet <- ret
			}
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
	return rMap(NewSliceData(input), func(k int, v Ti) (To, error) {
		return funcDo(k, v), nil
	})
}

func SliceExcludeError[Ti any, To any](input []Ti, funcDo func(int, Ti) (To, error)) []To {
	return rMap(NewSliceData(input), funcDo)
}

func SliceNoResult[Ti any](input []Ti, funcDo func(int, Ti)) {
	rMap(NewSliceData(input), func(k int, v Ti) (bool, error) {
		funcDo(k, v)
		return true, nil // 无用返回值，只是为了符合 rMap 接口
	})
}

func Map[Tk comparable, Tv any, To any](input map[Tk]Tv, funcDo func(k Tk, v Tv) To) []To {
	return rMap(NewMapData(input), func(k Tk, v Tv) (To, error) {
		return funcDo(k, v), nil
	})
}

func MapExcludeError[Tk comparable, Tv any, To any](input map[Tk]Tv, funcDo func(k Tk, v Tv) (To, error)) []To {
	return rMap(NewMapData(input), funcDo)
}

func MapNoResult[Tk comparable, Tv any](input map[Tk]Tv, funcDo func(k Tk, v Tv)) {
	rMap(NewMapData(input), func(k Tk, v Tv) (bool, error) {
		funcDo(k, v)
		return true, nil // 无用返回值，只是为了符合 rMap 接口
	})
}
