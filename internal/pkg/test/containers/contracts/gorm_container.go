package contracts

import (
	"context"
	"testing"

	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/database/options"
)

type PostgresContainerOptions struct {
	Database  string
	Host      string
	Port      string
	HostPort  int
	UserName  string
	Password  string
	ImageName string
	Name      string
	Tag       string
}

type GormContainer interface {
	PopulateContainerOptions(
		ctx context.Context,
		t *testing.T,
		options ...*PostgresContainerOptions,
	) (*options.GormOptions, error)
	Cleanup(ctx context.Context) error
}
