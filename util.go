package di

func callProvide[T any](c *container, dep dependency) *T {
	provider, ok := dep.provider.(Provider[T])
	if !ok {
		return dep.provider.(*T)
	}
	return provider.Provide(c)
}

func callProvideValue[T any](c *container, dep dependency) T {
	provider, ok := dep.provider.(Provider[T])
	if !ok {
		return dep.provider.(T)
	}
	return *provider.Provide(c)
}

func findDependencyWithName[T comparable](c *container, name string) dependency {
	for _, d := range c.deps {
		if d.config.name != name {
			continue
		}
		result := getDependency[T](d)
		if result.provider == nil {
			continue
		}
		return result
	}
	return dependency{}
}

func findDependencyWithType[T comparable](c *container) dependency {
	for _, d := range c.deps {
		result := getDependency[T](d)
		if result.provider == nil {
			continue
		}
		return result
	}
	return dependency{}
}

func getDependency[T comparable](d dependency) dependency {
	switch p := d.provider.(type) {
	case *customProvider[T]:
		if d.config.singleton {
			return dependency{
				provider: p.provider,
				config:   d.config,
			}
		}
		return dependency{
			provider: new(T),
			config:   d.config,
		}
	case *T:
		if !d.config.singleton {
			d.provider = new(T)
		}
		return d
	default:
		return dependency{}
	}
}
