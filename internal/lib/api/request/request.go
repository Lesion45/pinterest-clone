package request

// Request fields
// * login
// * password
// * sessionInfo - NOT RELEASED - IN PROCESS
type RegisterReq struct {
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}
