package core

type userSignUp struct {
	Username	string	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type userSignIn struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`	
}

type webResponse struct {
	Code	int			`json:"code"`
	Status	string		`json:"status"`
	Data 	interface{}	`json:"data"`
}

