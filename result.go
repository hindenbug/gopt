package gopt

type Result[T any] struct {
	ok  T
	err error
}

func Ok[T any](ok T) Result[T] {
	return Result[T]{ok: ok}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsError returns true when value is invalid.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Error returns error when value is invalid or nil.
func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Ok() Option[T] {
	if r.IsOk() {
		return Some(r.ok)
	}

	return None[T]()
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("cant unwrap")
	}

	return r.ok
}
