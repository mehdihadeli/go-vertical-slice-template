package database

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/config/environemnt"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database/options"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/dependency"

	"gorm.io/gorm"
)

func AddGorm(dependencies *dependency.ServiceCollection) {
	dependency.Add[*options.GormOptions](
		dependencies,
		func(sp *dependency.ServiceProvider) (*options.GormOptions, error) {
			environment := dependency.GetGenericRequiredService[environemnt.Environment](sp)
			return options.ConfigGormOptions(environment)
		},
	)

	dependency.Add[*gorm.DB](dependencies, func(sp *dependency.ServiceProvider) (*gorm.DB, error) {
		gormOptions := dependency.GetGenericRequiredService[*options.GormOptions](sp)
		return NewGorm(gormOptions)
	})
}
