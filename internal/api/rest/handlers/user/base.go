package user

type UserHandler struct {
	userService userServiceInterface
}

type userServiceInterface interface{}

func New(userService userServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
