package users

type CreateUserReq struct {
	Email    string
	Username string
	Password string
}

type DeleteUserReq struct {
	ID uint
}

type GetUserByIdReq struct {
	ID uint
}

type GetUserByEmailReq struct {
	Email string
}

type UpdateUserReq struct {
	ID       uint
	Email    string
	Username string
	Password string
}

type LoginReq struct {
	Username string
	Password string
}

type RegisterReq struct {
	Email    string
	Username string
	Password string
}
