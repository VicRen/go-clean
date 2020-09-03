package registry

import (
	"context"
	"fmt"
	"reflect"
)

type ConfigCreator func(ctx context.Context, config interface{}) (interface{}, error)

var typeCreatorRegistry = make(map[reflect.Type]ConfigCreator)

func RegisterConfig(config interface{}, configCreator ConfigCreator) error {
	configType := reflect.TypeOf(config)
	if _, found := typeCreatorRegistry[configType]; found {
		return fmt.Errorf("%s is already registered", configType.Name())
	}
	typeCreatorRegistry[configType] = configCreator
	return nil
}
