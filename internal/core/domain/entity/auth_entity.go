package entity

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessToken struct {
	AccessToken string
	ExpiresAt   int64
}
