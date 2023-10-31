package di

func Provide[T comparable](c Container, configs ...Config) *T {
	cfg := processConfig(nil, configs...)
	cc, ok := c.(*container)
	if !ok {
		return new(T)
	}
	var provider *T
	var dep dependency
	nameExist := len(cfg.name) > 0
	if nameExist {
		dep = findDependencyWithName[T](cc, cfg.name)
	}
	if !nameExist {
		dep = findDependencyWithType[T](cc)
	}
	if !dep.config.singleton {
		dep.provider = new(T)
		provider = callProvide[T](cc, dep)
	}
	if dep.config.singleton {
		provider = dep.provider.(*T)
	}
	return provider
}

func ProvideValue[T comparable](c Container, configs ...Config) T {
	cfg := processConfig(nil, configs...)
	cc, ok := c.(*container)
	if !ok {
		return *new(T)
	}
	var provider T
	var dep dependency
	nameExist := len(cfg.name) > 0
	if nameExist {
		dep = findDependencyWithName[T](cc, cfg.name)
	}
	if !nameExist {
		dep = findDependencyWithType[T](cc)
	}
	if !dep.config.singleton {
		provider = callProvideValue[T](cc, dep)
	}
	if dep.config.singleton {
		provider = dep.provider.(T)
	}
	return provider
}
