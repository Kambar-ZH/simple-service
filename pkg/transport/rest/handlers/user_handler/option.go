package user_handler

import "github.com/Kambar-ZH/simple-service/pkg/services/user_service"

type Option func(ctl *User)

func WithUserService(userService user_service.User) Option {
	return func(ctl *User) {
		ctl.userService = userService
	}
}
