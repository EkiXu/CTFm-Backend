package request

// User register structure
type RegisterStruct struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

// User login structure
type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EditUserStruct struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}
