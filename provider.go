package di

type Provider[T any] interface {
	Provide(c Container) *T
}

func CreateProvider[T any](provider *T) Provider[T] {
	return &customProvider[T]{provider: provider}
}

type customProvider[T any] struct {
	provider *T
}

func (s *customProvider[T]) Provide(c Container) *T {
	return s.provider
}
