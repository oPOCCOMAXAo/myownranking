package ginutils

import "github.com/gin-gonic/gin"

// CtxTypedPointer is helper for get-set pointer values to gin.Context.
//
// Example:
//
//	const SomeTypeValue CtxTypedPointer[SomeType] = "some_context_key"
//
//	func SomeHandler(ctx *gin.Context) {
//		// ...
//		SomeTypeValue.Set(ctx, &someValue)
//		// ...
//		someValue := SomeTypeValue.Get(ctx)
//	}
type CtxTypedPointer[T any] string

func NewTyped[T any](field string) CtxTypedPointer[T] {
	return CtxTypedPointer[T](field)
}

func (typed CtxTypedPointer[T]) Set(ctx *gin.Context, value *T) {
	ctx.Set(string(typed), value)
}

func (typed CtxTypedPointer[T]) Constant(value *T) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		typed.Set(ctx, value)
	}
}

func (typed CtxTypedPointer[T]) GetOK(ctx *gin.Context) (*T, bool) {
	value, ok := ctx.Get(string(typed))
	if !ok {
		return nil, false
	}

	res, ok := value.(*T)

	return res, ok
}

func (typed CtxTypedPointer[T]) Get(ctx *gin.Context) *T {
	res, _ := typed.GetOK(ctx)

	return res
}

func (typed CtxTypedPointer[T]) IsEmpty(ctx *gin.Context) bool {
	res, ok := typed.GetOK(ctx)

	return !ok || res == nil
}

// CtxTypedValue is helper for get-set values to gin.Context.
//
// Example:
//
//	const SomeTypeValue CtxTypedValue[SomeType] = "some_context_key"
//
//	func SomeHandler(ctx *gin.Context) {
//		// ...
//		SomeTypeValue.Set(ctx, someValue)
//		// ...
//		someValue := SomeTypeValue.Get(ctx)
//	}
type CtxTypedValue[T any] string

func NewTypedValue[T any](field string) CtxTypedValue[T] {
	return CtxTypedValue[T](field)
}

func (typed CtxTypedValue[T]) Set(ctx *gin.Context, value T) {
	ctx.Set(string(typed), value)
}

func (typed CtxTypedValue[T]) Constant(value T) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		typed.Set(ctx, value)
	}
}

func (typed CtxTypedValue[T]) GetOK(ctx *gin.Context) (T, bool) {
	value, ok := ctx.Get(string(typed))
	if !ok {
		var zero T

		return zero, false
	}

	res, ok := value.(T)

	return res, ok
}

func (typed CtxTypedValue[T]) Get(ctx *gin.Context) T {
	res, _ := typed.GetOK(ctx)

	return res
}

func (typed CtxTypedValue[T]) IsEmpty(ctx *gin.Context) bool {
	_, ok := typed.GetOK(ctx)

	return !ok
}
