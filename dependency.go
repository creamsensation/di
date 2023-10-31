package di

type Dependency interface {
	Provider() any
}

type dependency struct {
	provider any
	config   dependencyConfig
}

func (s dependency) Provider() any {
	return s.provider
}
