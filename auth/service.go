package auth

type AuthService interface {
	RegisterEmail(dto UserReq, token string) error
	FindByVerificationToken(verificationToken string) error
	NewUser(dto UserReq) (string, error)
	UpdateUser(dto UserReq) error
	FindByID(id string) (*UserReq, error)
	DeleteByID(id string) error
}

type AuthServiceImpl struct {
	repo AuthRepo
}

func NewAuthService(repo AuthRepo) AuthService {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) RegisterEmail(dto UserReq, token string) error {

}

func (s *AuthServiceImpl) FindByVerificationToken(verificationToken string) error {

}

func (s *AuthServiceImpl) NewUser(dto UserReq) (string, error) {

}

func (s *AuthServiceImpl) UpdateUser(dto UserReq) error {

}

func (s *AuthServiceImpl) FindByID(id string) (*UserReq, error) {

}

func (s *AuthServiceImpl) DeleteByID(id string) error {

}
