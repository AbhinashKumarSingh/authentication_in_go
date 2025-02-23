package requests

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
