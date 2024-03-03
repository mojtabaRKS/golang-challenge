package user

type userService struct {
	userRepository userRepositoryInterface
}

type userRepositoryInterface interface {
}

func New(userRepo userRepositoryInterface) *userService {
	return &userService{
		userRepository: userRepo,
	}
}
