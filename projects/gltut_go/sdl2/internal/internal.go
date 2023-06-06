package internal

import (
	"unsafe"
)

type CObjWrapper[T any] struct {
	CObj T
}

func Unwrap[W any, T any](w *W) T {
	var w1 = (*CObjWrapper[T])(unsafe.Pointer(w))
	return w1.CObj
}

func UnwrapRef[W any, T any](w *W) *T {
	var w1 = (*CObjWrapper[T])(unsafe.Pointer(w))
	return &w1.CObj
}

func WrapNew[W any, T any](o T) *W {
	var w = new(W)
	var w1 = (*CObjWrapper[T])(unsafe.Pointer(w))
	w1.CObj = o
	return w
}
