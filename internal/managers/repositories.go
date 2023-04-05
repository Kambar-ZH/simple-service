package managers

import (
	"github.com/Kambar-ZH/simple-service/internal/conf"
	"github.com/Kambar-ZH/simple-service/internal/repositories/common/user_repo"
	"sync"
)

var repositories = &Repositories{}

type Repositories struct {
	userRepositoryInit sync.Once
	userRepository     user_repo.User
}

func (r Repositories) User() user_repo.User {
	r.userRepositoryInit.Do(func() {
		r.userRepository = user_repo.New(conf.GlobalConfig.GormDB)
	})
	return r.userRepository
}
