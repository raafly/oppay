package helper

type ErrResponse struct {
	Status  bool `json:"status"`
	Message string `json:"message"`
	Data    any `json:"data"`
	Err     any `json:"err,omitempty"`
}

func (e *ErrResponse) Error() string {
	return e.Message
}

func NewErrNotFound() ErrResponse {
	return ErrResponse{
		Status: false,
		Message: "not found",
		Data: nil,
	}	
}

func NewErrBadRequest() ErrResponse {
	return ErrResponse{
		Status: false,
		Message: "not found",
		Data: nil,
	}	
}