package slice

func Reduce[T any, U any](items []T, initial U, reducer func(U, T) U) U {
	accumulated := initial
	for _, item := range items {
		accumulated = reducer(accumulated, item)
	}
	return accumulated
}

func Map[T any, U any](items []T, mapper func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

type FilterFunc[T any] func(T) bool

func Filter[T any](items []T, fn FilterFunc[T]) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}
