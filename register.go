package di

func Register[T any](p Provider[T], configs ...Config) Dependency {
	return dependency{
		provider: p,
		config:   processConfig(p, configs...),
	}
}
