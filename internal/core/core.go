package core

import (
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/transport/rest"
)

func Run() {
	conf.GlobalConfig.Init()
	rest.InitRouter()
}
