package core

type UserService interface {
	signUp(request userSignUp) (error)
	singIn(request userSignIn) (error)
	findById(userId string) (*User ,error)
}

type userService struct {
	port UserRepo
}

func NewUserService(Port UserRepo) UserService {
	return &userService{port: Port}
}
 
func (s userService) signUp(request userSignUp) error {
	user := User {
		Username: request.Username,
		Email: request.Email,
		Password: request.Password,
	}

	if err := s.port.signUp(user); err != nil {
		return err
	}
	return nil
}

func (s userService) singIn(request userSignIn) error {
	user := User {
		Email: request.Email,
		Password: request.Password,
	}

	if _, err := s.port.singIn(user); err != nil {
		return err
	}
	return nil
}
 
func (s userService) findById(userId string) (*User, error) {
	data, err := s.port.findById(userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}