package list

import "errors"

type ArrayList[T any] struct {
	arr    []T
	length int
}

func (al *ArrayList[T]) Prepend(vals ...T) {
	if cap(al.arr) < al.length+len(vals) {
		newArr := make([]T, max(cap(al.arr)*2, al.length+len(vals)))
		copy(newArr[len(vals):], al.arr)
		al.arr = newArr
	}
	for idx, v := range vals {
		al.arr[idx] = v
		al.length++
	}
}

func (al *ArrayList[T]) Append(vals ...T) {
	if cap(al.arr) < al.length+len(vals) {
		newArr := make([]T, max(cap(al.arr)*2, al.length+len(vals)))
		copy(newArr, al.arr)
		al.arr = newArr
	}
	for _, v := range vals {
		al.arr[al.length] = v
		al.length++
	}
}

func (al *ArrayList[T]) Get(idx int) (rVal T, err error) {
	if idx >= al.length {
		return rVal, errors.New("index out of bounds")
	}
	return al.arr[idx], nil
}

func (al *ArrayList[T]) Pop() (rVal T, err error) {
	if al.length == 0 {
		return rVal, errors.New("no elements left to pop")
	}
	rVal = al.arr[al.length-1]
	al.length--
	return rVal, nil
}

func (al *ArrayList[T]) Peek() (rVal T, err error) {
	if al.length == 0 {
		return rVal, errors.New("no elements left to peek")
	}
	return al.arr[al.length-1], nil
}
