package types

type UsersRegisterRequest struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
	Password    string `json:"password"`
}

type UsersRegisterResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UsersActivationRequest struct {
	Token string `json:"token"`
}

type UsersActivationResponse struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
	Type   string `json:"type"`
}

type PinActivationRequest struct {
	PIN    string `json:"pin"`
	UserID int
	Email  string
}
