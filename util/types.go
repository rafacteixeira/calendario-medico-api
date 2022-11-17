package util

type AuthRequest struct {
	Login    string
	Password string
}

type AuthError struct {
	Error   string
	Message string
}

type AuthResponse struct {
	Token string
}
