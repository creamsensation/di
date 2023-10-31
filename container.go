package di

type Container interface{}

type container struct {
	deps []dependency
}

func New(dependencies ...Dependency) Container {
	deps := make([]dependency, len(dependencies))
	for i, item := range dependencies {
		dep, ok := item.(dependency)
		if !ok {
			continue
		}
		deps[i] = dep
	}
	return &container{
		deps: deps,
	}
}
