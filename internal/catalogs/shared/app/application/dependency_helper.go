package application

import (
	"fmt"
	"reflect"

	"go.uber.org/dig"
)

// ResolveRequiredDependency resolves a dependency from the container by type and panics if it fails
func ResolveRequiredDependency[T any](c *dig.Container) T {
	var result T
	err := c.Invoke(func(t T) {
		result = t
	})
	if err != nil {
		panic(fmt.Sprintf("failed to resolve dependency for type %s: %v", reflect.TypeOf(result).String(), err))
	}
	return result
}

// ResolveDependency resolves a dependency from the container by type and returns an error if it fails
func ResolveDependency[T any](c *dig.Container) (T, error) {
	var result T
	err := c.Invoke(func(t T) {
		result = t
	})
	return result, err
}
