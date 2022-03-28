package global

import "github.com/yuxing/goprojects/blogservice/pkg/logger"

var (
	ServerSetting   *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
	Logger          *logger.Logger
)
