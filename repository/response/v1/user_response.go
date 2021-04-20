package v1

//UserResponse model for response
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	RoleID   string `json:"roleId"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	ID       string `json:"id"`
	RoleID   string `json:"roleId"`
}
