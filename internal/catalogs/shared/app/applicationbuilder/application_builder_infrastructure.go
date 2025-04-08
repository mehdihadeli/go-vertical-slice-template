package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database/options"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb"
	config2 "github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb/config"
)

func (b *ApplicationBuilder) AddInfrastructure() error {
	appOptions, err := config.ConfigAppOptions(b.Environment)
	if err != nil {
		return err
	}
	b.appOptions = appOptions

	gormOptions, err := options.ConfigGormOptions(b.Environment)
	if err != nil {
		return err
	}
	gormDB, err := database.NewGorm(gormOptions)
	if err != nil {
		return err
	}
	b.gormDB = gormDB

	echoOptions, err := config2.ConfigEchoOptions(b.Environment)
	if err != nil {
		return err
	}
	b.echoOptions = echoOptions

	e, err := echoweb.NewEcho(b.Logger)
	if err != nil {
		return err
	}
	b.echo = e

	return nil
}
