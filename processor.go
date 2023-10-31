package di

import (
	"fmt"
	"reflect"
)

func processConfig(provider any, configs ...Config) dependencyConfig {
	var result dependencyConfig
	for _, item := range configs {
		c, ok := item.(config)
		if !ok {
			continue
		}
		switch c.configType {
		case configName:
			result.name = fmt.Sprintf("%v", c.value)
		case configSingleton:
			result.singleton = true
		}
	}
	if len(result.name) == 0 && provider != nil {
		result.name = reflect.TypeOf(provider).String()
	}
	return result
}
