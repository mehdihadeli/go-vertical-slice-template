package defaultLogger

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
)

var l logger.Logger

func initLogger() {
	l = logger.NewZapLogger(
		&logger.LogOptions{CallerEnabled: false},
		constants.Dev,
	)
}

func GetLogger() logger.Logger {
	if l == nil {
		initLogger()
	}

	return l
}
