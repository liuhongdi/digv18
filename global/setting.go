package global

import (
	"github.com/liuhongdi/digv18/pkg/setting"
	"time"
)
//服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//静态目录配置
type StaticSettingS struct {
	StaticDir     string    //静态文件目录
}
//定义全局变量
var (
	ServerSetting   *ServerSettingS
	StaticSetting *StaticSettingS
)

//读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Static", &StaticSetting)
	if err != nil {
		return err
	}

	return nil
}

