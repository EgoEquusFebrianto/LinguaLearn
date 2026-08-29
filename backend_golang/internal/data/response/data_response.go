package response

type LoginServiceResponse struct {
	AccessToken  	string
	RefreshToken  	string
	User 			UserProfile
}

type RefreshServiceResponse struct {
	AccessToken  	string
	RefreshToken  	string
	User 			UserProfile
	RememberMe		bool
}

type UserProfile struct {
    ID       uint64 `json:"id"`
    FullName string `json:"full_name"`
    Email    string `json:"email"`
    Role     string `json:"role"`
}

type LoginResponse struct {
    AccessToken string      `json:"access_token"`
    User        interface{} `json:"user"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}