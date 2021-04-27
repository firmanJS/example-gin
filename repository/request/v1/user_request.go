package v1

//ChangePassReqModel ...
type ChangePassReqModel struct {
	ID          string `json:"id"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

//CreateUserReq ...
type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//EditUserReq ...
type EditUserReq struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
