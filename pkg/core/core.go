package core

import (
	"github.com/Kambar-ZH/simple-service/pkg/conf"
	"github.com/Kambar-ZH/simple-service/pkg/transport/rest"
)

func Run() {
	conf.GlobalConfig.Init()
	rest.InitRouter()
}
