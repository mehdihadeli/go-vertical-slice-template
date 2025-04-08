package dependency

import "reflect"

type DependencyType struct {
	Initializer func(sp *ServiceProvider) (interface{}, error)
	Type        reflect.Type
}

func NewDependencyType[T any](initializer func(sp *ServiceProvider) (T, error)) *DependencyType {
	var zero T
	return &DependencyType{
		Initializer: func(sp *ServiceProvider) (interface{}, error) {
			return initializer(sp)
		},
		Type: reflect.TypeOf(&zero).Elem(),
	}
}
