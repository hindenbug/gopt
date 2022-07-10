package gopt

import "fmt"

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (o Option[T]) Some(value T) Option[T] {
	o.value = &value
	return o
}

func (o Option[T]) None() Option[T] {
	return o
}

// Option[T].IsSome() returns true.
func (o Option[T]) IsSome() bool {
	return !o.IsNone()
}

// Option[T].IsNone() returns false.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("cannot unwrap on None")
	}

	return *o.value
}

func (o Option[T]) UnwrapOr(defVal T) T {
	if o.IsNone() {
		return defVal
	}

	return *o.value
}

func (o Option[T]) Expect(msg string) T {
	if o.IsNone() {
		panic(fmt.Errorf("%s", msg))
	}

	return *o.value
}
