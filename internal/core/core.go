package core

import (
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest"
)

func Run() (err error) {
	if err = conf.GlobalConfig.Init(); err != nil {
		return
	}
	if err = rest.InitRouter(); err != nil {
		return
	}
	return
}
