package di

import "reflect"

type Config interface{}

type config struct {
	configType int
	value      any
}

const (
	configName = iota
	configSingleton
)

func Interface[T any]() Config {
	return config{
		configType: configName,
		value:      reflect.TypeOf((*T)(nil)).Elem().String(),
	}
}

func Name(name string) Config {
	return config{
		configType: configName,
		value:      name,
	}
}

func Singleton() Config {
	return config{
		configType: configSingleton,
	}
}
