package dependency

import (
	"reflect"
)

// ServiceCollection with various AddDependency methods
type ServiceCollection struct {
	dependencies []*DependencyType
}

// AddDependency adds a pre-constructed DependencyType
func (dc *ServiceCollection) AddDependency(dependencyType *DependencyType) {
	for i, dep := range dc.dependencies {
		if dep.Type == dependencyType.Type {
			// Override existing dependency
			dc.dependencies[i] = dependencyType
			return
		}
	}

	dc.dependencies = append(dc.dependencies, dependencyType)
}

// AddFromFunc adds a dependency from a function returning (interface{}, error)
func (dc *ServiceCollection) AddFromFunc(initializer func(sp *ServiceProvider) (interface{}, error), typ interface{}) {
	dc.AddDependency(&DependencyType{
		Initializer: initializer,
		Type:        reflect.TypeOf(typ).Elem(),
	})
}

func (dc *ServiceCollection) Build() *ServiceProvider {
	return NewServiceProvider(dc.dependencies)
}

// Add adds a type-safe dependency initializer
func Add[T any](dc *ServiceCollection, initializer func(sp *ServiceProvider) (T, error)) {
	dc.AddDependency(NewDependencyType[T](initializer))
}
