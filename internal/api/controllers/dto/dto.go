package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetUserByIdReq struct {
	ID uint `json:"id"`
}

type AuthHeader struct {
	Token string `header:"Authorization"`
}
