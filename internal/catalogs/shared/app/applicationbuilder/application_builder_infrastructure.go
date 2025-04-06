package applicationbuilder

import (
	"github.com/mehdihadeli/go-vertical-slice-template/config"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/http/echoweb"
)

func (b *ApplicationBuilder) AddInfrastructure() {
	err := config.AddAppConfig(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = database.AddGorm(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = echoweb.AddEcho(b.Container)
	if err != nil {
		b.Logger.Fatal(err)
	}
}
