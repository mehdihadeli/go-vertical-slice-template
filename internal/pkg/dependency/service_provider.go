package dependency

import (
	"fmt"
	"reflect"
)

type ServiceProvider struct {
	dependencies []*DependencyType
	// simple map storage for singleton instances
	instances map[reflect.Type]interface{}
}

func NewServiceProvider(dependencies []*DependencyType) *ServiceProvider {
	return &ServiceProvider{
		dependencies: dependencies,
		instances:    make(map[reflect.Type]interface{}),
	}
}

func (sp *ServiceProvider) GetService(targetType reflect.Type) (interface{}, error) {
	// Check cache first
	if instance, exists := sp.instances[targetType]; exists {
		return instance, nil
	}

	for _, dep := range sp.dependencies {
		if dep.Type == targetType {
			instance, err := dep.Initializer(sp)
			if err != nil {
				return nil, fmt.Errorf("initialization failed for %v: %w", targetType, err)
			}

			// Cache the instance
			sp.instances[targetType] = instance

			return instance, nil
		}
	}

	return nil, fmt.Errorf("dependency not found for type: %v", targetType)
}

func (sp *ServiceProvider) GetRequiredService(targetType reflect.Type) interface{} {
	instance, err := sp.GetService(targetType)
	if err != nil {
		panic(fmt.Sprintf("dependency error: %v", err))
	}

	return instance
}

// GetGenericService retrieves a dependency by type
func GetGenericService[T any](sp *ServiceProvider) (T, error) {
	var zero T
	targetType := reflect.TypeOf((*T)(nil)).Elem()

	// Check cache first
	if instance, exists := sp.instances[targetType]; exists {
		if typedInstance, ok := instance.(T); ok {
			return typedInstance, nil
		}
		return zero, fmt.Errorf("cached instance has wrong type for %v", targetType)
	}

	for _, dep := range sp.dependencies {
		if dep.Type == targetType {
			instance, err := dep.Initializer(sp)
			if err != nil {
				return zero, fmt.Errorf("initialization failed for %v: %w", targetType, err)
			}

			if typedInstance, ok := instance.(T); ok {
				// Cache the instance
				sp.instances[targetType] = typedInstance

				return typedInstance, nil
			}
			return zero, fmt.Errorf("type mismatch for %v: expected %T, got %T",
				targetType, zero, instance)
		}
	}

	return zero, fmt.Errorf("dependency not found for type: %v", targetType)
}

// GetGenericRequiredService retrieves a dependency or panics
func GetGenericRequiredService[T any](sp *ServiceProvider) T {
	instance, err := GetGenericService[T](sp)
	if err != nil {
		panic(fmt.Sprintf("dependency error: %v", err))
	}

	return instance
}
