package request

type LoginRequest struct {
	Email    	string	`json:"email"`
	Password 	string	`json:"password"`
	RememberMe	bool	`json:"remember_me"`
}

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
