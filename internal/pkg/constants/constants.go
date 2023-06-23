package constants

import "time"

const (
	ConfigPath           = "CONFIG_PATH"
	AppEnv               = "APP_ENV"
	AppRootPath          = "APP_ROOT"
	PROJECT_NAME_ENV     = "PROJECT_NAME"
	Json                 = "json"
	GRPC                 = "GRPC"
	METHOD               = "METHOD"
	NAME                 = "NAME"
	METADATA             = "METADATA"
	REQUEST              = "REQUEST"
	REPLY                = "REPLY"
	TIME                 = "TIME"
	MaxHeaderBytes       = 1 << 20
	StackSize            = 1 << 10 // 1 KB
	BodyLimit            = "2M"
	ReadTimeout          = 15 * time.Second
	WriteTimeout         = 15 * time.Second
	GzipLevel            = 5
	WaitShotDownDuration = 3 * time.Second
	Dev                  = "development"
	Test                 = "test"
	Production           = "production"
)
