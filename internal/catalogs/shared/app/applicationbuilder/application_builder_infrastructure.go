package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb"
)

func (b *ApplicationBuilder) AddInfrastructure() error {
	config.AddAppConfig(b.ServiceCollection, b.Environment)

	database.AddGorm(b.ServiceCollection)

	echoweb.AddEcho(b.ServiceCollection)

	return nil
}
