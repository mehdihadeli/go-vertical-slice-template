package defaultLogger

import (
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/constants"
	"github.com/mehdihadeli/go-vertical-slice-template/internal/pkg/logger"
)

var log logger.Logger

func initLogger() {
	l, err := logger.NewZapLogger(
		constants.Dev,
		&logger.LogOptions{LogLevel: "debug"},
	)
	if err != nil {
		panic(err)
	}
	log = l
}

func GetLogger() logger.Logger {
	if log == nil {
		initLogger()
	}

	return log
}
