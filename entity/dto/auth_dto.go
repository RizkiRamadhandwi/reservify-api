package dto

type AuthRequestDto struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
}
