package application

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/catalogs/products/contracts"
)

func (a *Application) MapEndpoints() {
	a.ResolveRequiredDependencyFunc(func(endpoints []contracts.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	})
}
